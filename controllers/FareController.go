package controllers

import (
	"fmt"
	"net/http"
	"taxi/models"
	"taxi/shared/database"

	"github.com/gin-gonic/gin"
)

type FareController struct {
}
type addNewFareResponse struct {
	Status  bool
	Message string
}
type getFareResponse struct {
	ID               int64   `json:"fare_id"`
	VehicleTypeID    int64   `json:"vehicle_type_id"`
	LocationID       int64   `json:"location_id"`
	MinimumFare      float64 `json:"minimum_fare"`
	WaitingFare      float64 `json:"waiting_fare"`
	BaseFare         float64 `json:"base_fare"`
	BaseFareDistance float64 `json:"base_fare_distance"`
	BaseFareDuration float64 `json:"base_fare_duration"`
	DurationFare     float64 `json:"duration_fare"`
	DistanceFare     float64 `json:"distance_fare"`
	Tax              float64 `json:"tax"`
	IsActive         bool    `json:"is_active"`
	Currency         string  `json:"currency"`
	LocationName     string  `json:"location_name"`
	VehicleTypeName  string  `json:"vehicle_type_name"`
	VehicleTypeImage string  `json:"vehicle_type_image"`
}

func (a *FareController) GetActiveFare(c *gin.Context) {
	var list []getFareResponse
	database.Db.Raw("SELECT fares.* ,locations.currency,locations.name as location_name,vehicle_types.name as vehicle_type_name,vehicle_types.image_active as vehicle_type_image FROM fares INNER JOIN locations ON fares.location_id = locations.id INNER JOIN vehicle_types ON fares.vehicle_type_id = vehicle_types.id WHERE fares.is_active=true AND fares.deleted_at IS NULL").Find(&list)
	c.JSON(http.StatusOK, list)
}

func (a *FareController) GetActiveFareForLocation(c *gin.Context) {
	var list []getFareResponse
	database.Db.Raw("SELECT fares.* ,locations.currency,locations.name as location_name,vehicle_types.name as vehicle_type_name,vehicle_types.image_active as vehicle_type_image FROM fares INNER JOIN locations ON fares.location_id = locations.id INNER JOIN vehicle_types ON fares.vehicle_type_id = vehicle_types.id WHERE fares.is_active=true AND fares.deleted_at IS NULL  AND fares.location_id = ?", c.Param("locationId")).Find(&list)
	c.JSON(http.StatusOK, list)
}

type fareStatusChangeRequest struct {
	FareId int64
}

type fareStatusChangeResponse struct {
	Status  bool
	Message string
}

func (a *FareController) DisableFare(c *gin.Context) {
	var data fareStatusChangeRequest
	c.BindJSON(&data)
	var response = fareStatusChangeResponse{Status: true, Message: "Location Disabled Successfully"}
	database.Db.Where("id = ?", data.FareId).Delete(&models.Fare{})
	c.JSON(http.StatusOK, response)
}

func (a *FareController) AddNewFare(c *gin.Context) {
	var data models.Fare
	var response = addNewFareResponse{Status: false}
	fmt.Println("method called")
	err := c.BindJSON(&data)
	fmt.Printf("%+v\n", data)

	if err == nil {
		fmt.Println("data binded")
		fmt.Printf("%+v\n", data)
		if data.VehicleTypeID == 0 {
			response.Message = "Vehicle Type ID is required"
			c.JSON(http.StatusOK, response)
			return
		} else if data.OperatorID == 0 {
			response.Message = "Operator ID is required"
			c.JSON(http.StatusOK, response)
			return
		} else if data.DurationFare == 0 && data.DistanceFare == 0 {
			response.Message = "Either Duration Fare or Distance Fare is required"
			c.JSON(http.StatusOK, response)
			return
		} else {
			var count = 0
			database.Db.Model(&models.Fare{}).Where("vehicle_type_id = ? AND operator_id = ? AND is_active = true", data.VehicleTypeID, data.OperatorID).Count(&count)
			if count == 0 {
				data.IsActive = true
				result := database.Db.Create(&data)
				if result.Error != nil {
					response.Message = result.Error.Error()
					c.JSON(http.StatusOK, response)
					return
				} else {
					response.Status = true
					response.Message = "Fare created successfully"
					c.JSON(http.StatusOK, response)
					return
				}
			} else {
				response.Message = "The selected location already has fare for vehicle type."
				c.JSON(http.StatusOK, response)
				return
			}

		}
	} else {
		fmt.Println("data bind error")
		response.Message = err.Error()
		c.JSON(http.StatusOK, response)
		return
	}

}
