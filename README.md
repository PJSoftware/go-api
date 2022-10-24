# Go API Module

A general library for interfacing with APIs

## Simple Usage: No Auth

```go
api := goapi.New(apiURL)
res := api.Get(endpoint)
```

## Simple Usage: ThreeLegged OAuth

```go
api := goapi.New(apiURL).AuthThreeLegged(id, secret, authURL, callbackURL)
qry := api.NewQuery(endpoint)
qry.AddHeader(hdrName,hdrValue)
res := qry.Get()
```
