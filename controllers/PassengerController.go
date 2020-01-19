package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"taxi/models"
	"taxi/shared/config"
	"taxi/shared/database"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PassengerController struct {
}

type sendOtpRequest struct {
	DialCode     int64
	CountryCode  string
	MobileNumber string
}

type sendOtpResponse struct {
	Status  bool
	Message string
	IsNew   bool
}
type verifyOtpResponse struct {
	Status     bool
	Message    string
	UserDetail models.Passenger
}

type verifyOtpRequest struct {
	DialCode     int64
	CountryCode  string
	MobileNumber string
	Name         string
	Otp          string
}
type addBasicInfoRequest struct {
	UserID int64
	Name   string
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func (a *PassengerController) SendOtp(c *gin.Context) {
	var data sendOtpRequest
	var response = sendOtpResponse{Status: false, IsNew: false}
	c.BindJSON(&data)
	if (data.DialCode) == 0 {
		response.Message = "DialCode is required"
		c.JSON(http.StatusOK, response)
	} else if len(data.CountryCode) == 0 {
		response.Message = "CountryCode is required"
		c.JSON(http.StatusOK, response)
	} else if len(data.MobileNumber) < 6 {
		response.Message = "MobileNumber is required"
		c.JSON(http.StatusOK, response)
	} else {
		//Validation Success and so send the otp
		var customer models.Passenger
		database.Db.Where("dial_code = ? AND country_code = ? AND mobile_number = ?", data.DialCode, data.CountryCode, data.MobileNumber).First(&customer)
		if customer.ID == 0 {
			response.IsNew = true
		} else if !customer.IsActive {
			response.Message = "Sorry! Your account is blocked by administrator."
			c.JSON(http.StatusOK, response)
			return
		}
		var pin = random(1000, 9999)
		var newOtp = models.Otp{
			DialCode:     data.DialCode,
			CountryCode:  data.CountryCode,
			MobileNumber: data.MobileNumber,
			Otp:          strconv.Itoa(pin),
		}
		database.Db.Model(&models.Otp{}).Where("dial_code = ? AND country_code = ? AND mobile_number = ? AND is_used = ?", data.DialCode, data.CountryCode, data.MobileNumber, false).UpdateColumn("is_used", true)
		var dbResponse = database.Db.Create(&newOtp)
		if dbResponse.Error == nil {
			response.Status = true
			response.Message = fmt.Sprintf("Sms gateway is disabled in testing environment.So Use %s as Otp to verify your mobile number.", strconv.Itoa(pin))
			c.JSON(http.StatusOK, response)
		} else {
			response.Message = dbResponse.Error.Error()
			c.JSON(http.StatusOK, response)
		}
	}
}

func (a *PassengerController) VerifyOtp(c *gin.Context) {
	var data verifyOtpRequest
	var response = verifyOtpResponse{Status: false}
	c.BindJSON(&data)
	if (data.DialCode) == 0 {
		response.Message = "DialCode is required"
		c.JSON(http.StatusOK, response)
	} else if len(data.CountryCode) == 0 {
		response.Message = "CountryCode is required"
		c.JSON(http.StatusOK, response)
	} else if len(data.MobileNumber) < 6 {
		response.Message = "MobileNumber is required"
		c.JSON(http.StatusOK, response)
	} else if len(data.Otp) != 4 {
		response.Message = "Otp must contain 4 digits"
		c.JSON(http.StatusOK, response)
	} else {
		var otpDetails models.Otp
		database.Db.Where("dial_code = ? AND country_code = ? AND mobile_number = ? AND is_used = ?", data.DialCode, data.CountryCode, data.MobileNumber, false).First(&otpDetails)
		if data.Otp == otpDetails.Otp {

			var customer models.Passenger
			var customerCount = 0
			database.Db.Where("dial_code = ? AND country_code = ? AND mobile_number = ?", data.DialCode, data.CountryCode, data.MobileNumber).First(&customer).Count(&customerCount)
			if customerCount == 0 {
				if len(strings.TrimSpace(data.Name)) >= 3 {
					database.Db.Model(&otpDetails).UpdateColumn("is_used", true)
					customer.MobileNumber = data.MobileNumber
					customer.DialCode = data.DialCode
					customer.CountryCode = data.CountryCode
					customer.Name = data.Name
					customer.IsActive = true
					database.Db.Create(&customer)
					token := jwt.NewWithClaims(jwt.SigningMethodHS256, config.JwtClaims{
						customer.ID,
						"passenger",
						jwt.StandardClaims{
							ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
							Issuer:    "onride",
						},
					})
					tokenString, err := token.SignedString(config.JwtSecretKey)
					if err != nil {
						response.Message = err.Error()
						response.Status = false
					} else {
						database.Db.Model(&customer).UpdateColumn("auth_token", tokenString)
						response.UserDetail = customer
						response.Status = true
						response.Message = "Otp verified successfully"
						c.JSON(http.StatusOK, response)
					}
				} else {
					response.Message = "Name must contain 3 characters"
					c.JSON(http.StatusOK, response)
				}

			} else {
				database.Db.Model(&otpDetails).UpdateColumn("is_used", true)

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"user_id":   customer.ID,
					"user_type": "passenger",
				})
				tokenString, err := token.SignedString(config.JwtSecretKey)
				if err != nil {
					response.Message = err.Error()
					response.Status = false
				} else {
					database.Db.Model(&customer).UpdateColumn("auth_token", tokenString)
					response.UserDetail = customer
					response.Status = true
					response.Message = "Otp verified successfully"
					c.JSON(http.StatusOK, response)
				}

			}
		} else {
			response.Message = "Invalid Otp"
			c.JSON(http.StatusOK, response)
		}

	}
}

func (a *PassengerController) AddCustomerBasicDetails(c *gin.Context) {
	var data addBasicInfoRequest
	var response = sendOtpResponse{Status: false}
	c.BindJSON(&data)
	if len(data.Name) < 3 {
		response.Message = "Name is required"
		c.JSON(http.StatusOK, response)
	} else {
		var customer models.Passenger
		var customerCount = 0
		database.Db.Where("user_id = ? AND is_active = ?", data.UserID, true).First(&customer).Count(&customerCount)
		if customerCount == 0 {
			response.Message = "Customer not found"
			c.JSON(http.StatusOK, response)
		} else {
			response.Status = true
			response.Message = "Customer details updated"
			c.JSON(http.StatusOK, response)
		}
	}
}

func (a *PassengerController) CheckIsOnRide(c *gin.Context) {
	var ride models.Ride
	var response = rideBookingResponse{Status: false}
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	database.Db.Where("ride_status IN (0,1,2,3) AND passenger_id = ?", userData.UserID).First(&ride)
	if ride.ID != 0 {
		response.Status = true
		response.RideDetails = ride
	}
	c.JSON(http.StatusOK, response)
}

type GetNearByDriversRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type NearByDriver struct {
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (a *PassengerController) GetNearByDrivers(c *gin.Context) {
	var data GetNearByDriversRequest
	var response []NearByDriver
	c.BindJSON(&data)
	database.Db.Raw("SELECT driver_vehicle_assignments.driver_id as id,ST_X(vehicles.latlng) AS longitude,ST_Y(vehicles.latlng) AS latitude,ST_Distance(vehicles.latlng, ref_geom) AS distance from vehicles INNER JOIN driver_vehicle_assignments ON driver_vehicle_assignments.vehicle_id = vehicles.id AND driver_vehicle_assignments.is_online = true AND driver_vehicle_assignments.is_ride = false CROSS JOIN (SELECT ST_MakePoint(" + fmt.Sprintf("%f", data.Latitude) + "," + fmt.Sprintf("%f", data.Longitude) + ")::geography AS ref_geom) AS r  WHERE ST_DWithin(vehicles.latlng, ref_geom, 5000)  ORDER BY ST_Distance(vehicles.latlng, ref_geom)").Scan(&response)
	c.JSON(http.StatusOK, response)
}

func (a *PassengerController) GetAllPassengers(c *gin.Context) {
	var list []models.Passenger
	database.Db.Select([]string{"id", "name", "image", "dial_code", "mobile_number", "is_active"}).Find(&list)
	c.JSON(http.StatusOK, list)
}
