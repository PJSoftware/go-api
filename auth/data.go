package auth

type AuthEnum int32

const (
	None AuthEnum = iota
	APIKey
	TwoLeggedOAuth
	ThreeLeggedOAuth
)

type Data struct {
	AuthType AuthEnum
}
