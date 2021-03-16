package addtransport

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/basic"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http"
	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	publicKey = "-----BEGIN RSA PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA/0GpgXEy0vc+lZtHbjgZ\nSh5S/MeSSLdizKCd3gNAxIqc1/jJvPOQS76H+UUvHlxZVEt1oCYa0maxlxKlwKth\n6GDFNVkTi5GziIxPKVEgrLP0U/W5NHcaprZJxPHTduR7MOC34IV3Lz8OxhqTWOl0\nAnClGj9kPhHIEenuDaO/IE3oBpx8XBS2CNTIK9fu2UbRCNMceCjPFx6XIwFLhfTb\ns8TedX0c4O5FJlRkm1XRA80nDMSZWRYHfYyJyTkevASm/Bxfxzd/Jf8XwJu36PMn\nsBbMHyQj55y7RYnGFBkgR5MNr02b9VUjCYRfUisrkm1R5D+Kw0Bvp+RyY+Bjyw+U\nBwIDAQAB\n-----END RSA PUBLIC KEY-----"
)

var (
	basicUser = "lol"
	basicPass = "lul"
)

func getKey() jwtgo.Keyfunc {
	pub, err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return nil
	}

	return func(token *jwtgo.Token) (interface{}, error) {
		return pub, nil
	}
}

func AuthMiddleware(authType string, endpoint endpoint.Endpoint) endpoint.Endpoint {
	switch t := authType; t {
	case "jwt":
		return jwt.NewParser(getKey(), jwtgo.SigningMethodRS256, jwt.StandardClaimsFactory)(endpoint)
	case "basic":
		return basic.AuthMiddleware(basicUser, basicPass, "Example Realm")(endpoint)
	default:
		return endpoint
	}
}

func GetContext(authType string) http.RequestFunc {
	switch t := authType; t {
	case "jwt":
		return jwt.HTTPToContext()
	case "basic":
		return httptransport.PopulateRequestContext
	default:
		return httptransport.PopulateRequestContext
	}
}
