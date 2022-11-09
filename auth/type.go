package auth

import "fmt"

func Type(authType AuthEnum) {
	switch authType {
	case None:
		fmt.Println("Auth: None")
	case APIKey:
		fmt.Println("Auth: API Key")
	case TwoLeggedOAuth:
		fmt.Println("Auth: Two-Legged OAuth")
	case ThreeLeggedOAuth:
		fmt.Println("Auth: Three-Legged OAuth")
	}
}