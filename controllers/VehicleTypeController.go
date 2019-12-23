package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"taxi/models"
	"taxi/shared/database"
	"time"

	"github.com/gin-gonic/gin"
)

type VehicleTypeController struct {
}

type addVehicleTypeRequest struct {
	Name string
}

type addVehicleTypeResponse struct {
	Status  bool
	Message string
}

func (a *VehicleTypeController) GetVehicleTypes(c *gin.Context) {
	var list []models.VehicleType
	database.Db.Find(&list)
	c.JSON(http.StatusOK, list)
}
func (a *VehicleTypeController) GetVehicleTypesForCategory(c *gin.Context) {
	var list []models.VehicleType
	database.Db.Where("vehicle_category_id = ? ", c.Param("categoryId")).Find(&list)
	c.JSON(http.StatusOK, list)
}

func (a *VehicleTypeController) GetVehicleTypeWithID(c *gin.Context) {
	var vehicleType models.VehicleType
	database.Db.Where("id = ?", c.Param("vehicleTypeId")).First(&vehicleType)
	c.JSON(http.StatusOK, vehicleType)
}

func (a *VehicleTypeController) GetVehicleTypeCategories(c *gin.Context) {
	type vehicleTypeCategoryWithCount struct {
		ID                uint
		Name              string
		TotalVehicleTypes int64
		IsActive          bool
	}
	var list []vehicleTypeCategoryWithCount
	database.Db.Raw("SELECT (SELECT COUNT(*) as total_vehicle_types FROM vehicle_types  WHERE vehicle_types.vehicle_category_id = vehicle_categories.id),* FROM vehicle_categories").Scan(&list)
	c.JSON(http.StatusOK, list)
}

func (a *VehicleTypeController) GetVehicleTypesWithFare(c *gin.Context) {
	type vehicleTypeWithCategory struct {
		ID                  uint
		Name                string
		Image               string
		VehicleCategoryID   uint
		VehicleCategoryName string
		Description         string
		ImageActive         string
	}
	var list []vehicleTypeWithCategory
	fmt.Println("GetVehicleTypesWithFare : called")
	database.Db.Raw("SELECT vehicle_types.*,vehicle_categories.name as vehicle_category_name FROM vehicle_types INNER JOIN vehicle_categories ON vehicle_categories.id = vehicle_types.vehicle_category_id  INNER JOIN fares ON fares.vehicle_type_id = vehicle_types.id AND fares.is_active = true AND fares.location_id = ? WHERE vehicle_types.is_active = true", c.Param("locationId")).Scan(&list)
	c.JSON(http.StatusOK, list)
}

func (a *VehicleTypeController) GetUnAssignedVehicleTypes(c *gin.Context) {
	type vehicleTypeWithCategory struct {
		ID                  uint
		Name                string
		Image               string
		VehicleCategoryID   uint
		VehicleCategoryName string
		Description         string
		ImageActive         string
	}
	var list []vehicleTypeWithCategory
	database.Db.Raw("SELECT vehicle_types.*,vehicle_categories.name as vehicle_category_name FROM vehicle_types INNER JOIN vehicle_categories ON vehicle_types.vehicle_category_id = vehicle_categories.id WHERE NOT EXISTS (SELECT * FROM fares WHERE fares.location_id = ? AND fares.vehicle_type_id = vehicle_types.id AND fares.is_active = true AND deleted_at IS NULL)", c.Param("locationId")).Scan(&list)
	c.JSON(http.StatusOK, list)
}
func (a *VehicleTypeController) GetUnAssignedVehicleTypesForZone(c *gin.Context) {
	type vehicleTypeWithCategory struct {
		ID                  uint
		Name                string
		Image               string
		VehicleCategoryID   uint
		VehicleCategoryName string
		Description         string
		ImageActive         string
	}
	var list []vehicleTypeWithCategory
	database.Db.Raw("SELECT vehicle_types.*,vehicle_categories.name as vehicle_category_name FROM vehicle_types INNER JOIN vehicle_categories ON vehicle_types.vehicle_category_id = vehicle_categories.id WHERE NOT EXISTS (SELECT * FROM zone_fares WHERE zone_fares.zone_id = ? AND zone_fares.vehicle_type_id = vehicle_types.id AND zone_fares.is_active = true AND deleted_at IS NULL)", c.Param("locationId")).Scan(&list)
	c.JSON(http.StatusOK, list)
}

func (a *VehicleTypeController) GetActiveVehicleTypes(c *gin.Context) {
	var list []models.VehicleType
	database.Db.Where("is_active = ?", true).Find(&list)
	c.JSON(http.StatusOK, list)
}

func (a *VehicleTypeController) GetActiveVehicleTypeCategories(c *gin.Context) {
	var list []models.VehicleCategory
	database.Db.Where("is_active = ?", true).Find(&list)
	c.JSON(http.StatusOK, list)
}

func (a *VehicleTypeController) AddNewVehicleTypeCategory(c *gin.Context) {
	var data models.VehicleCategory
	var response = addNewVehicleCategoryResponse{Status: false}
	c.BindJSON(&data)
	if len(data.Name) < 4 {
		response.Message = "Name should contain atleast 4 characters"
		c.JSON(http.StatusOK, response)
		return
	} else {
		data.IsActive = true
		result := database.Db.Create(&data)
		if result.Error == nil {
			response.Status = true
			response.Message = "Vehicle category created successfully!"
			c.JSON(http.StatusOK, response)
			return
		} else {
			response.Message = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}
	}

}

func (a *VehicleTypeController) EditVehicleTypeCategory(c *gin.Context) {
	var data models.VehicleCategory
	var response = addNewVehicleCategoryResponse{Status: false}
	c.BindJSON(&data)
	if len(data.Name) < 4 {
		response.Message = "Name should contain atleast 4 characters"
		c.JSON(http.StatusOK, response)
		return
	} else {
		data.IsActive = true
		result := database.Db.Model(&models.VehicleCategory{}).Where("id = ?", data.ID).UpdateColumn("name", data.Name)
		if result.Error == nil {
			response.Status = true
			response.Message = "Vehicle category updated successfully!"
			c.JSON(http.StatusOK, response)
			return
		} else {
			response.Message = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}
	}

}

func (a *VehicleTypeController) AddNewVehicleType(c *gin.Context) {
	var response = addVehicleTypeResponse{Status: false}
	name := c.PostForm("name")
	description := c.PostForm("description")
	fmt.Println("name : " + name)
	vehicleCategoryID, convertError := strconv.Atoi(c.PostForm("vehicle_category_id"))
	if len(description) < 10 || len(description) > 80 {
		response.Message = "Description length should be between 10 to 50 characters"
		c.JSON(http.StatusOK, response)
		fmt.Println(response)
		return
	} else if convertError != nil || vehicleCategoryID == 0 {
		response.Message = "Vehicle Category is required"
		c.JSON(http.StatusOK, response)
		fmt.Println(response)
		return
	} else {
		var vehicleCategoryItemCount = 0
		database.Db.Model(&models.VehicleType{}).Where("vehicle_category_id = ? AND is_active = true", vehicleCategoryID).Count(&vehicleCategoryItemCount)
		if vehicleCategoryItemCount < 3 {
			form, _ := c.MultipartForm()
			fmt.Println("file count = %d", len(form.File))
			// Source
			activeImage, err := c.FormFile("active_image")
			if err != nil {
				fmt.Println(err)
				response.Message = "Active Image is required"
				c.JSON(http.StatusOK, response)
				fmt.Println(response)
				return
			}

			inActiveImage, err1 := c.FormFile("inactive_image")
			if err1 != nil {
				response.Message = "Inactive Image is required"
				c.JSON(http.StatusOK, response)
				return
			}

			activeImageFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + activeImage.Filename
			inActiveImageFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + inActiveImage.Filename

			if err := c.SaveUploadedFile(activeImage, "public/vehicletype/"+activeImageFileName); err != nil {
				response.Message = fmt.Sprintf("upload file err: %s", err.Error())
				c.JSON(http.StatusOK, response)
				return
			} else {
				if err := c.SaveUploadedFile(inActiveImage, "public/vehicletype/"+inActiveImageFileName); err != nil {
					response.Message = fmt.Sprintf("upload file err: %s", err.Error())
					c.JSON(http.StatusOK, response)
					return
				} else {
					if len(name) < 3 {
						response.Message = "Name should contain atleast 3 characters"
						c.JSON(http.StatusOK, response)
						return
					} else {

						var newVehicleType = models.VehicleType{
							Name:              name,
							Image:             "public/vehicletype/" + inActiveImageFileName,
							ImageActive:       "public/vehicletype/" + activeImageFileName,
							VehicleCategoryID: uint(vehicleCategoryID),
							Description:       description,
							IsActive:          true,
						}
						database.Db.Create(&newVehicleType)
						response.Status = true
						response.Message = "Vehicle Type Added Successfully"
						c.JSON(http.StatusOK, response)
					}
				}
			}
		} else {
			response.Message = "Sorry ! Cannot Add more than 3 vehicle type in a Vehicle Category."
			c.JSON(http.StatusOK, response)
		}
	}

}

func (a *VehicleTypeController) EnableVehicleType(c *gin.Context) {
	var response struct {
		Status bool
	}
	var vehicleTypeId = c.Param("vehicleTypeId")
	res := database.Db.Model(&models.VehicleType{}).Where("id = ?", vehicleTypeId).UpdateColumn("is_active", true)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}

func (a *VehicleTypeController) DisableVehicleType(c *gin.Context) {
	var response struct {
		Status bool
	}
	var vehicleTypeId = c.Param("vehicleTypeId")
	res := database.Db.Model(&models.VehicleType{}).Where("id = ?", vehicleTypeId).UpdateColumn("is_active", false)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}

func (a *VehicleTypeController) EnableVehicleTypeCategory(c *gin.Context) {
	var response struct {
		Status bool
	}
	var categoryId = c.Param("categoryId")
	res := database.Db.Model(&models.VehicleCategory{}).Where("id = ?", categoryId).UpdateColumn("is_active", true)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}

func (a *VehicleTypeController) DisableVehicleTypeCategory(c *gin.Context) {
	var response struct {
		Status bool
	}
	var categoryId = c.Param("categoryId")
	res := database.Db.Model(&models.VehicleCategory{}).Where("id = ?", categoryId).UpdateColumn("is_active", false)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}

func (a *VehicleTypeController) EditVehicleType(c *gin.Context) {
	var response = addVehicleTypeResponse{Status: false}
	name := c.PostForm("name")
	description := c.PostForm("description")
	isNewImages := c.PostForm("is_new_images")
	fmt.Println("name : " + name)
	vehicleCategoryID, convertError := strconv.Atoi(c.PostForm("vehicle_category_id"))
	vehicleTypeID, convertError := strconv.Atoi(c.PostForm("id"))
	if len(description) < 10 || len(description) > 80 {
		response.Message = "Description length should be between 10 to 50 characters"
		c.JSON(http.StatusOK, response)
		fmt.Println(response)
		return
	} else if convertError != nil || vehicleCategoryID == 0 || vehicleTypeID == 0 {
		response.Message = "Vehicle Category or Type is required"
		c.JSON(http.StatusOK, response)
		fmt.Println(response)
		return
	} else {
		var vehicleType models.VehicleType
		result := database.Db.Where("id = ?", vehicleTypeID).First(&vehicleType)
		if result.RowsAffected != 0 {
			if isNewImages == "true" {
				form, _ := c.MultipartForm()
				fmt.Println("file count = %d", len(form.File))
				// Source
				activeImage, err := c.FormFile("active_image")
				if err != nil {
					fmt.Println(err)
					response.Message = "Active Image is required"
					c.JSON(http.StatusOK, response)
					fmt.Println(response)
					return
				}

				inActiveImage, err1 := c.FormFile("inactive_image")
				if err1 != nil {
					response.Message = "Inactive Image is required"
					c.JSON(http.StatusOK, response)
					return
				}

				activeImageFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + activeImage.Filename
				inActiveImageFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + inActiveImage.Filename

				err = c.SaveUploadedFile(activeImage, "public/vehicletype/"+activeImageFileName)
				err = c.SaveUploadedFile(inActiveImage, "public/vehicletype/"+inActiveImageFileName)
				if len(name) < 3 {
					response.Message = "Name should contain atleaast 3 characters"
					c.JSON(http.StatusOK, response)
					return
				} else {

					var updatedVehicleType = models.VehicleType{
						Name:              name,
						VehicleCategoryID: uint(vehicleCategoryID),
						Description:       description,
					}
					fmt.Println("Using new images")

					if err == nil {
						updatedVehicleType.Image = "public/vehicletype/" + inActiveImageFileName
						updatedVehicleType.ImageActive = "public/vehicletype/" + activeImageFileName
					} else {
						fmt.Println(err.Error())
					}
					database.Db.Model(&vehicleType).UpdateColumns(&updatedVehicleType)
					response.Status = true
					response.Message = "Vehicle Type Updated Successfully"
					c.JSON(http.StatusOK, response)
				}

			} else {
				var updatedVehicleType = models.VehicleType{
					Name:              name,
					VehicleCategoryID: uint(vehicleCategoryID),
					Description:       description,
				}
				fmt.Println("Using older images")

				database.Db.Model(&vehicleType).UpdateColumns(&updatedVehicleType)
				response.Status = true
				response.Message = "Vehicle Type Updated Successfully"
				c.JSON(http.StatusOK, response)
			}

		} else {
			response.Message = "Sorry ! Vehicle Type not found."
			c.JSON(http.StatusOK, response)
		}
	}

}
