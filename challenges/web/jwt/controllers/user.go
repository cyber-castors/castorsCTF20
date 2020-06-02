package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"../models"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

//Once user is logged in
func (uc Controller) Success(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := c.Value

	//new isntance of claims with composite literal
	claims := &models.Claims{}

	//Parse jwt string and store in claims
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//if after AAAAALL that the user has valid token, welcome her/him
	fmt.Fprintf(w, "Welcome %s!", claims.Username)
}

func (uc Controller) Refresh(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknString := c.Value
	claims := &models.Claims{}
	tkn, err := jwt.ParseWithClaims(tknString, claims, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//Only issue new token if old token is within 30 secs of expiring,
	//if not return bad request
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Create new token
	expTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(models.JwtKey)
	if !handleError(w, err) {
		return
	}

	//Set new token as cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenStr,
		Expires: expTime,
	})
}
