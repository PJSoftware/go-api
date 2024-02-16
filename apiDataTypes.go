package api

// APIData is the main export from go-api; it is generated via api.New()
type APIData struct {
	name    string
	rootURL string
}

// Each Endpoint should be individually managed by the client code. An Endpoint
// is generated via api.NewEndpoint()
type Endpoint struct {
	endpointURL string
	parent      *APIData
}

// An individual Request is used to communicate with the external API. A Request
// is generated via (*Endpoint).NewRequest()
type Request struct {
	endPoint *Endpoint
	queries  []reqQuery
	headers  []reqHeader
	bodyKV   []reqBody
	bodyTXT  string
	hasBody  bool
}

type reqQuery keyValuePair
type reqHeader keyValuePair
type reqBody keyValuePair

type keyValuePair struct {
	key   string
	value string
}

type Response struct {
	Status int
	Body   string
}
