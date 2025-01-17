package main

import (
	"github.com/LavaJover/dronwallet/auth/db"
)

func main(){
	dsn := "host=localhost user=postgres password=admin dbname=dronwallet port=5432 sslmode=disable"
	db.InitDB(dsn)
}