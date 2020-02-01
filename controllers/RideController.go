package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"taxi/models"
	"taxi/shared/config"
	"taxi/shared/database"
	"time"

	"github.com/gin-gonic/gin"
)

type RideController struct {
}
type RideAcceptCustomerResponse struct {
	DriverDetails  models.Driver
	RideDetails    models.Ride
	Payload        string
}

type RideAcceptDriverResponse struct {
	Status        bool
	Message       string
	PassengerName string
	RideDetails   models.Ride
}
type GetRidesRequest struct {
	RideStatus int64
}
type RideListItem struct {
	PassengerName    string
	ServiceArea      string
	PickupLocation   string
	DropLocation     string
	RideDateTime     time.Time
	RideStartTime    time.Time
	RideEndTime      time.Time
	RideType         int64
	IsRideLater      bool
	ID               uint
	Distance         float64
	Duration         float64
	DurationReadable string
	BaseFare         float64
	DistanceFare     float64
	DurationFare     float64
	Tax              float64
	TotalFare        float64

	RideStatus int64
}

type RideDetail struct {
	PassengerName     string
	PassengerMobile   string
	PassengerDialCode int64

	DriverName         string
	DriverMobile       string
	DriverDialCode     int64
	VehicleName        string
	VehicleNumber      string
	ServiceArea        string
	Currency           string
	PickupLocation     string
	DropLocation       string
	RideDateTime       time.Time
	RideStartTime      time.Time
	RideEndTime        time.Time
	RideType           int64
	IsRideLater        bool
	ID                 uint
	Distance           float64
	Duration           float64
	DurationReadable   string
	BaseFare           float64
	KmInBaseFare       float64
	DurationInBaseFare float64
	BaseDistanceFare   float64
	BaseDurationFare   float64

	PassengerRating float64
	DriverRating    float64
	DriverReview    string
	PassengerReview string

	DistanceFare  float64
	DurationFare  float64
	Tax           float64
	TaxPercentage float64
	TotalFare     float64
	RideStatus    int64
}

func (r *RideController) GetRides(c *gin.Context) {
	var data GetRidesRequest
	var list []RideListItem
	c.BindJSON(&data)

	if data.RideStatus == -1 {
		database.Db.Raw("SELECT rides.*,locations.name as service_area,passengers.name as passenger_name FROM rides INNER JOIN locations ON rides.location_id = locations.id INNER JOIN passengers ON passengers.id = rides.passenger_id ORDER BY rides.created_at DESC").Scan(&list)
	} else {
		database.Db.Raw("SELECT rides.*,locations.name as service_area,passengers.name as passenger_name FROM rides INNER JOIN locations ON rides.location_id = locations.id INNER JOIN passengers ON passengers.id = rides.passenger_id WHERE rides.ride_status = " + strconv.Itoa(int(data.RideStatus)) + " ORDER BY rides.created_at DESC").Scan(&list)
	}
	c.JSON(http.StatusOK, list)
}

func (r *RideController) GetRidesForPassenger(c *gin.Context) {
	var list []RideListItem
	database.Db.Raw("SELECT rides.*,locations.name as service_area,passengers.name as passenger_name FROM rides INNER JOIN locations ON rides.location_id = locations.id INNER JOIN passengers ON passengers.id = " + c.Param("passengerId") + " WHERE rides.passenger_id = " + c.Param("passengerId") + "  ORDER BY rides.created_at DESC").Scan(&list)
	c.JSON(http.StatusOK, list)
}

func (r *RideController) GetRideDetail(c *gin.Context) {
	var detail RideDetail
	var ride models.Ride
	database.Db.Where("id=?", c.Param("rideId")).First(&ride)
	if ride.ZoneFareID == 0 {
		if ride.RideStatus == 4 {
			database.Db.Raw("SELECT rides.*,locations.name as service_area,locations.currency,fares.base_fare_distance as km_in_base_fare,fares.base_fare_duration as duration_in_base_fare,fares.base_fare,fares.duration_fare as base_duration_fare,fares.distance_fare as base_distance_fare,fares.tax as tax_percentage,passengers.name as passenger_name,passengers.dial_code as passenger_dial_code,passengers.mobile_number as passenger_mobile,drivers.name as driver_name,drivers.dial_code as driver_dial_code,drivers.mobile_number as driver_mobile,vehicles.name as vehicle_name,vehicles.vehicle_number FROM rides INNER JOIN vehicles ON vehicles.id = rides.vehicle_id INNER JOIN locations ON rides.location_id = locations.id INNER JOIN fares ON fares.id = rides.fare_id INNER JOIN drivers ON drivers.id = rides.driver_id INNER JOIN passengers ON passengers.id = rides.passenger_id WHERE rides.id = " + c.Param("rideId")).Scan(&detail)
		} else {
			database.Db.Raw("SELECT rides.*,locations.name as service_area,locations.currency,passengers.name as passenger_name,passengers.dial_code as passenger_dial_code,passengers.mobile_number as passenger_mobile FROM rides INNER JOIN locations ON rides.location_id = locations.id INNER JOIN passengers ON passengers.id = rides.passenger_id WHERE rides.id = " + c.Param("rideId")).Scan(&detail)
		}
	} else {
		if ride.RideStatus == 4 {
			database.Db.Raw("SELECT rides.*,locations.name as service_area,locations.currency,zone_fares.base_fare_distance as km_in_base_fare,zone_fares.base_fare_duration as duration_in_base_fare,zone_fares.base_fare,zone_fares.duration_fare as base_duration_fare,zone_fares.distance_fare as base_distance_fare,zone_fares.tax as tax_percentage,passengers.name as passenger_name,passengers.dial_code as passenger_dial_code,passengers.mobile_number as passenger_mobile,drivers.name as driver_name,drivers.dial_code as driver_dial_code,drivers.mobile_number as driver_mobile,vehicles.name as vehicle_name,vehicles.vehicle_number FROM rides INNER JOIN vehicles ON vehicles.id = rides.vehicle_id INNER JOIN locations ON rides.location_id = locations.id INNER JOIN zone_fares ON zone_fares.id = rides.zone_fare_id INNER JOIN drivers ON drivers.id = rides.driver_id INNER JOIN passengers ON passengers.id = rides.passenger_id WHERE rides.id = " + c.Param("rideId")).Scan(&detail)
		} else {
			database.Db.Raw("SELECT rides.*,locations.name as service_area,locations.currency,passengers.name as passenger_name,passengers.dial_code as passenger_dial_code,passengers.mobile_number as passenger_mobile FROM rides INNER JOIN locations ON rides.location_id = locations.id INNER JOIN passengers ON passengers.id = rides.passenger_id WHERE rides.id = " + c.Param("rideId")).Scan(&detail)
		}
	}

	c.JSON(http.StatusOK, detail)
}

type RideAcceptRequest struct {
	RideID    uint `json:"ride_id"`
	VehicleID uint `json:"vehicle_id"`
}

func checkRideQueueOfDriver(driverId uint) RideAcceptDriverResponse {

	var driverResponse = RideAcceptDriverResponse{
		Status: false,
	}
	var ride models.Ride
	database.Db.Where("driver_id = ? AND ride_status = 7", driverId).First(&ride)
	if ride.ID != 0 {
		var passengerDetails models.Passenger
		database.Db.Where("id=?", ride.PassengerID).First(&passengerDetails)
		driverResponse.Status = true
		driverResponse.RideDetails = ride
		driverResponse.PassengerName = passengerDetails.Name
	}
	return driverResponse
}

func (r *RideController) CheckRideQueue(c *gin.Context) {
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var response = checkRideQueueOfDriver(userData.UserID)
	if response.Status {
		database.Db.Model(&models.Ride{}).Where("id = ?", response.RideDetails.ID).UpdateColumn("ride_status", 1)
		response.RideDetails.RideStatus = 1
	}
	c.JSON(http.StatusOK, response)

}

func (r *RideController) RideAccept(c *gin.Context) {

	var data RideAcceptRequest
	c.BindJSON(&data)

	var rideId = data.RideID
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var driverId = userData.UserID

	var driverResponse = RideAcceptDriverResponse{
		Status: false,
	}

	var ride models.Ride
	database.Db.Where("id=?", rideId).First(&ride)

	if ride.RideStatus == 0 {

		var driverDetails models.Driver
		database.Db.Where("id=?", driverId).First(&driverDetails)

		if driverDetails.IsRide {
			database.Db.Model(&ride).UpdateColumns(&models.Ride{ DriverID: driverId, RideStatus: 7})
		} else {
			database.Db.Model(&ride).UpdateColumns(&models.Ride{ DriverID: driverId, RideStatus: 1})
			database.Db.Model(&driverDetails).UpdateColumn("is_ride", true)
		}

		var passengerDetails models.Passenger

		database.Db.Where("id=?", ride.PassengerID).First(&passengerDetails)

		var passengerResponse = RideAcceptCustomerResponse{
			DriverDetails:  driverDetails,
			RideDetails:    ride,
		}

		var eventLog = models.RideEventLog{
			RideID:     ride.ID,
			RideStatus: ride.RideStatus,
			Message:    "Driver Assigned For Ride",
		}

		database.Db.Create(&eventLog)

		passengerData, err := json.Marshal(&passengerResponse)

		if err == nil {
			mqttController.Publish(fmt.Sprintf("passenger/%d/ride_accepted", ride.PassengerID), 2, string(passengerData))
		}

		driverResponse.Status = true
		driverResponse.RideDetails = ride
		driverResponse.PassengerName = passengerDetails.Name

	} else {
		driverResponse.Message = "Sorry ! This ride can't be accepted now it is either cancelled or taken by another driver."
	}

	c.JSON(http.StatusOK, driverResponse)

}

func (r *RideController) CheckOnRide(passengerId uint) {

	if passengerId != 0 {
		var ride models.Ride
		result := database.Db.Where("passenger_id = ? AND ride_status IN (1,2,3)", passengerId).First(&ride)
		if result.RowsAffected != 0 {
			var driverDetails models.Driver
			var passengerDetails models.Passenger
			database.Db.Where("id=?", ride.DriverID).First(&driverDetails)
			database.Db.Where("id=?", ride.PassengerID).First(&passengerDetails)
			var passengerResponse = RideAcceptCustomerResponse{
				DriverDetails:  driverDetails,
				RideDetails:    ride,
			}

			passengerData, err := json.Marshal(&passengerResponse)
			if err == nil {
				mqttController.Publish(fmt.Sprintf("passenger/%d/ride_accepted", ride.PassengerID), 2, string(passengerData))
			}

		}
	}

}

func (r *RideController) DriverArrived(c *gin.Context) {

	type responseFormat struct {
		Status  bool
		Message string
	}
	var response = responseFormat{Status: false}
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var rideID = c.Param("rideId")
	var ride models.Ride
	result := database.Db.Where("id = ? AND driver_id = ? AND ride_status = 1", rideID, userData.UserID).First(&ride)
	if result.RowsAffected == 1 {
		database.Db.Model(&ride).UpdateColumn("ride_status", 2)
		var eventLog = models.RideEventLog{
			RideID:     ride.ID,
			RideStatus: ride.RideStatus,
			Message:    "Driver Arrived At PickUp Location",
		}
		database.Db.Create(&eventLog)
		response.Status = true
		c.JSON(http.StatusOK, response)
		return
	} else {
		response.Message = "Ride status cannot be changed as arrived now."
		c.JSON(http.StatusOK, response)
		return
	}

}

func (r *RideController) StartTrip(c *gin.Context) {

	type responseFormat struct {
		Status  bool
		Message string
	}

	var response = responseFormat{Status: false}
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var rideID = c.Param("rideId")
	var ride models.Ride
	result := database.Db.Where("id = ? AND driver_id = ? AND ride_status = 2", rideID, userData.UserID).First(&ride)
	if result.RowsAffected == 1 {
		database.Db.Model(&ride).UpdateColumns(&models.Ride{RideStatus: 3, RideStartTime: time.Now()})
		var eventLog = models.RideEventLog{
			RideID:     ride.ID,
			RideStatus: ride.RideStatus,
			Message:    "Ride Started",
		}
		database.Db.Create(&eventLog)
		response.Status = true
		c.JSON(http.StatusOK, response)
		return
	} else {
		response.Message = "Ride status cannot be changed as arrived now."
		c.JSON(http.StatusOK, response)
		return
	}

}

type responseFormat struct {
	Status          bool
	Message         string
	RideDetails     models.Ride
	BaseFareDetails models.Fare
	Currency        string
}

func (r *RideController) GetRideDetailsForMobile(c *gin.Context) {

	/*var response = responseFormat{Status: false}
	var rideID = c.Param("rideId")
	var ride models.Ride
	result := database.Db.Where("id = ? ", rideID).First(&ride)
	if result.RowsAffected == 1 {
		response.RideDetails = ride
		var location models.Location
		database.Db.Where("id = ?", ride.LocationID).First(&location)
		response.Currency = location.Currency
		if ride.ZoneFareID == 0 {
			var fare models.Fare
			database.Db.Where("id = ?", ride.FareID).First(&fare)

			response.BaseFareDetails = fare
			response.Status = true
		} else {
			var fare models.ZoneFare
			database.Db.Where("id = ?", ride.ZoneFareID).First(&fare)
			response.BaseFareDetails = models.Fare{BaseFare: fare.BaseFare, BaseFareDistance: fare.BaseFareDistance, BaseFareDuration: fare.BaseFareDuration}
			response.Status = true
		}
	}
	c.JSON(http.StatusOK, response)
*/
}

func (r *RideController) GetRideTimeline(c *gin.Context) {

	var response []models.RideEventLog
	database.Db.Where("ride_id = ? ", c.Param("id")).Find(&response)
	c.JSON(http.StatusOK, response)

}

func (r *RideController) StopTrip(c *gin.Context) {


	var response = responseFormat{Status: false}
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var rideID = c.Param("rideId")
	var ride models.Ride
	result := database.Db.Where("id = ? AND driver_id = ? AND ride_status = 3", rideID, userData.UserID).First(&ride)

	if result.RowsAffected == 1 {
		database.Db.Model(&ride).UpdateColumn("ride_status", 4)
		var eventLog = models.RideEventLog{
			RideID:     ride.ID,
			RideStatus: ride.RideStatus,
			Message:    "Ride Completed",
		}
		database.Db.Create(&eventLog)

		type dist struct {
			Distance float64
		}
		var rideDistance dist
		database.Db.Raw("UPDATE rides SET distance = RideDistance.distance  FROM (SELECT (round( CAST(float8 (st_Length(ST_MakeLine(RideLocations.latlng)::geography)/1000) as numeric), 2)+0.01)::float as distance FROM (SELECT latlng,ride_id FROM ride_locations Where ride_id =" + rideID + " ORDER BY time ASC) as RideLocations GROUP by RideLocations.ride_id ) as RideDistance").Scan(&rideDistance)
		database.Db.Where("id = ? ", rideID).First(&ride)

		var endTime = time.Now()
		var diff = endTime.Sub(ride.RideStartTime)
		var duration = diff.Minutes()
		duration = math.Ceil(duration*100) / 100
		if ride.ZoneFareID == 0 {
			var fare models.Fare
			fareResult := database.Db.Where("id = ?", ride.FareID).First(&fare)
			if fareResult.RowsAffected != 0 {
				totalFare := fare.BaseFare
				var distanceFare, durationFare float64
				if ride.Distance > fare.BaseFareDistance {
					distanceFare = (ride.Distance - fare.BaseFareDistance) * fare.DistanceFare
					distanceFare = math.Ceil(distanceFare*100) / 100
				}
				if duration > fare.BaseFareDuration {
					durationFare = (duration - fare.BaseFareDuration) * fare.DurationFare
					durationFare = math.Ceil(durationFare*100) / 100
				}
				totalFare = totalFare + distanceFare + durationFare
				var tax = (fare.Tax / 100) * totalFare
				tax = math.Ceil(tax*100) / 100
				totalFare += tax
				totalFare = math.Ceil(totalFare*100) / 100
				database.Db.Model(&ride).UpdateColumns(&models.Ride{
					Duration:         duration,
					RideEndTime:      endTime,
					DistanceFare:     distanceFare,
					DurationFare:     durationFare,
					Tax:              tax,
					TotalFare:        totalFare,
					DurationReadable: diff.String(),
				})
				database.Db.Model(&models.Driver{}).Where("id = ? ", ride.DriverID).UpdateColumns(&models.Driver{IsRide:false,IsOnline:true})
				var location models.Operator
				database.Db.Where("id = ?", ride.OperatorID).First(&location)
				response.RideDetails = ride
				response.Currency = location.Currency
				response.BaseFareDetails = fare
				response.Status = true

				//paytmpg ends here
				data, err := json.Marshal(&response)
				if err == nil {
					mqttController.Publish(fmt.Sprintf("passenger/%d/ride_invoice", ride.PassengerID), 0, string(data))
				}

				c.JSON(http.StatusOK, response)
			}
		} else {
			var fare models.ZoneFare
			fareResult := database.Db.Where("id = ?", ride.ZoneFareID).First(&fare)
			if fareResult.RowsAffected != 0 {
				totalFare := fare.BaseFare
				var distanceFare, durationFare float64
				if ride.Distance > fare.BaseFareDistance {
					distanceFare = (ride.Distance - fare.BaseFareDistance) * fare.DistanceFare
					distanceFare = math.Ceil(distanceFare*100) / 100
				}
				if duration > fare.BaseFareDuration {
					durationFare = (duration - fare.BaseFareDuration) * fare.DurationFare
					durationFare = math.Ceil(durationFare*100) / 100

				}
				totalFare = totalFare + distanceFare + durationFare
				var tax = (fare.Tax / 100) * totalFare
				tax = math.Ceil(tax*100) / 100
				totalFare += tax
				totalFare = math.Ceil(totalFare*100) / 100
				database.Db.Model(&ride).UpdateColumns(&models.Ride{
					Duration:         duration,
					RideEndTime:      endTime,
					DistanceFare:     distanceFare,
					DurationFare:     durationFare,
					Tax:              tax,
					TotalFare:        totalFare,
					DurationReadable: diff.String(),
				})
				database.Db.Model(&models.Driver{}).Where("id = ? ", ride.DriverID).UpdateColumns(&models.Driver{IsRide:false,IsOnline:true})
				var location models.Operator
				database.Db.Where("id = ?", ride.OperatorID).First(&location)
				response.RideDetails = ride
				response.Currency = location.Currency
				response.BaseFareDetails = models.Fare{BaseFare: fare.BaseFare, BaseFareDistance: fare.BaseFareDistance, BaseFareDuration: fare.BaseFareDuration}
				response.Status = true

				data, err := json.Marshal(&response)
				if err == nil {
					mqttController.Publish(fmt.Sprintf("passenger/%d/ride_invoice", ride.PassengerID), 0, string(data))
				}
				c.JSON(http.StatusOK, response)
			}
		}

		return
	} else {

		response.Message = "Ride status cannot be changed as arrived now."
		c.JSON(http.StatusOK, response)
		return

	}

}

func (r *RideController) GetRideLocations(c *gin.Context) {
	type latlng struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
	var rideLocations []latlng
	database.Db.Raw("SELECT ST_X(latlng) as lat, ST_Y(latlng) as lng,ride_id FROM ride_locations Where ride_id = " + c.Param("rideId") + " ORDER BY time ASC").Scan(&rideLocations)
	c.JSON(http.StatusOK, rideLocations)
}

type rideLocationItem struct {
	Latitude  float64
	Longitude float64
	UnixTime  int64
}

type rideLocationUpdateRequest struct {
	RideID    int64
	Locations []rideLocationItem
}
type rideLocationResponse struct {
	Status bool
}

func (r *RideController) UpdateRideLocations(c *gin.Context) {
	var locationUpdateRequest rideLocationUpdateRequest
	c.BindJSON(&locationUpdateRequest)
	var response = rideLocationResponse{Status: false}
	var query = "INSERT INTO ride_locations (ride_id,time,latlng) VALUES "

	if locationUpdateRequest.RideID != 0 {
		for i, location := range locationUpdateRequest.Locations {
			var prefix = ""
			if i != 0 {
				prefix = ","
			}
			query += (prefix + "(" + strconv.Itoa(int(locationUpdateRequest.RideID)) + ",'" + (time.Unix(location.UnixTime, 0).Format(time.RFC3339)) + "',ST_GeometryFromText('POINT(" + fmt.Sprintf("%f", location.Latitude) + " " + fmt.Sprintf("%f", location.Longitude) + ")'))")
		}
		result := database.Db.Exec(query)
		fmt.Println(query)
		if result.Error == nil {
			response.Status = true
		} else {
			fmt.Println(result.Error.Error())
		}
	} else {
		fmt.Println("ride id not valid")
	}
	c.JSON(http.StatusOK, response)
}

type rateRequest struct {
	RideID int64
	Review string
	Rating float64
}

func (r *RideController) RateDriver(c *gin.Context) {
	var response = rideLocationResponse{Status: false}
	var ride models.Ride
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var data rateRequest
	c.BindJSON(&data)
	database.Db.Where("id = ? AND passenger_id = ?", data.RideID, userData.UserID).First(&ride)
	if ride.ID != 0 {
		result := database.Db.Model(&ride).UpdateColumns(&models.Ride{DriverRating: data.Rating, PassengerReview: data.Review})
		if result.Error == nil {
			response.Status = true
		}
	}
	c.JSON(http.StatusOK, response)
}

func (r *RideController) RatePassenger(c *gin.Context) {
	var response = rideLocationResponse{Status: false}
	var ride models.Ride
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var data rateRequest
	c.BindJSON(&data)
	database.Db.Where("id = ? AND driver_id = ?", data.RideID, userData.UserID).First(&ride)
	if ride.ID != 0 {
		result := database.Db.Model(&ride).UpdateColumns(&models.Ride{PassengerRating: data.Rating, DriverReview: data.Review})
		if result.Error == nil {
			response.Status = true
		}
	}
	c.JSON(http.StatusOK, response)
}
