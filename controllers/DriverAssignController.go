package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"taxi/models"
	"taxi/shared/database"
	"time"
)

var mqttController = MqttController{}

type DriverAssignController struct {
}
type rideDriverDetail struct {
	DriverID uint    `json:"driver_id"`
	Distance float64 `json:"distance"`
}

func AssignDriverForRide(ride models.Ride) {

	var nearestDriver rideDriverDetail

	database.Db.Raw("SELECT driver_vehicle_assignments.driver_id as driver_id,ST_Distance(vehicles.latlng, ref_geom) AS distance from vehicles INNER JOIN driver_vehicle_assignments ON driver_vehicle_assignments.vehicle_id = vehicles.id AND driver_vehicle_assignments.is_online = true AND driver_vehicle_assignments.is_ride = false CROSS JOIN (SELECT ST_MakePoint(" + fmt.Sprintf("%f", ride.PickupLatitude) + "," + fmt.Sprintf("%f", ride.PickupLongitude) + ")::geography AS ref_geom) AS r  WHERE ST_DWithin(vehicles.latlng, ref_geom, 10000) AND vehicles.vehicle_type_id =" + strconv.Itoa(int(ride.VehicleTypeID)) + "  ORDER BY ST_Distance(vehicles.latlng, ref_geom) LIMIT 1").Scan(&nearestDriver)
	fmt.Println(nearestDriver.DriverID)
	fmt.Println(nearestDriver.Distance)

	if nearestDriver.DriverID != 0 {
		data, err := json.Marshal(&ride)
		if err == nil {
			mqttController.Publish(fmt.Sprintf("driver/%d/new_ride_request", nearestDriver.DriverID), 0, string(data))
		} else {
			mqttController.Publish(fmt.Sprintf("driver/%d/new_ride_request", nearestDriver.DriverID), 0, string(data))
		}
		var request = models.SentRideRequest{
			DriverID: nearestDriver.DriverID,
			RideID:   ride.ID,
			IsActive: true,
		}
		database.Db.Create(&request)
	}
	scheduleNextAssignment(ride.ID)

}

func scheduleNextAssignment(rideID uint) {
	timer := time.NewTimer(time.Second * 24)
	<-timer.C
	CheckDriverAssignmentForRide(rideID)
}

func CheckDriverAssignmentForRide(rideId uint) {
	fmt.Println("\n\n Scheduler Running for ride id = " + strconv.Itoa(int(rideId)) + "\n\n")
	var ride models.Ride
	database.Db.Where("id = ?", rideId).First(&ride)
	if ride.ID != 0 {
		if ride.RideStatus == 0 {
			var nearestDriver rideDriverDetail
			var previousRequest []models.SentRideRequest
			database.Db.Where("ride_id = ?", rideId).Find(&previousRequest)
			if len(previousRequest) <= 5 {

				var previousDriverList string = ""

				for i, req := range previousRequest {
					if i != 0 {
						previousDriverList += ","
					}
					previousDriverList += fmt.Sprintf("%d", req.DriverID)
				}

				if len(previousRequest) > 0 {
					database.Db.Raw("SELECT driver_vehicle_assignments.driver_id as driver_id,ST_Distance(vehicles.latlng, ref_geom) AS distance from vehicles INNER JOIN driver_vehicle_assignments ON driver_vehicle_assignments.vehicle_id = vehicles.id AND driver_vehicle_assignments.is_online = true AND driver_vehicle_assignments.is_ride = false  AND driver_vehicle_assignments.driver_id NOT IN (" + previousDriverList + ")   CROSS JOIN (SELECT ST_MakePoint(" + fmt.Sprintf("%f", ride.PickupLatitude) + "," + fmt.Sprintf("%f", ride.PickupLongitude) + ")::geography AS ref_geom) AS r  WHERE ST_DWithin(vehicles.latlng, ref_geom, 5000) AND vehicles.vehicle_type_id =" + strconv.Itoa(int(ride.VehicleTypeID)) + "  ORDER BY ST_Distance(vehicles.latlng, ref_geom) LIMIT 1").Scan(&nearestDriver)
				} else {
					database.Db.Raw("SELECT driver_vehicle_assignments.driver_id as driver_id,ST_Distance(vehicles.latlng, ref_geom) AS distance from vehicles INNER JOIN driver_vehicle_assignments ON driver_vehicle_assignments.vehicle_id = vehicles.id AND driver_vehicle_assignments.is_online = true AND driver_vehicle_assignments.is_ride = false  CROSS JOIN (SELECT ST_MakePoint(" + fmt.Sprintf("%f", ride.PickupLatitude) + "," + fmt.Sprintf("%f", ride.PickupLongitude) + ")::geography AS ref_geom) AS r  WHERE ST_DWithin(vehicles.latlng, ref_geom, 5000) AND vehicles.vehicle_type_id =" + strconv.Itoa(int(ride.VehicleTypeID)) + "  ORDER BY ST_Distance(vehicles.latlng, ref_geom) LIMIT 1").Scan(&nearestDriver)
				}

				fmt.Println(nearestDriver.DriverID)
				fmt.Println(nearestDriver.Distance)

				if nearestDriver.DriverID != 0 {
					data, err := json.Marshal(&ride)
					if err == nil {
						mqttController.Publish(fmt.Sprintf("driver/%d/new_ride_request", nearestDriver.DriverID), 0, string(data))
					} else {
						mqttController.Publish(fmt.Sprintf("driver/%d/new_ride_request", nearestDriver.DriverID), 0, string(data))
					}
					var request = models.SentRideRequest{
						DriverID: nearestDriver.DriverID,
						RideID:   ride.ID,
					}
					database.Db.Create(&request)
					scheduleNextAssignment(rideId)
				} else {
					checkDriversGoingToComplete(ride)
				}

			} else {

				database.Db.Model(&ride).UpdateColumn("ride_status", 5)
				data, _ := json.Marshal(&ride)
				mqttController.Publish(fmt.Sprintf("passenger/%d/driver_unavailable", ride.PassengerID), 2, string(data))

			}

		}
	}
}

func checkDriversGoingToComplete(ride models.Ride) {
	var nearestDriver rideDriverDetail
	var previousRequest []models.SentRideRequest
	database.Db.Where("ride_id = ?", ride.ID).Find(&previousRequest)
	var previousDriverList string = ""

	for i, req := range previousRequest {
		if i != 0 {
			previousDriverList += ","
		}
		previousDriverList += fmt.Sprintf("%d", req.DriverID)
	}

	database.Db.Raw("SELECT driver_vehicle_assignments.driver_id as driver_id,ST_Distance(vehicles.latlng, ref_geom) AS distance from vehicles INNER JOIN driver_vehicle_assignments ON driver_vehicle_assignments.vehicle_id = vehicles.id AND driver_vehicle_assignments.is_ride = true AND driver_vehicle_assignments.driver_id NOT IN (" + previousDriverList + ")  INNER JOIN rides ON rides.driver_id = driver_vehicle_assignments.driver_id AND rides.ride_status = 4  CROSS JOIN (SELECT ST_MakePoint(" + fmt.Sprintf("%f", ride.PickupLatitude) + "," + fmt.Sprintf("%f", ride.PickupLongitude) + ")::geography AS ref_geom) AS r  WHERE ST_DWithin((SELECT ST_MakePoint(rides.drop_latitude,rides.drop_longitude)::geography), ref_geom, 5000) AND vehicles.vehicle_type_id =" + strconv.Itoa(int(ride.VehicleTypeID)) + "  ORDER BY ST_Distance((SELECT ST_MakePoint(rides.drop_latitude,rides.drop_longitude)::geography), ref_geom) LIMIT 1").Scan(&nearestDriver)

	if nearestDriver.DriverID != 0 {
		data, err := json.Marshal(&ride)
		if err == nil {
			mqttController.Publish(fmt.Sprintf("driver/%d/new_ride_request", nearestDriver.DriverID), 0, string(data))
		} else {
			mqttController.Publish(fmt.Sprintf("driver/%d/new_ride_request", nearestDriver.DriverID), 0, string(data))
		}
		var request = models.SentRideRequest{
			DriverID: nearestDriver.DriverID,
			RideID:   ride.ID,
		}
		database.Db.Create(&request)
		timer := time.NewTimer(time.Second * 24)
		<-timer.C
		checkDriversGoingToComplete(ride)
	} else {
		database.Db.Model(&ride).UpdateColumn("ride_status", 5)
		data, _ := json.Marshal(&ride)
		mqttController.Publish(fmt.Sprintf("passenger/%d/driver_unavailable", ride.PassengerID), 2, string(data))
	}
}

type acceptRideRequest struct {
	RideID int64 `json:"ride_id"`
}

/*

0 - waiting,
1 - accepted,
2 - arrived,
3 - started,
4 - stopped,
5 - driver unavailable
6 - cancelled,
7 - Queued

*/
