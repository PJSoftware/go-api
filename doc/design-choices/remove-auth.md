# Remove Auth Support

As currently written, `go-api` has minimal support in place for `auth` handling.

My only current experience with `Auth0` is via a single API -- and it recently
changed its approach. I suspect I don't know enough to be able to implement a
"general" Auth module here, so I have taken the decision to remove it.

However, I shall document the little I had, in case I ever decide to re-add auth
support.

I suspect that if I do ever decide to implement a general auth package, it will
be independent of `go-api` (if that makes more sense!)

## Relevant Code

### auth/authDataTypes.go

```go
package auth

type authEnum int32

const (
  none authEnum = iota
  apiKey
  twoLeggedOAuth
  threeLeggedOAuth
)

type Data struct {
  authType authEnum
}
```

### auth/types.go

```go
package auth

func (d *Data) Type() string {
  switch d.authType {
  case none:
    return "Auth: None"
  case apiKey:
    return "Auth: API Key"
  case twoLeggedOAuth:
    return "Auth: Two-Legged OAuth"
  case threeLeggedOAuth:
    return "Auth: Three-Legged OAuth"
  }
  return ""
}

func (d *Data) NoAuth() {
  d.authType = none
}
```
