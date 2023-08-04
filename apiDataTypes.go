package api

type APIData struct {
	name    string
	rootURL string
}

type Endpoint struct {
	endpointURL string
	parent      *APIData
}

type nameValuePair struct {
	name  string
	value string
}

type reqQuery nameValuePair
type reqHeader nameValuePair
type reqBody nameValuePair

type Request struct {
	endPoint *Endpoint
	queries  []reqQuery
	headers  []reqHeader
	body     []reqBody
}

type Result struct {
	Body []byte
}
