package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"code.google.com/p/goauth2/oauth"
	"code.google.com/p/google-api-go-client/calendar/v3"
	"github.com/masahide/get-cybozu-schedule/lib"
)

func main() {

	flag.Usage = lib.Usage
	flag.Parse()

	if *lib.Version {
		fmt.Printf("%s\n", lib.ShowVersion())
		return
	}

	port := 3000
	config, err := lib.Parse("google.json")
	if err != nil {
		log.Fatalf("Error Server: %v", err)
		return
	}
	transport := oauth.Transport{
		Config: &oauth.Config{
			ClientId:     config.ClientID,
			ClientSecret: config.ClientSecret,
			RedirectURL:  fmt.Sprintf("%s:%d", "http://localhost", port),
			Scope:        "https://www.googleapis.com/auth/calendar",
			AuthURL:      "https://accounts.google.com/o/oauth2/auth",
			TokenURL:     "https://accounts.google.com/o/oauth2/token",
			TokenCache:   oauth.CacheFile("cache.json"),
		},
	}

	err = lib.GoogleOauth(&lib.GoogleToken{&transport}, lib.LocalServerConfig{port, 30, runtime.GOOS})
	if err != nil {
		log.Fatalf("Error Server: %v", err)
		return
	}

	var svc *calendar.Service
	var cl *calendar.CalendarList

	svc, err = calendar.New(transport.Client())

	if err != nil {
		log.Fatalf("Error calendar.New: %v", err)
		return
	}

	cl, err = svc.CalendarList.List().Do()

	if err != nil {
		log.Fatalf("Error CalendarList.List(): %v", err)
		return
	}

	fmt.Printf("--- Your calendars ---\n")

	for _, item := range cl.Items {
		//fmt.Printf("%v, %v\n", item.Summary, item.Description)
		fmt.Printf("%# v\n", item)
	}

}
