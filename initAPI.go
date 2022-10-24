package goapi

import (
	"fmt"

	"github.com/pjsoftware/goapi/auth"
)

type APIData struct {
	RootURL  string
	AuthData auth.Data
}

func New(rootURL string) *APIData {
	fmt.Printf("API: %s\n", rootURL)
	rv := APIData{}
	rv.RootURL = rootURL
	rv.AuthData.AuthType = auth.None
	return &rv
}
