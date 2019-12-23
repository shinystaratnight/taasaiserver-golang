package controllers

import (
	"fmt"
	"net/http"
	"taxi/models"
	"taxi/shared/database"

	"github.com/gin-gonic/gin"
)

type ZoneFareController struct {
}
type addNewZoneFareResponse struct {
	Status  bool
	Message string
}
type getZoneFareResponse struct {
	ID               int64   `json:"fare_id"`
	VehicleTypeID    int64   `json:"vehicle_type_id"`
	ZoneID           int64   `json:"zone_id"`
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

type zoneFareStatusChangeRequest struct {
	FareId int64
}

type zoneFareStatusChangeResponse struct {
	Status  bool
	Message string
}

func (a *ZoneFareController) AddNewZoneFare(c *gin.Context) {
	var data models.ZoneFare
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
		} else if data.ZoneID == 0 {
			response.Message = "Zone ID is required"
			c.JSON(http.StatusOK, response)
			return
		} else if data.DurationFare == 0 && data.DistanceFare == 0 {
			response.Message = "Either Duration Fare or Distance Fare is required"
			c.JSON(http.StatusOK, response)
			return
		} else {
			var count = 0
			database.Db.Model(&models.ZoneFare{}).Where("vehicle_type_id = ? AND zone_id = ? AND is_active = true", data.VehicleTypeID, data.ZoneID).Count(&count)
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

func (a *ZoneFareController) GetActiveZoneFare(c *gin.Context) {
	var list []getFareResponse
	database.Db.Raw("SELECT zone_fares.* ,vehicle_types.name as vehicle_type_name,vehicle_types.image_active as vehicle_type_image FROM zone_fares INNER JOIN vehicle_types ON zone_fares.vehicle_type_id = vehicle_types.id WHERE zone_fares.is_active=true AND zone_fares.deleted_at IS NULL AND zone_fares.zone_id = " + c.Param("zoneId")).Find(&list)
	c.JSON(http.StatusOK, list)
}

func (a *ZoneFareController) DisableZoneFare(c *gin.Context) {
	var data fareStatusChangeRequest
	c.BindJSON(&data)
	var response = fareStatusChangeResponse{Status: true, Message: "Zone Fare Disabled Successfully"}
	database.Db.Where("id = ?", data.FareId).Delete(&models.ZoneFare{})
	c.JSON(http.StatusOK, response)
}
