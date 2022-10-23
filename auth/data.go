package auth

type AuthType int32

const (
	None AuthType = iota
	APIKey
	TwoLeggedOAuth
	ThreeLeggedOAuth
)

type Data struct {
	AuthType AuthType
}
