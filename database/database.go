package database

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)
func (d database) Get(key string) ([]byte, error) {
	return d.db.Get([]byte(key), nil)
}

func (d database) Set(key string,value []byte) error {
	return d.db.Put([]byte(key), value, nil)
}

func (d database) Delete(key string) error {
	return d.db.Delete([]byte(key), nil)
}

func (d database) Close() error {
	//return d.db.Close()
	return nil
}



func New() (LevelDB, error) {
	db, err := leveldb.OpenFile("dbms-db-a", nil)
	if err != nil {
		fmt.Printf("Error in creating DATABASE: %s\n", err)
		return nil, err
	}

	return database{ db }, err
}


type (
	LevelDB interface {
		Get(key string) ([]byte, error)
		Set(key string, value []byte) error
		Delete(key string) error
		Close() error
	}

	database struct {
		db *leveldb.DB
	}
)