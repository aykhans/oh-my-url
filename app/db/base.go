package db

import (
	"fmt"
	"github.com/aykhans/oh-my-url/app/config"
	"github.com/gocql/gocql"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type DB interface {
	Init()
	CreateURL(url string) (string, error)
	GetURL(key string) (string, error)
}

func GetDB() DB {
	db := config.GetDB()
	switch db {
	case "postgres":
		postgresConf := config.GetPostgresConfig()
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			postgresConf.HOST,
			postgresConf.USER,
			postgresConf.PASSWORD,
			postgresConf.DBNAME,
			postgresConf.PORT,
			"disable",
			"UTC",
		)

		var db *gorm.DB
		var err error
		for i := 0; i < 5; i++ {
			db, err = gorm.Open(gormPostgres.Open(dsn), &gorm.Config{})
			if err == nil {
				break
			}
			time.Sleep(3 * time.Second)
		}
		if err != nil {
			panic(err)
		}
		return &Postgres{gormDB: db}

	case "cassandra":
		cassandraConf := config.GetCassandraConfig()
		cluster := gocql.NewCluster(cassandraConf.CLUSTERS...)
		cluster.Keyspace = cassandraConf.KEYSPACE
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: cassandraConf.USER,
			Password: cassandraConf.PASSWORD,
		}

		var db *gocql.Session
		var err error
		for i := 0; i < 60; i++ {
			db, err = cluster.CreateSession()
			if err == nil {
				break
			}
			time.Sleep(3 * time.Second)
		}
		if err != nil {
			panic(err)
		}
		return &Cassandra{
			db: db,
			currentID: &CurrentID{},
		}
	}
	panic("unknown db")
}
