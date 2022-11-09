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
