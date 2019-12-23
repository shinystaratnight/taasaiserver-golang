package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"taxi/models"
	"taxi/shared/database"
	"time"

	"github.com/gin-gonic/gin"
)

type VehicleController struct {
}
type getVehicleResponse struct {
	ID              int64  `json:"vehicle_id"`
	Name            string `json:"vehicle_name"`
	VehicleTypeId   int64  `json:"vehicle_type_id"`
	LocationId      int64  `json:"location_id"`
	Brand           string `json:"brand"`
	Model           string `json:"model"`
	Color           string `json:"color"`
	VehicleNumber   string `json:"vehicle_number"`
	SeatCapacity    int64  `json:"seat_capacity"`
	Image           string `json:"image"`
	IsActive        bool   `json:"is_active"`
	LocationName    string `json:"location_name"`
	VehicleTypeName string `json:"vehicle_type_name"`
}
type addNewVehicleCategoryResponse struct {
	Status  bool
	Message string
}

func (a *VehicleController) GetVehicles(c *gin.Context) {
	var list []getVehicleResponse
	database.Db.Raw("SELECT vehicle_types.name as vehicle_type_name,locations.name as location_name,vehicles.* FROM vehicles  INNER JOIN vehicle_types ON vehicle_types.id = vehicles.vehicle_type_id INNER JOIN locations ON locations.id = vehicles.location_id").Scan(&list)
	c.JSON(http.StatusOK, list)
}

func (a *VehicleController) GetVehiclesOfCompany(c *gin.Context) {
	var list []getVehicleResponse
	database.Db.Raw("SELECT vehicle_types.name as vehicle_type_name,locations.name as location_name,vehicles.* FROM  vehicles INNER JOIN company_location_assignments ON vehicles.company_location_assignment_id = company_location_assignments.id AND company_location_assignments.company_id = " + c.Param("companyId") + " INNER JOIN vehicle_types ON vehicle_types.id = vehicles.vehicle_type_id INNER JOIN locations ON locations.id = company_location_assignments.location_id").Scan(&list)
	c.JSON(http.StatusOK, list)
}

func (a *VehicleController) GetVehiclesForCompany(c *gin.Context) {
	var list []getVehicleResponse
	database.Db.Raw("SELECT vehicle_types.name as vehicle_type_name,locations.name as location_name,vehicles.* FROM vehicles INNER JOIN company_location_assignments ON vehicles.company_location_assignment_id = company_location_assignments.id AND company_location_assignments.company_id = " + c.Param("companyId") + "  INNER JOIN vehicle_types ON vehicle_types.id = vehicles.vehicle_type_id INNER JOIN locations ON locations.id = company_location_assignments.location_id").Scan(&list)
	c.JSON(http.StatusOK, list)
}

func (v *VehicleController) AddNewVehicle(c *gin.Context) {
	var data models.Vehicle
	var response = addNewDriverResponse{Status: false}
	data.Name = c.PostForm("name")
	data.Brand = c.PostForm("brand")
	data.VehicleModel = c.PostForm("model")
	data.Color = c.PostForm("color")
	data.VehicleNumber = c.PostForm("vehicle_number")
	var convertError error
	var locationID, vehicleTypeID int
	locationID, convertError = strconv.Atoi(c.PostForm("location_id"))
	vehicleTypeID, convertError = strconv.Atoi(c.PostForm("vehicle_type_id"))
	if convertError == nil {
		data.VehicleTypeID = uint(vehicleTypeID)
		data.CompanyLocationAssignmentID = uint(locationID)
		driverImage, err := c.FormFile("image")
		if err != nil {
			fmt.Println(err)
			response.Message = "Image is required"
			c.JSON(http.StatusOK, response)
			fmt.Println(response)
			return
		}
		if len(data.Brand) == 0 {
			response.Message = "Brand is required"
			c.JSON(http.StatusOK, response)
			return
		} else if len(data.VehicleModel) == 0 {
			response.Message = "Model is required"
			c.JSON(http.StatusOK, response)
			return
		} else if len(data.Name) < 3 {
			response.Message = "Vehicle name is required"
			c.JSON(http.StatusOK, response)
			return
		} else if len(data.VehicleNumber) < 3 {
			response.Message = "Vehicle Registration Number is required"
			c.JSON(http.StatusOK, response)
			return
		} else if data.CompanyLocationAssignmentID == 0 {
			response.Message = "Location Id is required"
			c.JSON(http.StatusOK, response)
			return
		} else if data.VehicleTypeID == 0 {
			response.Message = "Vehicle Type Id is required"
			c.JSON(http.StatusOK, response)
			return
		} else if data.CompanyLocationAssignmentID == 0 {
			response.Message = "CompanyLocationAssignmentID is required"
			c.JSON(http.StatusOK, response)
			return
		} else {
			imageFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + driverImage.Filename
			if err := c.SaveUploadedFile(driverImage, "public/vehicle/"+imageFileName); err != nil {
				response.Message = fmt.Sprintf("upload file err: %s", err.Error())
				c.JSON(http.StatusOK, response)
				return
			} else {
				data.VehicleNumber = strings.Trim(data.VehicleNumber, " ")
				data.Image = "public/vehicle/" + imageFileName
				data.IsActive = true
				result := database.Db.Create(&data)
				if result.Error == nil {
					response.Status = true
					response.Message = "Vehicle added successfully!"
					c.JSON(http.StatusOK, response)
					return
				} else {
					response.Message = result.Error.Error()
					c.JSON(http.StatusOK, response)
					return
				}

			}
		}
	} else {
		response.Message = "Location Id is required"
		c.JSON(http.StatusOK, response)
		return
	}
}

func (a *VehicleController) EnableVehicle(c *gin.Context) {
	var response struct {
		Status bool
	}
	var categoryId = c.Param("vehicleId")
	res := database.Db.Model(&models.Vehicle{}).Where("id = ?", categoryId).UpdateColumn("is_active", true)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}

func (a *VehicleController) DisableVehicle(c *gin.Context) {
	var response struct {
		Status bool
	}
	var categoryId = c.Param("vehicleId")
	res := database.Db.Model(&models.Vehicle{}).Where("id = ?", categoryId).UpdateColumn("is_active", false)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}
