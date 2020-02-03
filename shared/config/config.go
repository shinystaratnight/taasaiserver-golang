package config

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
)

type JwtClaims struct {
	UserID   uint   `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}
//var SendGridKey = "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
var MqttKey = "CGDJqpVAGOREo6P1ddiXq0JkrADJK5e4.GwXx4CCKcqYO1LjN3Om6NkPFXWKfc3hx"
var GeoFenceClient, _ = redis.Dial("tcp", ":9851")
var JwtSecretKey = []byte("e50fe02b850e4d40817e9be34e147f686828a7e87fb84540a0c944a0fb0eed577f8266410485491c8025b7cb04e041ca")

var StripePublishableKey = "pk_test_IaUh6JLXuxbxVjieMBoIzuKi00BrNd1scQ"
var StripeSecretKey = "sk_test_NmqdkctVUxk7neGylJ81aRAI002PRJd9vu"

var tlsCert = ""
var tlsKey = ""