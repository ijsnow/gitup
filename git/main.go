package main

import (
	"log"
	"net/http"
	"os"
	"os/user"
	"regexp"

	"github.com/ijsnow/gittp"

	"gitup.io/isaac/gitup/datastore"
	"gitup.io/isaac/gitup/settings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := datastore.Connect(settings.App.Database.MongoDB.URL); err != nil {
		panic(err)
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("~")

	path := re.ReplaceAllLiteralString(settings.App.RepoDir, usr.HomeDir)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	r := gittp.NewHandler(path, check)

	log.Printf("Listening at 0.0.0.0:%s\n", settings.App.Server.Port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+settings.App.Server.Port, r))
}

const (
	myGitUsername = "isaac"
	myGitPassord  = "supersecret"
)

func check(r gittp.RequestInfo) (bool, int) {
	return true, 0
}
