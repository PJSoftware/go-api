package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"

	goapi "github.com/pjsoftware/go-api"
)

type Object struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil)).With("app_name", "send-json")
	slog.SetDefault(logger)

	api := goapi.New("https://ZZZZZ.wiremockapi.cloud")

	getJSON(api)
	postJSON(api)
}

func getJSON(api *goapi.APIData) {
	ep := api.NewEndpoint("json/1")

	fmt.Printf("EP: %s\n", ep.URL())
	req := ep.NewRequest()
	res, err := req.GET()
	if err != nil {
		log.Fatalf("GET error: %v", err)
	}

	fmt.Printf("RESPONSE STATUS: %v\n", res.Status)
	fmt.Printf("RESPONSE BODY: %v\n", res.Body)
}

func postJSON(api *goapi.APIData) {
	ep := api.NewEndpoint("json")

	obj := &Object{}
	obj.ID = 12345
	obj.Value = "abc-def-ghi"

	fmt.Printf("EP: %s\n", ep.URL())
	req := ep.NewRequest()
	req.SetBodyJSON(obj)

	res, err := req.POST()
	if err != nil {
		if errors.Is(err, goapi.ErrSuccess) {
			fmt.Printf("Success returned: %v\n", err)
		} else {
			log.Fatalf("POST error: %v\n", err)
		}
	}

	fmt.Printf("RESPONSE STATUS: %v\n", res.Status)
	fmt.Printf("RESPONSE BODY: %v\n", res.Body)
}
