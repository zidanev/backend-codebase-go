package token

import (
	"context"
	"crypto/rsa"
	"fmt"
	"time"

	"codebase-go/bin/pkg/utils"

	"github.com/dgrijalva/jwt-go"
)

func Generate(ctx context.Context, privateKey *rsa.PrivateKey, payload *Claim, expired time.Duration) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		now := time.Now()
		exp := now.Add(expired)

		token := jwt.New(jwt.SigningMethodRS256)

		claims := jwt.MapClaims{
			"exp": exp.Unix(),
			"iat": now.Unix(),
			"sub": payload.UserId,
		}

		token.Claims = claims

		tokenString, err := token.SignedString(privateKey)
		if err != nil {
			output <- utils.Result{Error: err}
			return
		}

		output <- utils.Result{Data: tokenString}

	}()

	return output
}

func Validate(ctx context.Context, publicKey *rsa.PublicKey, tokenString string) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		tokenParse, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})

		var errToken string
		switch ve := err.(type) {
		case *jwt.ValidationError:
			if ve.Errors == jwt.ValidationErrorExpired {
				errToken = "token has been expired"
			} else {
				errToken = "token parsing error"
			}
		}

		if len(errToken) > 0 {
			output <- utils.Result{Error: fmt.Errorf(errToken)}
			return
		}
		if !tokenParse.Valid {
			output <- utils.Result{Error: fmt.Errorf("token parsing error")}
			return
		}
		mapClaims, _ := tokenParse.Claims.(jwt.MapClaims)
		var tokenClaim Claim
		tokenClaim.UserId, _ = mapClaims["sub"].(string)

		output <- utils.Result{Data: tokenClaim}
	}()

	return output
}
