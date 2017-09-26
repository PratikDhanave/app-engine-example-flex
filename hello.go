package main

import (
	"fmt"
	"net/http"
	"cloud.google.com/go/datastore"
	"log"
	"context"
	"os"
)


var datastoreClient *datastore.Client


func main() {

	ctx := context.Background()

	projectID := os.Getenv("projectID") //"golangproject-180913"

	var err error

	datastoreClient, err = datastore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/save", handler)
	http.HandleFunc("/retrieve", retrieve)
  log.Fatal(http.ListenAndServe(":8080", nil))
	
}




type Store struct {
	Input string
}



func retrieve(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	projectID := "golangproject-180913"

	var err error

	datastoreClient, err = datastore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatal(err)
	}

var val []*Store

	html := ""

	keys, err := datastoreClient.GetAll(ctx, datastore.NewQuery("Store"), &val)

	for i, key := range keys {
	    fmt.Println(key)
	    fmt.Println(val[i])
			html = html + "\n" + val[i].Input
	}

	w.Header().Set("Content-Type", "text/html")
        fmt.Fprint(w,html)
}


func handler(w http.ResponseWriter, r *http.Request) {

  ctx := context.Background()

	param := r.URL.Query().Get("input")

	entity := &Store{}

  entity.Input = param

	key := datastore.IncompleteKey("Store", nil)

	_, err := datastoreClient.Put(ctx, key, entity)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, "Value = ",param,"\tstored in Database")
}
