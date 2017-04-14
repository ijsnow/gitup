package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ijsnow/gitup/api/config"
	"github.com/ijsnow/gitup/api/router"
	"github.com/ijsnow/gitup/datastore"
	"github.com/ijsnow/gitup/utils/fspaths"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := datastore.Connect(fspaths.ExpandTildePath(config.App.Database.Path)); err != nil {
		panic(err)
	}

	r := router.InitRouter()

	log.Printf("Listening at 0.0.0.0:%s\n", config.App.Server.Port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+config.App.Server.Port, r))

	fmt.Println("Closing database")
	datastore.Close()
}
