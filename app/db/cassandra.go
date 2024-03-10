package db

import (
	"errors"
	"log"
	"sync"

	"github.com/aykhans/oh-my-url/app/utils"
	"github.com/gocql/gocql"
)

type CurrentID struct {
	ID int
	mu sync.RWMutex
}

type Cassandra struct {
	db        *gocql.Session
	currentID *CurrentID
}

func (c *Cassandra) Init() {
	err := c.db.Query(`
		CREATE TABLE IF NOT EXISTS url (
			id int,
			key text,
			url text,
			count int,
			PRIMARY KEY ((key), count)
		) WITH CLUSTERING ORDER BY (count DESC)
		AND compaction = {'class': 'LeveledCompactionStrategy'};
	`).Consistency(gocql.All).Exec()
	if err != nil {
		panic(err)
	}

	var id int
	err = c.db.
		Query(`SELECT MAX(id) FROM url LIMIT 1`).
		Consistency(gocql.One).
		Scan(&id)

	if err != nil {
		panic(err)
	}
	c.currentID.ID = id
}

func (c *Cassandra) CreateURL(url string) (string, error) {
	c.currentID.mu.Lock()
	defer c.currentID.mu.Unlock()

	id := c.currentID.ID + 1
	key := utils.GenerateKey(id)
	m := make(map[string]interface{})

	query := `INSERT INTO url (id, key, url, count) VALUES (?, ?, ?, ?) IF NOT EXISTS`
	applied, err := c.db.Query(query, id, key, url, 0).Consistency(gocql.All).MapScanCAS(m)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if !applied {
		log.Println("Failed to insert unique key")
		return "", errors.New("an error occurred, please try again later")
	}
	c.currentID.ID = id

	return key, nil
}

func (c *Cassandra) GetURL(key string) (string, error) {
	var url string
	err := c.db.
		Query(`SELECT url FROM url WHERE key = ? LIMIT 1`, key).
		Consistency(gocql.One).
		Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}

// just in case
// const urlKeyCreateFunction = `
// CREATE FUNCTION IF NOT EXISTS oh_my_url.generate_url_key(n int) RETURNS NULL ON NULL INPUT RETURNS text LANGUAGE java AS '
//     String keyCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
//     StringBuilder result = new StringBuilder();
//     int base = keyCharacters.length();

//     while (n > 0) {
//         n--;
//         result.append(keyCharacters.charAt(n % base));
//         n = (int) Math.floor(n / base);
//     }

//     return result.reverse().toString();
// ';
// `
