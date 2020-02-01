package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"taxi/models"
	"taxi/shared/config"
	"taxi/shared/database"

	jwt "github.com/dgrijalva/jwt-go"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
)

type MqttController struct {
}

var client MQTT.Client

func (m *MqttController) Connect() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID(config.MqttKey)
	opts.SetUsername(config.MqttKey)
	opts.SetPassword(config.MqttKey)
	opts.SetCleanSession(false)

	client = MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Println("\n\n\n Mqtt Connection Success !\n\n\n")
	}

}

func (m *MqttController) Publish(topic string, qos int, payload string) {
	token := client.Publish(topic, byte(qos), false, payload)
	if token.Wait() && token.Error() != nil {
		fmt.Println("\n\n\n topic : " + topic + " , " + token.Error().Error() + "!\n\n\n")

	} else {
		fmt.Println("\n\n\n" + topic + "Mqtt message Sent Success !\n\n\n")
	}
}

type mqttAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ClientID string `json:"client_id"`
}
type mqttAuthResponse struct {
	Result string `json:"result"`
}

type mqttWebHookRequest struct {
	ClientId string `json:"client_id"`
	Username string `json:"username"`
	Topic    string `json:"topic"`
	Payload  string `json:"payload"`
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.AbortWithStatus(code)
}

type LocationUpdateRequest struct {
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}

func (a *MqttController) WebHook(c *gin.Context) {
	var requestData mqttWebHookRequest
	c.BindJSON(&requestData)
	token, err := jwt.ParseWithClaims(requestData.Username, &config.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JwtSecretKey, nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*config.JwtClaims); ok && token.Valid {
			payloadString, payloadError := base64.StdEncoding.DecodeString(requestData.Payload)
			if payloadError == nil {
				if requestData.Topic == "locationUpdate" {
					var locationUpdateRequest LocationUpdateRequest
					err := json.Unmarshal(payloadString, &locationUpdateRequest)
					if err == nil {
						var newLocationUpdateResponse = database.Db.Exec("UPDATE drivers SET latlng = ST_GeometryFromText('POINT(" + fmt.Sprintf("%f", locationUpdateRequest.Latitude) + " " + fmt.Sprintf("%f", locationUpdateRequest.Longitude) + ")') where id = " + fmt.Sprintf("%d", claims.UserID))
						database.Db.Model(&models.Driver{}).Where("id = ? ", claims.UserID).UpdateColumn("is_online", true)
						if newLocationUpdateResponse.Error != nil {
							fmt.Println("latlng update error!")
						} else {
							fmt.Println("latlng update success!")
						}
					}
				} else if requestData.Topic == "check_ride" {
					//var passengerID = claims.UserID
					//var rideController = RideController{}
					//rideController.CheckOnRide(passengerID)

				}
			} else {
				fmt.Println("payload decode error")
			}
		}
	}

}

type authorizationResponse struct {
}

func (a *MqttController) ClientGone(c *gin.Context) {
	var requestData mqttWebHookRequest
	c.BindJSON(&requestData)
	clientData := strings.Split(requestData.ClientId, "#")
	if clientData[0] == "driver" {
		database.Db.Model(&models.Driver{}).Where("id = ?", clientData[1]).UpdateColumn("is_online", false)
	} else if clientData[0] == "passenger" {

	}
}

func (a *MqttController) HandleMqttAuthorization(c *gin.Context) {
	c.JSON(http.StatusOK, authorizationResponse{})
}

func (a *MqttController) HandleMqttAuth(c *gin.Context) {

	var data mqttAuthRequest
	c.BindJSON(&data)
	var username = data.Username

	if username == "" {
		fmt.Println("password empty")

		respondWithError(401, "API token required", c)
		return
	} else if username == config.MqttKey {
		c.JSON(http.StatusOK, mqttAuthResponse{Result: "ok"})
		return
	} else {
		token, err := jwt.ParseWithClaims(username, &config.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return config.JwtSecretKey, nil
		})
		if err == nil {
			if claims, ok := token.Claims.(*config.JwtClaims); ok && token.Valid {
				var count = 0
				fmt.Println("%+v", claims)
				fmt.Println("id", claims.UserID)

				if claims.UserType == "passenger" {
					database.Db.Model(&models.Passenger{}).Where("id = ? AND auth_token = ? AND is_active = true", claims.UserID, username).Count(&count)
				} else if claims.UserType == "driver" {
					database.Db.Model(&models.Driver{}).Where("id = ? AND auth_token = ? ", claims.UserID, username).Count(&count)
				} else if claims.UserType == "admin" {
					database.Db.Model(&models.Admin{}).Where("id = ? AND auth_token = ? AND is_active = true", claims.UserID, username).Count(&count)
				}
				if count == 0 {
					respondWithError(401, "Invalid API token", c)
					return
				} else {
					c.JSON(http.StatusOK, mqttAuthResponse{Result: "ok"})
				}
			} else {
				fmt.Println("cannot parse token")
				respondWithError(401, "Invalid API token", c)
				return
			}
		} else {
			fmt.Println("error : " + err.Error())
			respondWithError(401, "Invalid API token", c)
			return
		}
	}

}
