package models

import (
	"text/template"

	"github.com/dgrijalva/jwt-go"
)

//Key is in Rockyou.txt on purpose
var JwtKey = []byte("chocolatito")

//Templates
var Tpl *template.Template

//Local "Database"
var DBUsers = map[string]User{}
var DBSessions = map[string]string{}

//Exported user struct
type User struct {
	Username  string `json:"name"`
	Password  string `json:"passwd"`
	Firstname string `json:"fname"`
	Lastname  string `json:"lname"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var Users = map[string]string{
	"user1":    "kTQiLQPN497H4TqzzSLr",
	"achiles":  "tjE2bBJMzNkqBDigUL68",
	"medusa":   "v6arvd9cl+sHyqSUPdoh",
	"cronos":   "ygaCGm4VLnLeazgEQTv3",
	"zeus":     "sEb6uRtAqqaa8CSLH9s0",
	"thanos":   "vDa4IA2/VCA4S83p9+cq",
	"godofwar": "uBMbStwArnKqQlg9VimE",
	"bazos":    "9CeKiZRYNq+jp4Rfdpmo",
	"creed":    "sfkSbKxwPkIOwrPCafVA",
	"ellie":    "niGXKAPkbiYOS1Nl14Yr",
	"covid":    "Lh0KsODue3rowMbNRu3r",
	"admin":    "ndJS1O1DLVtbiCIf5zlj",
	"aaron":    "ONflxlxC2VQSGdaao6hc",
	"smith":    "ua7KDrWDHZ7EtEAXdiv8",
	"john":     "0YhgB9/eT8qdQXTwCIKj",
	"paul":     "5rDqPuX7YfxSjkVIriGR",
	"sharon":   "5THgc3SITCCTyvMNqwKn",
	"manager":  "p6FyfIviVNKAKgYH5LBP",
}
