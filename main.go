package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jaysonesmith/imageswitch/api"
)

func main() {
	api := api.NewAPI()

	fmt.Println("starting up imageswitch on port 9001")

	log.Fatal(http.ListenAndServe(":9001", api.Router))
}
