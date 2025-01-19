package patterns

import "sync"

type Database struct {
	Connection string
}

var instance *Database
var once sync.Once

func GetDatabaseInstance() *Database {
	once.Do(func() {
		instance = &Database{Connection: "Connected to Employee Database"}
	})
	return instance
}
