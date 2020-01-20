package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kelvins/geocoder"
	"math"
	"net/http"
	"taxi/models"
	"taxi/shared/config"
	"taxi/shared/database"
	"taxi/shared/googleMap"
	"time"

	"googlemaps.github.io/maps"

	"github.com/gin-gonic/gin"
)

type RideBookingController struct {
}

type estimateFareRequest struct {
	PickupLatitude  string `json:"pickup_latitude"`
	PickupLongitude string `json:"pickup_longitude"`
	DropLatitude    string `json:"drop_latitude"`
	DropLongitude   string `json:"drop_longitude"`
}

type estimateFareResponse struct {
	Status  bool
	Message string
	Data    []estimatedFare
}
type estimatedFare struct {
	ID                       uint    `json:"fare_id"`
	CategoryID               string  `json:"category_id"`
	VehicleTypeID            uint    `json:"vehicle_type_id"`
	LocationID               uint    `json:"location_id"`
	MinimumFare              float64 `json:"minimum_fare"`
	WaitingFare              float64 `json:"waiting_fare"`
	BaseFare                 float64 `json:"base_fare"`
	BaseFareDistance         float64 `json:"base_fare_distance"`
	BaseFareDuration         float64 `json:"base_fare_duration"`
	DurationFare             float64 `json:"duration_fare"`
	DistanceFare             float64 `json:"distance_fare"`
	Tax                      float64 `json:"tax"`
	EstimatedFare            float64 `json:"estimated_fare"`
	IsActive                 bool    `json:"is_active"`
	Currency                 string  `json:"currency"`
	LocationName             string  `json:"location_name"`
	CategoryName             string  `json:"category_name"`
	VehicleTypeName          string  `json:"vehicle_type_name"`
	VehicleTypeDesc          string  `json:"vehicle_type_desc"`
	VehicleTypeImage         string  `json:"vehicle_type_image"`
	VehicleTypeImageInactive string  `json:"vehicle_type_image_inactive"`
}

func (r *RideBookingController) GetEstimatedFare(c *gin.Context) {
	var data estimateFareRequest
	var response = estimateFareResponse{Status: false}
	c.BindJSON(&data)
	var intersectLocation models.Location
	database.Db.Where("is_active = true AND ST_Contains(polygon,ST_GeometryFromText('POINT(" + data.PickupLatitude + " " + data.PickupLongitude + ")'))").First(&intersectLocation)
	if intersectLocation.ID != 0 {

		distanceRequest := &maps.DistanceMatrixRequest{
			Origins:      []string{data.PickupLatitude + "," + data.PickupLongitude},
			Destinations: []string{data.DropLatitude + "," + data.DropLongitude},
			Mode:         maps.TravelModeDriving,
		}
		distanceMatrixResponse, distanceReqError := googleMap.Client.DistanceMatrix(context.Background(), distanceRequest)

		if distanceReqError == nil {
			if len(distanceMatrixResponse.Rows) > 0 && len(distanceMatrixResponse.Rows[0].Elements) > 0 {
				estimatedDistance := float64(distanceMatrixResponse.Rows[0].Elements[0].Distance.Meters / 1000.0)
				estimatedDuration := distanceMatrixResponse.Rows[0].Elements[0].Duration.Minutes()

				var fareList []estimatedFare
				fareResult := database.Db.Raw("SELECT fares.*,vehicle_categories.id as category_id,vehicle_categories.name as category_name ,locations.currency,locations.name as location_name,vehicle_types.description as vehicle_type_desc,vehicle_types.name as vehicle_type_name,vehicle_types.image as vehicle_type_image_inactive,vehicle_types.image_active as vehicle_type_image FROM fares INNER JOIN locations ON fares.location_id = locations.id AND locations.is_active = true INNER JOIN vehicle_types ON fares.vehicle_type_id = vehicle_types.id AND vehicle_types.is_active = true INNER JOIN vehicle_categories ON vehicle_types.vehicle_category_id = vehicle_categories.id AND vehicle_categories.is_active = true WHERE fares.is_active=true AND fares.deleted_at IS NULL AND fares.location_id = ?", intersectLocation.ID).Find(&fareList)
				if fareResult.RowsAffected != 0 {
					for index, fare := range fareList {
						totalFare := fare.BaseFare
						if estimatedDistance > fare.BaseFareDistance {
							totalFare += (estimatedDistance - fare.BaseFareDistance) * fare.DistanceFare
						}
						if estimatedDuration > fare.BaseFareDuration {
							totalFare += (estimatedDuration - fare.BaseFareDuration) * fare.DurationFare
						}
						totalFare += (fare.Tax / 100) * totalFare
						fareList[index].EstimatedFare = math.Ceil(totalFare*100) / 100
						fmt.Println("estimatef fare : %f", totalFare)

					}
					//check if location comes under a zone
					var intersectZoneLocation models.Zone
					database.Db.Where("is_active = true AND location_id = ?  AND ST_Contains(polygon,ST_GeometryFromText('POINT("+data.PickupLatitude+" "+data.PickupLongitude+")'))", intersectLocation.ID).First(&intersectZoneLocation)
					if intersectZoneLocation.ID != 0 {
						var zoneFareList []models.ZoneFare
						database.Db.Where("is_active = true AND deleted_at IS NULL AND zone_id = ?", intersectZoneLocation.ID).Find(&zoneFareList)
						if len(zoneFareList) != 0 {
							for _, fare := range zoneFareList {
								totalFare := fare.BaseFare
								if estimatedDistance > fare.BaseFareDistance {
									totalFare += (estimatedDistance - fare.BaseFareDistance) * fare.DistanceFare
								}
								if estimatedDuration > fare.BaseFareDuration {
									totalFare += (estimatedDuration - fare.BaseFareDuration) * fare.DurationFare
								}
								totalFare += (fare.Tax / 100) * totalFare
								for index, normalFare := range fareList {
									if normalFare.VehicleTypeID == fare.VehicleTypeID {
										fareList[index].EstimatedFare = math.Ceil(totalFare*100) / 100
										break
									}
								}
							}
						}
					}

					response.Status = true
					response.Message = "Service Available"
					response.Data = fareList

					c.JSON(http.StatusOK, response)
				} else {
					response.Message = "Sorry! No cars available at the moment."
					c.JSON(http.StatusOK, response)
				}
			}

		}
	} else {
		response.Message = "Sorry! Service not available at the pickup location specified."
		c.JSON(http.StatusOK, response)
	}

}

type rideBookingResponse struct {
	Status      bool
	Message     string
	RideDetails models.Ride
}

func (a *RideBookingController) GetBookingHistory(c *gin.Context) {
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var rideList []models.Ride
	database.Db.Where("passenger_id = ?", userData.UserID).Find(&rideList)
	c.JSON(http.StatusOK, rideList)

}

func (a *RideBookingController) GetDriverBookingHistory(c *gin.Context) {
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var rideList []models.Ride
	database.Db.Where("driver_id = ?", userData.UserID).Find(&rideList)
	c.JSON(http.StatusOK, rideList)

}

func (a *RideBookingController) BookRide(c *gin.Context) {
	var data models.Ride
	c.BindJSON(&data)
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)

	var response = rideBookingResponse{Status: false}
	if userData.UserID == 0 {
		response.Message = "UserID not found"
		c.JSON(http.StatusOK, response)
		return
	} else if userData.UserType != "passenger" {
		response.Message = "Sorry! Non passenger accounts cannot book rides."
		c.JSON(http.StatusOK, response)
		return
	} else if data.VehicleTypeID == 0 {
		response.Message = "VehicleTypeID is required."
		c.JSON(http.StatusOK, response)
		return
	} else if data.PickupLatitude == 0 {
		response.Message = "PickupLatitude is required"
		c.JSON(http.StatusOK, response)
		return
	} else if data.PickupLongitude == 0 {
		response.Message = "PickupLongitude is required"
		c.JSON(http.StatusOK, response)
		return
	} else if data.DropLatitude == 0 {
		response.Message = "DropLatitude is required"
		c.JSON(http.StatusOK, response)
		return
	} else if data.DropLongitude == 0 {
		response.Message = "DropLongitude is required"
		c.JSON(http.StatusOK, response)
		return
	} else {
		var intersectLocation models.Location
		database.Db.Where("is_active = true AND ST_Contains(polygon,ST_GeometryFromText('POINT(" + fmt.Sprintf("%f", data.PickupLatitude) + " " + fmt.Sprintf("%f", data.PickupLongitude) + ")'))").First(&intersectLocation)
		if intersectLocation.ID != 0 {
			data.PassengerID = userData.UserID
			data.LocationID = intersectLocation.ID
			var intersectZoneLocation models.Zone
			database.Db.Where("is_active = true AND location_id = ?  AND ST_Contains(polygon,ST_GeometryFromText('POINT("+fmt.Sprintf("%f", data.PickupLatitude)+" "+fmt.Sprintf("%f", data.PickupLongitude)+")'))", intersectLocation.ID).First(&intersectZoneLocation)
			if intersectZoneLocation.ID != 0 {
				var zoneFare models.ZoneFare
				database.Db.Where("is_active = true AND vehicle_type_id = ? AND deleted_at IS NULL AND zone_id = ?", data.VehicleTypeID, intersectZoneLocation.ID).Find(&zoneFare)
				if zoneFare.ID != 0 {
					data.ZoneID = intersectZoneLocation.ID
					data.ZoneFareID = zoneFare.ID
				}
			}
			var fare models.Fare
			database.Db.Where("is_active = true AND vehicle_type_id = ? AND deleted_at IS NULL AND location_id = ?", data.VehicleTypeID, intersectLocation.ID).Find(&fare)
			if fare.ID != 0 {
				data.FareID = fare.ID
				data.RideDateTime = time.Now()
				geocoder.ApiKey = "AIzaSyCmua_JtLFnNux2uKsi1sACWNm_qrSxlBo"
				pickupLocation := geocoder.Location{
					Latitude:  data.PickupLatitude,
					Longitude: data.PickupLongitude,
				}

				// Convert location (latitude, longitude) to a slice of addresses
				addresses, err := geocoder.GeocodingReverse(pickupLocation)

				if err == nil {
					// Usually, the first address returned from the API
					// is more detailed, so let's work with it
					address := addresses[0]
					data.PickupLocation = address.FormatAddress()
				}
				dropLocation := geocoder.Location{
					Latitude:  data.DropLatitude,
					Longitude: data.DropLongitude,
				}

				// Convert location (latitude, longitude) to a slice of addresses
				addresses, err = geocoder.GeocodingReverse(dropLocation)

				if err == nil {
					// Usually, the first address returned from the API
					// is more detailed, so let's work with it
					address := addresses[0]
					data.DropLocation = address.FormatAddress()
				}
				result := database.Db.Create(&data)
				if result.Error == nil {
					var eventLog = models.RideEventLog{
						RideID:     data.ID,
						RideStatus: data.RideStatus,
						Message:    "Ride Booking Accepted By Operator",
					}
					database.Db.Create(&eventLog)
					response.Message = "Ride booking success"
					response.Status = true
					response.RideDetails = data
					c.JSON(http.StatusOK, response)
					go func() {
						AssignDriverForRide(data)
					}()
					return
				} else {
					response.Message = result.Error.Error()
					c.JSON(http.StatusOK, response)
					return
				}
			} else {
				response.Message = "Sorry! Currently vehicle type you have chosen is not available."
				c.JSON(http.StatusOK, response)
				return
			}

		} else {
			response.Message = "Sorry! Location is outside our service area."
			c.JSON(http.StatusOK, response)
			return
		}

	}
}

type cancelRideRequest struct {
	RideID int64 `json:"ride_id"`
}
type cancelRideResponse struct {
	Status  bool
	Message string
}

func (r *RideBookingController) CancelRide(c *gin.Context) {
	var data cancelRideRequest
	var response = cancelRideResponse{Status: false}
	c.BindJSON(&data)
	if data.RideID == 0 {
		response.Message = "Ride Id is required"
		c.JSON(http.StatusOK, response)
		return
	} else {
		var ride models.Ride
		database.Db.Where("id = ?", data.RideID).First(&ride)
		if ride.ID == 0 {
			response.Message = "Ride details not found"
			c.JSON(http.StatusOK, response)
			return
		} else if ride.RideStatus == 6 {
			response.Message = "Ride already cancelled!"
			c.JSON(http.StatusOK, response)
			return
		} else if ride.RideStatus != 0 && ride.RideStatus != 1 && ride.RideStatus != 2 {
			response.Message = "Ride cannot be cancelled now!"
			c.JSON(http.StatusOK, response)
			return
		} else {
			database.Db.Model(&ride).UpdateColumn("ride_status", 6)
			var eventLog = models.RideEventLog{
				RideID:     ride.ID,
				RideStatus: ride.RideStatus,
				Message:    "Ride Cancelled",
			}
			database.Db.Create(&eventLog)
			database.Db.Model(&models.DriverVehicleAssignment{}).Where("driver_id = ? AND vehicle_id = ?", ride.DriverID, ride.VehicleID).UpdateColumn("is_ride", false)
			data, err := json.Marshal(&ride)
			if err == nil {
				mqttController.Publish(fmt.Sprintf("driver/%d/ride_cancelled", ride.DriverID), 2, string(data))
			} else {
				mqttController.Publish(fmt.Sprintf("driver/%d/ride_cancelled", ride.DriverID), 2, string(data))
			}
			response.Message = "Ride cancelled successfully!"
			response.Status = true
			c.JSON(http.StatusOK, response)
			return
		}
	}
}
