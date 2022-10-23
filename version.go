package goapi

import (
	"fmt"

	"github.com/pjsoftware/goapi/auth"
)

type APIData struct {
	AuthData auth.Data
}

func New() *APIData {
	rv := APIData{}
	rv.AuthData.AuthType = auth.None
	return &rv
}

func (a *APIData) Version() {
	fmt.Println("goAPI v0.0.1")
}

func (a *APIData) Type() {
	switch a.AuthData.AuthType {
	case auth.None:
		fmt.Println("Auth: None")
	case auth.APIKey:
		fmt.Println("Auth: API Key")
	case auth.TwoLeggedOAuth:
		fmt.Println("Auth: Two-Legged OAuth")
	case auth.ThreeLeggedOAuth:
		fmt.Println("Auth: Three-Legged OAuth")
	}

}