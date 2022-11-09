package api

import (
	"fmt"

	"github.com/pjsoftware/go-api/auth"
)

const Version = "0.0.1"

func version() {
	fmt.Printf("go-api version: %s\n", Version)
}

func (a *APIData) Type() {
	auth.Type(a.AuthData.AuthType)
}