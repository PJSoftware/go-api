# go-api

## Getting started

At the command prompt for your project, use:

```sh
go get github.com/pjsoftware/go-api
```

In your `Go` code:

```go
import {
  goapi "github.com/pjsoftware/go-api"
}

myAPI := goapi.New(apiURL)
ep  := myAPI.NewEndpoint(endpointURL)
req := ep.NewRequest().AddQuery(qryName,qryValue)
r := req.GET()
```
