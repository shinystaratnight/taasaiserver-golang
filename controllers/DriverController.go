package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"taxi/models"
	"taxi/shared/config"
	"taxi/shared/database"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type DriverController struct {
}

type addNewDriverResponse struct {
	Status  bool
	Message string
	ID      uint
}
type getDriverResponse struct {
	ID                 int64  `json:"driver_id"`
	Name               string `json:"driver_name"`
	OperatorID         int64  `json:"operator_id"`
	DialCode           string `json:"dial_code"`
	MobileNumber       string `json:"mobile_number"`
	LicenseNumber      string `json:"license_number"`
	DriverImage        string `json:"driver_image"`
	FcmID              string `json:"fcm_id"`
	IsActive           bool   `json:"is_active"`
	IsProfileCompleted bool   `json:"is_profile_completed"`
	LocationName       string `json:"location_name"`
	OperatorName       string `json:"operator_name"`
	VehicleName        string `json:"vehicle_name"`
	VehicleNumber      string `json:"vehicle_number"`
	VehicleImage       string `json:"vehicle_image"`
	FleetManager       string `json:"fleet_manager"`
}

// type getDriverHistoryResponse struct {
// 	ID                 int64  `json:"ride_id"`
// 	Name               string `json:"driver_name"`
// 	OperatorID         int64  `json:"operator_id"`
// 	DialCode           string `json:"dial_code"`
// 	MobileNumber       string `json:"mobile_number"`
// 	LicenseNumber      string `json:"license_number"`
// 	DriverImage        string `json:"driver_image"`
// 	FcmID              string `json:"fcm_id"`
// 	IsActive           bool   `json:"is_active"`
// 	IsProfileCompleted bool   `json:"is_profile_completed"`
// 	LocationName       string `json:"location_name"`
// 	OperatorName       string `json:"operator_name"`
// 	VehicleName        string `json:"vehicle_name"`
// 	VehicleNumber      string `json:"vehicle_number"`
// 	VehicleImage       string `json:"vehicle_image"`
// }

type addVehicleAssignmentResponse struct {
	Status  bool
	Message string
}

func (a *DriverController) GetDrivers(c *gin.Context) {
	var list []getDriverResponse
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	if userData.UserType == "admin" {
		database.Db.Raw("SELECT D.* ,O.name AS operator_name, O.location_name, F.id fleet_manager FROM drivers D INNER JOIN operators O ON D.operator_id = O.id LEFT JOIN fleet_managers F ON F.id = D.fleet_id ").Find(&list)

	} else {
		database.Db.Raw("SELECT D.* ,O.name AS operator_name, O.location_name, F.id fleet_manager FROM drivers D INNER JOIN operators O ON D.operator_id = O.id AND operators.id = " + strconv.Itoa(int(userData.UserID)) + " LEFT JOIN fleet_managers F ON F.id = D.fleet_id").Find(&list)

	}
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
	var convertError error
	id, convertError := strconv.Atoi(c.PostForm("id"))
	if convertError != nil {
		data.ID = 0
	} else {
		data.ID = uint(id)
	}
	data.Name = c.PostForm("name")
	data.MobileNumber = c.PostForm("mobile_number")

	operatorID, convertError := strconv.Atoi(c.PostForm("operator_id"))
	if convertError != nil {
		response.Message = "Operator Id is required"
		c.JSON(http.StatusOK, response)
		return
	}
	data.OperatorID = operatorID

	data.LicenseNumber = c.PostForm("license_number")
	data.VehicleName = c.PostForm("vehicle_name")

	var vehicleTypeId uint64 = 0
	vehicleTypeId, convertError = strconv.ParseUint(c.PostForm("vehicle_type_id"), 10, 64)
	if convertError != nil {
		response.Message = "Vehicle Type Id is required"
		c.JSON(http.StatusOK, response)
		return
	}
	data.VehicleTypeID = uint(vehicleTypeId)

	data.VehicleBrand = c.PostForm("vehicle_brand")
	data.VehicleModel = c.PostForm("vehicle_model")
	data.VehicleColor = c.PostForm("vehicle_color")
	data.VehicleNumber = c.PostForm("vehicle_number")
	data.IsActive = false

	var dialCode = 0
	dialCode, convertError = strconv.Atoi(c.PostForm("dial_code"))
	if convertError == nil {
		data.DialCode = int64(dialCode)
		data.OperatorID = operatorID
		form, _ := c.MultipartForm()
		fmt.Println("file count = %d", len(form.File))
		var vehicleImageErr error
		var driverImageErr error
		vehicleImage, vehicleImageErr := c.FormFile("vehicle_image")
		driverImage, driverImageErr := c.FormFile("driver_image")
		// if err != nil {
		// 	fmt.Println(err)
		// 	response.Message = "Active Image is required"
		// 	c.JSON(http.StatusBadRequest, response)
		// 	fmt.Println(response)
		// 	return
		// }
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
		} else {
			// if vehicleImageErr != nil {
			// 	fmt.Println(vehicleImageErr)
			// 	response.Message = "Vehicle Image is required"
			// 	c.JSON(http.StatusBadRequest, response)
			// 	fmt.Println(response)
			// 	return
			// } else {
			if vehicleImageErr == nil {
				imageFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + vehicleImage.Filename
				if err := c.SaveUploadedFile(vehicleImage, "public/vehicle/"+imageFileName); err != nil {
					response.Message = fmt.Sprintf("upload file err: %s", err.Error())
					c.JSON(http.StatusBadRequest, response)
					return
				}
				data.VehicleImage = "public/driver/" + imageFileName
			}
			// if driverImageErr != nil {
			// 	fmt.Println(driverImageErr)
			// 	response.Message = "Driver Image is required"
			// 	c.JSON(http.StatusBadRequest, response)
			// 	fmt.Println(response)
			// 	return
			// } else {
			if driverImageErr == nil {
				imageFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + driverImage.Filename
				if err := c.SaveUploadedFile(driverImage, "public/driver/"+imageFileName); err != nil {
					response.Message = fmt.Sprintf("upload file err: %s", err.Error())
					c.JSON(http.StatusBadRequest, response)
					return
				}
				data.DriverImage = "public/driver/" + imageFileName
			}

			data.IsActive = false
			var count = 0
			rows := database.Db.Model(&models.Driver{}).Where("id = ?", data.ID)
			rows.Count(&count)
			if count == 0 {
				if vehicleImageErr != nil {
					fmt.Println(vehicleImageErr)
					response.Message = "Vehicle Image is required"
					c.JSON(http.StatusBadRequest, response)
					fmt.Println(response)
					return
				} else if driverImageErr != nil {
					fmt.Println(driverImageErr)
					response.Message = "Driver Image is required"
					c.JSON(http.StatusBadRequest, response)
					fmt.Println(response)
					return
				}
				database.Db.Model(&models.Driver{}).Where("dial_code = ? AND mobile_number = ?", data.DialCode, data.MobileNumber).Count(&count)
				if count == 0 {
					result := database.Db.Create(&data)
					var dbResult models.Driver
					result.Last(&dbResult)
					if result.Error == nil {
						response.Status = true
						response.Message = "Driver added successfully!"
						response.ID = dbResult.ID
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
			} else {
				var row models.Driver
				rows.First(&row)
				if vehicleImageErr != nil {
					data.VehicleImage = row.VehicleImage
				}
				if driverImageErr != nil {
					data.DriverImage = row.DriverImage
				}
				result := database.Db.Model(&models.Driver{}).Where("id = ?", data.ID).Update(&data)
				if result.Error == nil {
					response.Status = true
					response.Message = "Driver updated successfully!"
					c.JSON(http.StatusOK, response)
					return
				} else {
					response.Message = result.Error.Error()
					c.JSON(http.StatusOK, response)
					return
				}
				// response.Message = "Driver mobile number already exists!"
				// c.JSON(http.StatusOK, response)
				// return
			}
		}
	} else {
		response.Message = "Operator Id is required"
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
		database.Db.Model(&models.Driver{}).Where("dial_code = ? AND mobile_number = ?", data.DialCode, data.MobileNumber).Count(&driverCount)
		if driverCount == 0 {
			response.IsNew = true
		}
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

type verifyOtpDriverResponse struct {
	Status        bool
	Message       string
	DriverDetails models.Driver
}

type ApproveDriverRequest struct {
	DriverID int
}

func (a *DriverController) ApproveDriver(c *gin.Context) {
	var response = GenericResponse{Status: true}
	var data ApproveDriverRequest
	c.BindJSON(&data)
	database.Db.Model(&models.Driver{}).Where("id = ?", data.DriverID).UpdateColumn("is_active", true)
	c.JSON(http.StatusOK, response)

}
func (a *DriverController) SubmitForApproval(c *gin.Context) {
	var response = verifyOtpDriverResponse{Status: false}

	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var driver models.Driver
	database.Db.Where("id = ?", userData.UserID).First(&driver)
	if driver.ID == userData.UserID {
		var documentsRequired []models.DriverDocument
		database.Db.Where("operator_id = ? ", driver.OperatorID).Find(&documentsRequired)
		isAllDocumentSubmitted := true
		for i := 0; i < len(documentsRequired); i++ {
			var count = 0
			database.Db.Model(&models.DriverDocumentUpload{}).Where("doc_id = ? AND driver_id = ? AND is_active = true", documentsRequired[i].ID, userData.UserID).Count(&count)
			if count == 0 {
				isAllDocumentSubmitted = false
				break
			}

		}
		if isAllDocumentSubmitted {
			database.Db.Model(&driver).UpdateColumn("is_profile_completed", true)
			response.Status = true
		} else {
			response.Message = "Kindly Upload All Documents Before Submitting For Approval"
		}
	}
	c.JSON(http.StatusOK, response)

}

func (a *DriverController) UploadDriverDocument(c *gin.Context) {
	var response = GenericResponse{Status: false}

	docID, docIdError := strconv.Atoi(c.PostForm("id"))
	if docIdError == nil {
		form, _ := c.MultipartForm()
		fmt.Println("file count = %d", len(form.File))
		// Source
		documentImage, err := c.FormFile("document_image")
		if err != nil {
			fmt.Println(err)
			response.Message = "Doc Image is required"
			c.JSON(http.StatusOK, response)
			fmt.Println(response)
			return
		}

		driverImageFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + documentImage.Filename

		if err := c.SaveUploadedFile(documentImage, "public/driver/"+driverImageFileName); err != nil {
			response.Message = fmt.Sprintf("upload file err: %s", err.Error())
			c.JSON(http.StatusOK, response)
			return
		} else {
			var userData = c.MustGet("jwt_data").(*config.JwtClaims)

			database.Db.Model(&models.DriverDocumentUpload{}).Where("doc_id = ? AND driver_id = ?", docID, userData.UserID).UpdateColumn("is_active", false)

			var newDocUpload = models.DriverDocumentUpload{
				DocID:    uint(docID),
				DriverID: userData.UserID,
				Image:    driverImageFileName,
				IsActive: true,
			}

			database.Db.Create(&newDocUpload)
			response.Status = true
		}
	} else {
		response.Message = docIdError.Error()
	}

	c.JSON(http.StatusOK, response)

}

func (a *DriverController) CreateDriverAccount(c *gin.Context) {
	var response = verifyOtpDriverResponse{Status: false}

	name := c.PostForm("Name")
	otp := c.PostForm("Otp")
	mobile := c.PostForm("MobileNumber")
	countryCode := c.PostForm("CountryCode")
	vehicleName := c.PostForm("VehicleName")
	vehicleBrand := c.PostForm("VehicleBrand")
	vehicleModel := c.PostForm("VehicleModel")
	vehicleColor := c.PostForm("VehicleColor")
	vehicleNumber := c.PostForm("VehicleNumber")
	licenseNumber := c.PostForm("LicenseNumber")

	fmt.Println("Name = " + c.PostForm("Name"))
	fmt.Println("MobileNumber " + c.PostForm("MobileNumber"))
	fmt.Println("VehicleNumber " + c.PostForm("VehicleNumber"))
	fmt.Println("LicenseNumber " + c.PostForm("LicenseNumber"))
	fmt.Println("VehicleTypeID = " + c.PostForm("VehicleTypeID"))
	fmt.Println("DialCode = " + c.PostForm("DialCode"))
	fmt.Println("OperatorID = " + c.PostForm("OperatorID"))

	vehicleTypeID, vehicleTypeIdError := strconv.Atoi(c.PostForm("VehicleTypeID"))
	dialCode, dialCodeError := strconv.Atoi(c.PostForm("DialCode"))
	operatorID, operatorIDError := strconv.Atoi(c.PostForm("OperatorID"))
	fmt.Println("VehicleTypeID = " + c.PostForm("VehicleTypeID"))
	fmt.Println("DialCode = " + c.PostForm("DialCode"))
	fmt.Println("OperatorID = " + c.PostForm("OperatorID"))

	if vehicleTypeIdError == nil && dialCodeError == nil && operatorIDError == nil {

		form, _ := c.MultipartForm()
		fmt.Println("file count = %d", len(form.File))
		fmt.Printf("file count = %+v", (form.File["files[]"]))
		// Source
		driverImage, err := c.FormFile("driver_image")
		if err != nil {
			fmt.Println(err)
			response.Message = "Driver Image is required"
		} else {
			vehicleImage, err1 := c.FormFile("vehicle_image")
			if err1 != nil {
				response.Message = "Vehicle Image is required"
			} else {
				driverImageFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + driverImage.Filename
				vehicleImageFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + "_" + vehicleImage.Filename

				if err := c.SaveUploadedFile(driverImage, "public/driver/"+driverImageFileName); err != nil {
					response.Message = fmt.Sprintf("upload file err: %s", err.Error())
				} else {
					if err := c.SaveUploadedFile(vehicleImage, "public/vehicle/"+vehicleImageFileName); err != nil {
						response.Message = fmt.Sprintf("upload file err: %s", err.Error())
					} else {
						var otpDetails models.Otp
						database.Db.Where("dial_code = ? AND country_code = ? AND mobile_number = ? AND is_used = ?", dialCode, countryCode, mobile, false).First(&otpDetails)
						if otp == otpDetails.Otp {
							database.Db.Model(&otpDetails).UpdateColumn("is_used", true)
							var driver = models.Driver{
								Name:               name,
								DialCode:           int64(dialCode),
								MobileNumber:       mobile,
								OperatorID:         operatorID,
								VehicleName:        vehicleName,
								VehicleTypeID:      uint(vehicleTypeID),
								VehicleBrand:       vehicleBrand,
								VehicleModel:       vehicleModel,
								VehicleColor:       vehicleColor,
								VehicleNumber:      vehicleNumber,
								LicenseNumber:      licenseNumber,
								VehicleImage:       "public/vehicle/" + vehicleImageFileName,
								DriverImage:        "public/driver/" + driverImageFileName,
								IsProfileCompleted: false,
								IsActive:           false,
							}

							database.Db.Create(&driver)
							token := jwt.NewWithClaims(jwt.SigningMethodHS256, config.JwtClaims{
								driver.ID,
								"driver",
								jwt.StandardClaims{
									ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
									Issuer:    "taasai",
								},
							})
							tokenString, err := token.SignedString(config.JwtSecretKey)
							if err != nil {
								response.Message = err.Error()
								response.Status = false
							} else {
								database.Db.Model(&driver).UpdateColumn("auth_token", tokenString)
								response.DriverDetails = driver
								response.Status = true
								response.Message = "Driver Account Created And Submitted For Approval"
							}

						} else {
							response.Message = "Invalid Otp"
						}

					}
				}
			}
		}
	} else {
		if vehicleTypeIdError != nil {
			response.Message = "vehicleTypeIdError : " + vehicleTypeIdError.Error()
		} else if dialCodeError != nil {
			response.Message = " dialCodeError : " + dialCodeError.Error()
		} else if operatorIDError != nil {
			response.Message = "operatorIDError : " + operatorIDError.Error()
		}
	}
	c.JSON(http.StatusOK, response)

}

func (a *DriverController) GetDriverDetails(c *gin.Context) {
	var response = verifyOtpDriverResponse{Status: true}
	var driverDetails models.Driver

	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	database.Db.Model(&models.Driver{}).Where("id = ? ", userData.UserID).First(&driverDetails)
	response.DriverDetails = driverDetails
	c.JSON(http.StatusOK, response)

}

type GetDriverDetailsWithDocResponse struct {
	DriverDetails   models.Driver
	OperatorDetails models.Operator
	DocsRequired    []models.DriverDocument
	UploadedDocs    []models.DriverDocumentUpload
}

func (a *DriverController) GetDriverDetailsWithDoc(c *gin.Context) {
	var response = GetDriverDetailsWithDocResponse{}
	var driverDetails models.Driver
	var operatorDetails models.Operator
	var docsRequired []models.DriverDocument
	var docsuploaded []models.DriverDocumentUpload
	database.Db.Model(&models.Driver{}).Where("id = ? ", c.Param("id")).First(&driverDetails)
	database.Db.Model(&models.Operator{}).Where("id = ? ", driverDetails.OperatorID).First(&driverDetails)
	database.Db.Where("operator_id = ?", driverDetails.OperatorID).Find(&docsRequired)
	database.Db.Where("driver_id = ? AND is_active = true", driverDetails.ID).Find(&docsuploaded)
	response.DriverDetails = driverDetails
	response.DocsRequired = docsRequired
	response.OperatorDetails = operatorDetails
	response.UploadedDocs = docsuploaded
	c.JSON(http.StatusOK, response)

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
			database.Db.Model(&models.Driver{}).Where("dial_code = ? AND mobile_number = ? ", data.DialCode, data.MobileNumber).First(&driverDetails)
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, config.JwtClaims{
				driverDetails.ID,
				"driver",
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
					Issuer:    "taasai",
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
	var response = DriverStatusResponse{Status: true, Message: "Success! Now you are Online"}
	database.Db.Model(&models.Driver{}).Where("id = ?  AND is_active = true", userData.UserID).UpdateColumn("is_online", true)
	database.Db.Model(&models.Driver{}).Where("id = ?  AND is_active = true", userData.UserID).UpdateColumn("is_ride", false)
	c.JSON(http.StatusOK, response)
	return
}

func (d *DriverController) GoOffline(c *gin.Context) {
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var response = DriverStatusResponse{Status: true, Message: "Success! Now you are Offline"}
	database.Db.Model(&models.Driver{}).Where("id = ?", userData.UserID).UpdateColumn("is_online", false)
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

/*
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
*/

func (a *DriverController) GetDocument(c *gin.Context) {
	fileName := c.Param("name")
	downloadName := fileName
	header := c.Writer.Header()
	header["Content-type"] = []string{"application/octet-stream"}
	header["Content-Disposition"] = []string{"attachment; filename= " + downloadName}

	file, err := os.Open("./public/driver/" + fileName)
	if err != nil {
		c.String(http.StatusOK, "%v", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusOK, "%v", err)
		return
	}
}
