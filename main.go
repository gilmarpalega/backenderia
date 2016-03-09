package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/guilhermebr/backenderia/apps/page"
	"github.com/nuveo/utils/adapter/mongo"
)

func main() {
	//Register Apps

	config := map[string]string{
		"db_uri": fmt.Sprintf("%s:%s", os.Getenv("DB_PORT_27017_TCP_ADDR"),
			os.Getenv("DB_PORT_27017_TCP_PORT")),
		"db_name": "backenderia",
	}

	db := mongo.New(config["db_uri"], config["db_name"])
	err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}

	n := negroni.Classic()
	r := mux.NewRouter()

	page.Register(r, db)

	n.UseHandler(r)

	n.Run(":3000")
}