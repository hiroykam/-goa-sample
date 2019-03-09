package services

import (
	"crypto/rsa"
	"fmt"
	"github.com/hiroykam/goa-sample/app"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/satori/go.uuid"
)

const tokenExpiredDuration = time.Hour
const refreshTokenExpiredDuration = 24 * time.Hour

// JWTController implements the JWT resource.
type AuthSharedService struct {
	privateKey *rsa.PrivateKey
}

func NewAuthSharedService() (*AuthSharedService, *sample_error.SampleError) {
	b, err := ioutil.ReadFile("./jwtkey/jwt.key")
	if err != nil {
		return nil, sample_error.NewSampleError(sample_error.InternalError, err.Error())
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, sample_error.NewSampleError(sample_error.InternalError, fmt.Sprintf("jwt: failed to load private key: %s", err))
	}
	return &AuthSharedService{
		privateKey: privateKey,
	}, nil
}

// Signin creates JWTs for use by clients to access the secure endpoints.
func (s *AuthSharedService) signedToken(id int, exp time.Time) (string, *sample_error.SampleError) {
	// Generate JWT
	token := jwt.New(jwt.SigningMethodRS512)
	exp = time.Now().Add(tokenExpiredDuration)
	token.Claims = jwt.MapClaims{
		"iss":    "Issuer",                         // who creates the token and signs it
		"aud":    "Audience",                       // to whom the token is intended to be sent
		"exp":    exp.Unix(),                       // time when the token will expire (10 minutes from now)
		"jti":    uuid.Must(uuid.NewV4()).String(), // a unique identifier for the token
		"iat":    time.Now().Unix(),                // when the token was issued/created (now)
		"nbf":    2,                                // time before which the token is not yet valid (2 minutes ago)
		"sub":    id,                               // the subject/principal is whom the token is about
		"scopes": "api:access",                     // token scope - not a standard claim
	}
	signedToken, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", sample_error.NewSampleError(sample_error.InternalError, fmt.Sprintf("failed to sign token: %s", err))
	}
	return signedToken, nil
}


func (s *AuthSharedService) IssueTokens(id int) (*app.Auth, *sample_error.SampleError) {
	exp := time.Now().Add(tokenExpiredDuration)

	tk, err := s.signedToken(id, time.Now().Add(tokenExpiredDuration))
	if err != nil {
		return nil, err
	}

	return &app.Auth{
		Token: &app.Token{
			ExpiredAt: exp,
			Token:     tk,
		},
	}, nil
}

// Secure runs the secure action.
func (c *AuthSharedService) GetId(token *jwt.Token) (int, *sample_error.SampleError) {
	// Retrieve the token claims
	if token == nil {
		return 0, sample_error.NewSampleError(sample_error.InternalError, "JWT token is missing from context")
	}
	claims := token.Claims.(jwt.MapClaims)

	// Use the claims to authorize
	return int(claims["sub"].(float64)), nil
}
