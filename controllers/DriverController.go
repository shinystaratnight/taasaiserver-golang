package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"taxi/models"
	"taxi/shared/config"
	"taxi/shared/database"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type DriverController struct {
}

type addNewDriverResponse struct {
	Status  bool
	Message string
}
type getDriverResponse struct {
	ID            int64  `json:"driver_id"`
	Name          string `json:"driver_name"`
	LocationId    int64  `json:"location_id"`
	DialCode      string `json:"dial_code"`
	MobileNumber  string `json:"mobile_number"`
	LicenseNumber string `json:"license_number"`
	AuthToken     string `json:"auth_token"`
	Image         string `json:"image"`
	FcmID         string `json:"fcm_id"`
	IsActive      bool   `json:"is_active"`
	LocationName  string `json:"location_name"`
}

type addVehicleAssignmentResponse struct {
	Status  bool
	Message string
}

func (a *DriverController) GetVehicleAssignmentsForID(c *gin.Context) {
	type response struct {
		Name          string
		ID            uint
		Image         string
		VehicleNumber string
		IsActive      bool
	}
	var list []response
	database.Db.Raw("SELECT vehicles.name,driver_vehicle_assignments.is_active,vehicles.image,vehicles.vehicle_number,driver_vehicle_assignments.id as ID FROM vehicles INNER JOIN driver_vehicle_assignments ON driver_vehicle_assignments.driver_id = ? AND driver_vehicle_assignments.vehicle_id = vehicles.id ", c.Param("driverId")).Scan(&list)
	c.JSON(http.StatusOK, list)

}

func (a *DriverController) GetVehicleAssignments(c *gin.Context) {
	var list []models.Vehicle
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	database.Db.Raw("SELECT vehicles.* FROM vehicles INNER JOIN driver_vehicle_assignments ON driver_vehicle_assignments.driver_id = ? AND driver_vehicle_assignments.vehicle_id = vehicles.id ", userData.UserID).Scan(&list)
	c.JSON(http.StatusOK, list)

}

func (a *DriverController) EnableAssignment(c *gin.Context) {
	var response struct {
		Status bool
	}
	var assignmentId = c.Param("id")
	res := database.Db.Model(&models.DriverVehicleAssignment{}).Where("id = ?", assignmentId).UpdateColumn("is_active", true)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}

func (a *DriverController) DisableAssignment(c *gin.Context) {
	var response struct {
		Status bool
	}
	var assignmentId = c.Param("id")
	res := database.Db.Model(&models.DriverVehicleAssignment{}).Where("id = ?", assignmentId).UpdateColumn("is_active", false)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}

func (a *DriverController) AddNewVehicle(c *gin.Context) {
	var data models.DriverVehicleAssignment
	var response = addVehicleAssignmentResponse{Status: false}
	c.BindJSON(&data)
	if data.DriverID == 0 {
		response.Message = "Driver ID is required"
		c.JSON(http.StatusOK, response)
		return
	} else if data.VehicleID == 0 {
		response.Message = "Vehicle ID is required"
		c.JSON(http.StatusOK, response)
		return
	} else {
		var vehicle models.Vehicle
		var driver models.Driver
		database.Db.Where("id = ? AND is_active = true", data.DriverID).First(&driver)
		database.Db.Where("id = ? AND is_active = true", data.VehicleID).First(&vehicle)
		if vehicle.ID == 0 || !vehicle.IsActive {
			response.Message = "Vehicle is invalid or disabled"
			c.JSON(http.StatusOK, response)
			return
		} else if driver.ID == 0 || !driver.IsActive {
			response.Message = "Driver is invalid or disabled"
			c.JSON(http.StatusOK, response)
			return
		} else if driver.CompanyLocationAssignmentID != vehicle.CompanyLocationAssignmentID {
			response.Message = "Either driver & vehicle is not in the same location"
			c.JSON(http.StatusOK, response)
			return
		}
		var assignmentCount = 0
		database.Db.Model(&models.DriverVehicleAssignment{}).Where("driver_id = ? AND vehicle_id = ?", data.DriverID, data.VehicleID).Count(&assignmentCount)
		if assignmentCount == 0 {
			data.IsActive = true
			database.Db.Create(&data)
			response.Status = true
			response.Message = "Vehicle assigned successfully!"
			c.JSON(http.StatusOK, response)
		} else {
			response.Message = "Driver is already assigned with that vehicle."
			c.JSON(http.StatusOK, response)
			return
		}
	}
}

func (a *DriverController) GetDrivers(c *gin.Context) {
	var list []getDriverResponse
	database.Db.Raw("SELECT drivers.* ,locations.name as location_name FROM drivers INNER JOIN locations ON drivers.location_id = locations.id ").Find(&list)
	c.JSON(http.StatusOK, list)
}

func (a *DriverController) GetDriversForCompany(c *gin.Context) {
	var list []getDriverResponse
	database.Db.Raw("SELECT drivers.* ,locations.name as location_name FROM drivers INNER JOIN company_location_assignments ON drivers.company_location_assignment_id = company_location_assignments.id AND company_location_assignments.company_id = " + c.Param("companyId") + " INNER JOIN locations ON company_location_assignments.location_id = locations.id ").Find(&list)
	c.JSON(http.StatusOK, list)
}

func (a *DriverController) AddNewDriver(c *gin.Context) {
	var data models.Driver
	var response = addNewDriverResponse{Status: false}
	data.Name = c.PostForm("name")
	data.MobileNumber = c.PostForm("mobile_number")
	data.LicenseNumber = c.PostForm("license_number")
	locationID, convertError := strconv.Atoi(c.PostForm("location_id"))
	var dialCode = 0
	dialCode, convertError = strconv.Atoi(c.PostForm("dial_code"))
	if convertError == nil {
		data.DialCode = int64(dialCode)
		data.CompanyLocationAssignmentID = uint(locationID)
		form, _ := c.MultipartForm()
		fmt.Println("file count = %d", len(form.File))
		driverImage, err := c.FormFile("image")
		if err != nil {
			fmt.Println(err)
			response.Message = "Active Image is required"
			c.JSON(http.StatusBadRequest, response)
			fmt.Println(response)
			return
		}
		if (data.DialCode) == 0 {
			response.Message = "Dial code is required"
			c.JSON(http.StatusOK, response)
			return
		} else if len(data.MobileNumber) < 6 {
			response.Message = "Mobile number is required"
			c.JSON(http.StatusOK, response)
			return
		} else if len(data.Name) < 3 {
			response.Message = "Driver name is required"
			c.JSON(http.StatusOK, response)
			return
		} else if len(data.LicenseNumber) < 5 {
			response.Message = "License Number is required"
			c.JSON(http.StatusOK, response)
			return
		} else if data.CompanyLocationAssignmentID == 0 {
			response.Message = "Location Id is required"
			c.JSON(http.StatusOK, response)
			return
		} else {
			imageFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + driverImage.Filename
			if err := c.SaveUploadedFile(driverImage, "public/driver/"+imageFileName); err != nil {
				response.Message = fmt.Sprintf("upload file err: %s", err.Error())
				c.JSON(http.StatusBadRequest, response)
				return
			} else {
				data.Image = "public/driver/" + imageFileName
				data.IsActive = true
				var count = 0
				database.Db.Model(&models.Driver{}).Where("dial_code = ? AND mobile_number = ?", data.DialCode, data.MobileNumber).Count(&count)
				if count == 0 {
					result := database.Db.Create(&data)
					if result.Error == nil {
						response.Status = true
						response.Message = "Driver added successfully!"
						c.JSON(http.StatusOK, response)
						return
					} else {
						response.Message = result.Error.Error()
						c.JSON(http.StatusOK, response)
						return
					}
				} else {
					response.Message = "Driver mobile number already exists!"
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

func (a *DriverController) SendOtp(c *gin.Context) {
	var data sendOtpRequest
	var response = sendOtpResponse{Status: false}
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

		//check driver account exist
		var driverCount = 0
		database.Db.Model(&models.Driver{}).Where("dial_code = ? AND mobile_number = ? AND is_active = true", data.DialCode, data.MobileNumber).Count(&driverCount)
		if driverCount == 0 {
			response.Message = "Driver account is not available"
			c.JSON(http.StatusOK, response)
		} else {
			//Validation Success and so send the otp
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
}

type verifyOtpDriverResponse struct {
	Status        bool
	Message       string
	DriverDetails models.Driver
}

func (a *DriverController) VerifyOtp(c *gin.Context) {
	var data verifyOtpRequest
	var response = verifyOtpDriverResponse{Status: false}
	c.BindJSON(&data)
	if (data.DialCode) == 0 {
		response.Message = "DialCode is required"
		c.JSON(http.StatusOK, response)
		return
	} else if len(data.CountryCode) == 0 {
		response.Message = "CountryCode is required"
		c.JSON(http.StatusOK, response)
		return
	} else if len(data.MobileNumber) < 6 {
		response.Message = "MobileNumber is required"
		c.JSON(http.StatusOK, response)
		return
	} else if len(data.Otp) != 4 {
		response.Message = "Otp must contain 4 digits"
		c.JSON(http.StatusOK, response)
		return
	} else {
		var otpDetails models.Otp
		database.Db.Where("dial_code = ? AND country_code = ? AND mobile_number = ? AND is_used = ?", data.DialCode, data.CountryCode, data.MobileNumber, false).First(&otpDetails)
		if data.Otp == otpDetails.Otp {
			database.Db.Model(&otpDetails).UpdateColumn("is_used", true)
			var driverDetails models.Driver
			database.Db.Model(&models.Driver{}).Where("dial_code = ? AND mobile_number = ? AND is_active = true", data.DialCode, data.MobileNumber).First(&driverDetails)
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, config.JwtClaims{
				driverDetails.ID,
				"driver",
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
					Issuer:    "onride",
				},
			})
			tokenString, err := token.SignedString(config.JwtSecretKey)
			if err != nil {
				response.Message = err.Error()
				response.Status = false
				c.JSON(http.StatusOK, response)
				return
			} else {
				database.Db.Model(&driverDetails).UpdateColumn("auth_token", tokenString)
				response.DriverDetails = driverDetails
				response.Status = true
				response.Message = "Otp verified successfully"
				c.JSON(http.StatusOK, response)
				return
			}

		} else {
			response.Message = "Invalid Otp"
			c.JSON(http.StatusOK, response)
			return
		}

	}

}

type DriverStatusChangeRequest struct {
	VehicleId int64 `json:"vehicle_id"`
}
type DriverStatusResponse struct {
	Status  bool
	Message string
}

func (d *DriverController) GoOnline(c *gin.Context) {
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var data DriverStatusChangeRequest
	var response = DriverStatusResponse{Status: false}
	c.BindJSON(&data)
	if data.VehicleId == 0 {
		response.Message = "Invalid vehicle id"
		c.JSON(http.StatusOK, response)
		return
	} else {
		var isAlreadyOnline = 0
		database.Db.Model(&models.DriverVehicleAssignment{}).Where("driver_id = ? AND vehicle_id != ? AND is_online = true AND is_active = true", userData.UserID, data.VehicleId).Count(&isAlreadyOnline)
		if isAlreadyOnline == 1 {
			response.Message = "Sorry ! You cannot go online for more than 1 vehicle in same time."
			c.JSON(http.StatusOK, response)
			return
		} else {
			var vehicleAssignment models.DriverVehicleAssignment
			database.Db.Where("driver_id = ? AND vehicle_id = ? AND is_active = true", userData.UserID, data.VehicleId).First(&vehicleAssignment)
			if vehicleAssignment.ID == 0 {
				response.Message = "Sorry ! Access to the vehicle is currently removed.Please try another vehicle."
				c.JSON(http.StatusOK, response)
				return
			} else {
				var count = 0
				database.Db.Model(&models.DriverVehicleAssignment{}).Where("vehicle_id = ? AND driver_id != ? AND is_active = true AND is_online = true", data.VehicleId, userData.UserID).Count(&count)
				if count == 0 {
					database.Db.Model(&vehicleAssignment).UpdateColumn("is_online", true)
					response.Status = true
					response.Message = "Success! Now you are Online"
				} else {
					response.Message = "The vehicle is already in use.Please select another vehicle."
				}
				c.JSON(http.StatusOK, response)
				return
			}
		}

	}

}

func (d *DriverController) GoOffline(c *gin.Context) {
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var data DriverStatusChangeRequest
	var response = DriverStatusResponse{Status: false}
	c.BindJSON(&data)
	database.Db.Model(&models.DriverVehicleAssignment{}).Where("driver_id = ? AND vehicle_id = ? AND is_online = true", userData.UserID, data.VehicleId).UpdateColumn("is_online", false)
	response.Status = true
	response.Message = "Success! Now you are Offline"
	c.JSON(http.StatusOK, response)
	return

}

func (a *DriverController) EnableDriver(c *gin.Context) {
	var response struct {
		Status bool
	}
	var categoryId = c.Param("driverId")
	res := database.Db.Model(&models.Driver{}).Where("id = ?", categoryId).UpdateColumn("is_active", true)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}

func (a *DriverController) DisableDriver(c *gin.Context) {
	var response struct {
		Status bool
	}
	var categoryId = c.Param("driverId")
	res := database.Db.Model(&models.Driver{}).Where("id = ?", categoryId).UpdateColumn("is_active", false)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}
