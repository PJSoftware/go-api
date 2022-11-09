package api

import "github.com/pjsoftware/go-api/auth"

type APIData struct {
	name     string
	rootURL  string
	auth     auth.Data
}

type Endpoint struct {
	endpointURL string
	parent *APIData
}

type reqQuery struct {
	name string
	value string
}
type reqHeader struct {
	name string
	value string
}

type Request struct {
	endPoint *Endpoint
	queries []reqQuery
	headers []reqHeader
}

type Result struct {
	Body []byte
}
