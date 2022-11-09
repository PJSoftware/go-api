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
