package db

import (
	"log"
	"strconv"
	"sync"

	"github.com/tidwall/buntdb"
)

var db database

// database is the collection of key-value pairs containing a mutex
type database struct {
	conn *buntdb.DB
	mux  sync.Mutex
}

// Create or connect with database
func Connect(file string) {
	conn, err := buntdb.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	db = database{
		conn: conn,
	}
}

// Close will sync the database to file and nil the object
func Close() {
	log.Fatal(db.conn.Close())
}

// Lock will lock the databse from read/writes
func Lock() {
	db.mux.Lock()
}

// Unlock will unlock the databse from read/writes
func Unlock() {
	db.mux.Unlock()
}

// Set will insert/update the key and string value
func Set(key, value string) error {
	return db.conn.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
}

// Set will insert/update the key and int value
func SetINT(key string, value int) error {
	return Set(key, strconv.Itoa(value))
}

// Get will return the string value for given key
func Get(key string) (value string) {
	_ = db.conn.View(func(tx *buntdb.Tx) (err error) {
		value, err = tx.Get(key)
		return
	})
	return
}

// GetINT will return the int value for given key
func GetINT(key string) int {
	return str2int(Get(key))
}

// str2int converts string to int and defaults to 0
func str2int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		i = 0
	}
	return i
}
