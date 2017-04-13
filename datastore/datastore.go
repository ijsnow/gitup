package datastore

import (
	boltdb "github.com/boltdb/bolt"
	"gitup.io/isaac/gitup/datastore/bolt"
)

// DataStore is the type for a database session
type DataStore struct {
	db *boltdb.DB

	Users    *bolt.Users
	Sessions *bolt.Sessions
}

func initialDataStore(db *boltdb.DB) *DataStore {
	return &DataStore{
		db:       db,
		Users:    nil,
		Sessions: nil,
	}
}

// NewDataStore returns a new datastore with a copied session
func (ds *DataStore) NewDataStore() *DataStore {
	return &DataStore{
		db:       ds.db,
		Users:    bolt.NewUsers(ds.db),
		Sessions: bolt.NewSessions(ds.db),
	}
}

// Close the session
func Close() {
	Store.db.Close()
}
