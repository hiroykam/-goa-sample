package services

import (
	"crypto/rsa"
	"fmt"
	"github.com/hiroykam/goa-sample/sample_middleware"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hiroykam/goa-sample/app"
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/pkg/errors"
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
func (s *AuthSharedService) signedToken(id int, exp time.Time) (string, string, *sample_error.SampleError) {
	// Generate JWT
	token := jwt.New(jwt.SigningMethodRS512)
	exp = time.Now().Add(tokenExpiredDuration)
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", "", sample_error.NewSampleError(sample_error.InternalError, err.Error())
	}
	jti := uuid.String()
	token.Claims = jwt.MapClaims{
		"iss":    "Issuer",                         // who creates the token and signs it
		"aud":    "Audience",                       // to whom the token is intended to be sent
		"exp":    exp.Unix(),                       // time when the token will expire (10 minutes from now)
		"jti":    jti,                              // a unique identifier for the token
		"iat":    time.Now().Unix(),                // when the token was issued/created (now)
		"nbf":    2,                                // time before which the token is not yet valid (2 minutes ago)
		"sub":    id,                               // the subject/principal is whom the token is about
		"scopes": "api:access",                     // token scope - not a standard claim
	}
	signedToken, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", "", sample_error.NewSampleError(sample_error.InternalError, fmt.Sprintf("failed to sign token: %s", err))
	}
	return signedToken, jti, nil
}


func (s *AuthSharedService) IssueTokens(id int) (*app.Auth, string, *sample_error.SampleError) {
	tk, _, err := s.signedToken(id, time.Now().Add(tokenExpiredDuration))
	if err != nil {
		return nil, "", err
	}

	rtk, jti, err := s.signedToken(id, time.Now().Add(tokenExpiredDuration))
	if err != nil {
		return nil, "", err
	}

	return &app.Auth{
		Token: &app.Token{
			ExpiredAt: time.Now().Add(tokenExpiredDuration),
			Token:     tk,
		},
		RefreshToken: &app.RefreshToken{
			ExpiredAt:    time.Now().Add(refreshTokenExpiredDuration),
			RefreshToken: rtk,
		},
	}, jti, nil
}

// Secure runs the secure action.
func (c *AuthSharedService) GetId(token *jwt.Token) (int, *sample_error.SampleError) {
	// Retrieve the token claims
	if token == nil {
		return 0, sample_error.NewSampleError(sample_error.UnAuthorized, "JWT token is missing from context")
	}
	claims := token.Claims.(jwt.MapClaims)


	// Use the claims to authorize
	return int(claims["sub"].(float64)), nil
}

func (c *AuthSharedService) VerifyToken(tokenString string) (string, *sample_error.SampleError) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		keys, err := sample_middleware.LoadJWTPublicKeys()
		if err != nil {
			return nil, err
		}
		return keys[0], nil
	})

	if err != nil {
		fmt.Println("test")
		return "", sample_error.NewSampleError(sample_error.UnAuthorized, err.Error())
	}

	err = parsedToken.Claims.Valid()
	if err !=nil {
		return "", sample_error.NewSampleError(sample_error.UnAuthorized, err.Error())
	}

	claims := parsedToken.Claims.(jwt.MapClaims)

	return claims["jti"].(string), nil
}
