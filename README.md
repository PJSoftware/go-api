# Go API Module

A general library for interfacing with APIs. The basic procedure is:

- Call `goapi.New()` to create a new `api` object. This can be stored for future use.
- Use `api.AuthXXX()` to perform any necessary authentication.
- Create a new `ep` endpoint with `api.NewEndpoint()`. Again, these endpoints can be stored and reused.
- Create a single-use `req` request object: `ep.NewRequest()`
- Modify with `req.AddQuery()`, `req.AddHeader()`, etc as required.
  - AddQuery, AddHeader, etc can be chained: `ep.NewRequest().AddHeader()`
- Call `req.GET()` or `req.POST()` as needed. By default, the request is discarded once used.
- All processing of the returned data is done by the client application. `go-api` makes no attempt to interpret what it receives; it simply passes it back in a `goapi.Result` struct.

## Simple Usage: No Auth

```go
api := goapi.New(apiURL)
ep  := api.NewEndpoint(endpointURL)
req := ep.NewRequest().AddQuery(qryName,qryValue)
r := req.GET()
```

A simple example of this from the "Chuck Norris" API (see [go-api-chuck](https://github.com/PJSoftware/go-api-chuck))
is as follows:

```go
chuck := goapi.New("http://api.chucknorris.io/jokes")
ep  := chuck.NewEndpoint("/random")
req := ep.NewRequest().AddQuery("category", "food")
r := req.GET()
```

After receiving the results, recommended processing may look like this:

```go
r := req.GET()
data := &MyDataStruct{}
json.Unmarshal(r.Body, data)
```

## Simple Usage: ThreeLegged OAuth

```go
api := goapi.New(apiURL)
api.AuthThreeLegged(id, secret, authURL, callbackURL)

ep  := api.NewEndpoint(endpointURL)
req := ep.NewRequest().AddHeader(hdrName,hdrValue)
r := req.GET()
```
