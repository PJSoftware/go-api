package goapi

import (
	"fmt"

	"github.com/pjsoftware/goapi/auth"
)

const Version = "0.0.1"

func (a *APIData) Version() {
	fmt.Printf("goAPI version: %s\n", Version)
}

func (a *APIData) Type() {
	auth.Type(a.AuthData.AuthType)
}