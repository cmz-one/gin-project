package main

import (
	db "rest/databases"
)







func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8080")
}
