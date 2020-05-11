package demo_pool

import "database/sql"

type PoolByte chan []byte

func NewPoolByte() PoolByte {
	return make(chan []byte)
}
func (p PoolByte) Put(b []byte) {
	select {
	case p <- b:
	default:
	}
}
func (p PoolByte) Get() []byte {
	var v []byte
	select {
	case v = <-p:
	default:
		v = make([]byte, 3)
	}
	return v
}

type PoolDB chan *sql.DB

func NewPoolDB(cap int) PoolDB {
	return make(chan *sql.DB, cap)
}
func (p PoolDB) Put(db *sql.DB) {
	select {
	case p <- db:
	default:

	}
}
func ConDB() *sql.DB {
	db, err := sql.Open("mysql", "")
	if err != nil {

	}
	return db
}
func (p PoolDB) Get() *sql.DB {
	var db *sql.DB
	select {
	case db = <-p:
	default:
		db = ConDB()
	}
	return db
}
