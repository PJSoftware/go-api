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
	version()
	fmt.Printf("API: %s\n", rootURL)
	rv := APIData{}
	rv.RootURL = rootURL
	rv.AuthData.AuthType = auth.None
	return &rv
}

func (a *APIData) NewQuery(endPoint string) *APIQuery {
	qry := &APIQuery{}
	qry.EndPoint = endPoint
	return qry
}

func (a *APIData) Get(endPoint string) *QRYResult {
	fmt.Printf("EndPoint: %s\n", endPoint)
	qry := a.NewQuery(endPoint)
	rv := qry.Get()
	return rv
}
