package db

import (
	"gorm.io/gorm"
)

type Postgres struct {
	gormDB *gorm.DB
}

type url struct {
	ID  uint   `gorm:"primaryKey"`
	Key string `gorm:"unique;not null;size:15;default:null"`
	URL string `gorm:"not null;default:null"`
	Count int `gorm:"not null;default:0"`
}

func (p *Postgres) Init() {
	err := p.gormDB.AutoMigrate(&url{})
	if err != nil {
		panic("failed to migrate database")
	}
	tx := p.gormDB.Exec(urlKeyCreateTrigger)
	if tx.Error != nil {
		panic("failed to create trigger")
	}
}

func (p *Postgres) CreateURL(mainUrl string) (string, error) {
	url := url{URL: mainUrl}
	tx := p.gormDB.Create(&url)
	if tx.Error != nil {
		return "", tx.Error
	}
	return url.Key, nil
}

func (p *Postgres) GetURL(key string) (string, error) {
	var result url
	tx := p.gormDB.Where("key = ?", key).First(&result)
	if tx.Error != nil {
		return "", tx.Error
	}
	result.Count++
	p.gormDB.Save(&result)
	return result.URL, nil
}

const urlKeyCreateTrigger = `
CREATE OR REPLACE FUNCTION generate_url_key()
	RETURNS TRIGGER AS $$
DECLARE
	key_characters TEXT := 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
	key TEXT := '';
	base INT := LENGTH(key_characters);
	n INT := NEW.id;
BEGIN
	WHILE n > 0 LOOP
	n := n - 1;
	key := SUBSTRING(key_characters FROM (n % base) + 1 FOR 1) || key;
	n := n / base;
	END LOOP;
	NEW.key := key;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_generate_url_key on "public"."urls";

CREATE TRIGGER trigger_generate_url_key
	BEFORE INSERT ON urls
	FOR EACH ROW
	EXECUTE FUNCTION generate_url_key();
`
