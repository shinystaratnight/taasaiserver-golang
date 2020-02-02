package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"taxi/models"
	"taxi/shared/config"
	"taxi/shared/database"

	"github.com/gin-gonic/gin"
)

type OperatorController struct {
}
type polyPoint struct {
	Lat float64
	Lng float64
}

type AddOperatorRequest struct {
	Name     string
	Currency string
	Password string
	ConfirmPassword   string
	LocationName       string
	Email      string
	PlatformCommission float64
	OperatorCommission float64
	WorkTime int
	RestTime int
	Polygon  []polyPoint
	Docs []Doc
}

type Doc struct{
	Name string
}

type addZoneRequest struct {
	Name       string
	LocationID uint
	PickupPoints []models.PickupPoint
	Polygon    []polyPoint
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func (a *OperatorController) GetOperators(c *gin.Context) {
	type locationWithFareCount struct {
		ID             uint
		Name           string
		LocationName   string
		DriverWorkTime   float64
		DriverRestTime   float64
		Currency       string
		IsActive       bool
		TotalFareCount int64
	}
	var list []locationWithFareCount
	database.Db.Raw("SELECT (SELECT COUNT(*) as total_fare_count FROM fares  WHERE fares.operator_id = operators.id AND fares.is_active = true),* FROM operators").Scan(&list)
	c.JSON(http.StatusOK, list)
}

func (a *OperatorController) GetZones(c *gin.Context) {
	var list []models.Zone
	database.Db.Where("operator_id = ?", c.Param("locationId")).Find(&list)
	c.JSON(http.StatusOK, list)
}

func (a *OperatorController) GetCoordinates(c *gin.Context) {
	type locationWithCoordinate struct {
		Coordinates string
	}
	var location locationWithCoordinate
	database.Db.Raw("select btrim(st_astext(polygon), 'POLYGON()') as coordinates from operators Where operators.id = " + c.Param("locationId")).Scan(&location)
	c.JSON(http.StatusOK, location)

}

func (a *OperatorController) GetActiveOperators(c *gin.Context) {
	var list []models.Operator
	database.Db.Where("is_active = ?", true).Find(&list)
	c.JSON(http.StatusOK, list)
}

type DriverDocument struct {
	ID uint
	OperatorID     uint
	Name  string
	IsUploaded bool
	IsActive       bool
}
func (a *OperatorController) GetDriverDocs(c *gin.Context) {
	var list []DriverDocument
	database.Db.Where("operator_id = ?",c.Param("id")).Find(&list)
	var docsuploaded []models.DriverDocumentUpload
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	database.Db.Where("driver_id = ?",userData.UserID).Find(&docsuploaded)
	for i, doc := range list {
		for _,uploadedDoc := range docsuploaded{
			if uploadedDoc.DocID == doc.ID{
				list[i].IsUploaded = true
			}
		}
	}
	c.JSON(http.StatusOK, list)
}

func (a *OperatorController) GetActiveLocationsForCompany(c *gin.Context) {
	type result struct {
		Name       string
		ID         uint
		LocationID uint
	}
	var list []result
	database.Db.Raw("SELECT locations.id as location_id,locations.name,company_location_assignments.id FROM company_location_assignments INNER JOIN locations ON company_location_assignments.location_id = locations.id AND locations.is_active = true WHERE company_location_assignments.company_id = " + c.Param("companyId")).Scan(&list)
	c.JSON(http.StatusOK, list)
}

func (a *OperatorController) GetOperatorById(c *gin.Context) {
	var locationId = c.Param("locationId")
	var data models.Operator
	database.Db.Where("id = ?", locationId).First(&data)
	c.JSON(http.StatusOK, data)
}



func (a *OperatorController) EnableLocation(c *gin.Context) {
	var response struct {
		Status bool
	}
	var locationId = c.Param("locationId")
	res := database.Db.Model(&models.Operator{}).Where("id = ?", locationId).UpdateColumn("is_active", true)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}
func (a *OperatorController) DisableLocation(c *gin.Context) {
	var response struct {
		Status bool
	}
	var locationId = c.Param("locationId")
	res := database.Db.Model(&models.Operator{}).Where("id = ?", locationId).UpdateColumn("is_active", false)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}

func (a *OperatorController) EnableZone(c *gin.Context) {
	var response struct {
		Status bool
	}
	var locationId = c.Param("locationId")
	res := database.Db.Model(&models.Zone{}).Where("id = ?", locationId).UpdateColumn("is_active", true)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}
func (a *OperatorController) DisableZone(c *gin.Context) {
	var response struct {
		Status bool
	}
	var locationId = c.Param("locationId")
	res := database.Db.Model(&models.Zone{}).Where("id = ?", locationId).UpdateColumn("is_active", false)
	if res.Error == nil {
		response.Status = true
	} else {
		response.Status = false
	}
	c.JSON(http.StatusOK, response)
}

func (a *OperatorController) AddNewOperator(c *gin.Context) {
	var data AddOperatorRequest
	var response = sendOtpResponse{Status: false}
	c.BindJSON(&data)
	if len(data.Name) < 3 {
		response.Message = "Name is required"
		c.JSON(http.StatusOK, response)
	} else if data.Password != data.ConfirmPassword {
		response.Message = "Passwords doesn't match"
		c.JSON(http.StatusOK, response)
		return
	} else if !validateEmail(data.Email) {
		response.Message = "Email is not valid"
		c.JSON(http.StatusOK, response)
		return
	} else if data.PlatformCommission < 0 {
		response.Message = "Commission must be greater than or equal to 0"
		c.JSON(http.StatusOK, response)
		return
	}else if data.OperatorCommission < 0 {
		response.Message = "Commission must be greater than or equal to 0"
		c.JSON(http.StatusOK, response)
		return
	}else if data.WorkTime <= 0 {
		response.Message = "WorkTime must be greater than 0"
		c.JSON(http.StatusOK, response)
		return
	}else if data.RestTime <= 0 {
		response.Message = "RestTime must be greater than 0"
		c.JSON(http.StatusOK, response)
		return
	} else {
		hashedPassword, err := hashPassword(data.Password)
		if err == nil {
			var polyString = ""
			for i := 0; i < len(data.Polygon); i++ {
				if i != 0 {
					polyString += ","
				}
				polyString += FloatToString(data.Polygon[i].Lat) + " " + FloatToString(data.Polygon[i].Lng)
			}
			var intersectLocation models.Operator
			var res = database.Db.Where("ST_Intersects(polygon,ST_GeometryFromText('POLYGON((" + polyString + "))'))").First(&intersectLocation)
			log.Println("count = ", res.RowsAffected)
			if intersectLocation.ID == 0 {
				var dataString = fmt.Sprintf(" '%s' , '%s' , '%s', %f , %f , %d , %d ",data.LocationName,data.Email,hashedPassword,data.PlatformCommission,data.OperatorCommission,data.WorkTime,data.RestTime)
				var newLocationAddResponse = database.Db.Exec("INSERT INTO operators (name,currency, polygon,is_active,location_name,email,password,platform_commission,operator_commission,driver_work_time,driver_rest_time) VALUES ('" + data.Name + "','" + data.Currency + "',ST_GeometryFromText('POLYGON((" + polyString + "))'),true,"+dataString+");")
				if newLocationAddResponse.Error != nil {
					response.Message = newLocationAddResponse.Error.Error()
				} else {
					response.Message = "Operator added successfully"
					response.Status = true

					var operator models.Operator
					database.Db.Where("email = ?",data.Email).First(&operator)
					if operator.ID!=0{
						for  i:=0;i<len(data.Docs);i++{
							doc:=models.DriverDocument{OperatorID:operator.ID,Name:data.Docs[i].Name}
							database.Db.Create(&doc);
						}
					}
				}

			} else {
				response.Message = "Location intersects the previously created location named " + intersectLocation.Name

			}
		}

		c.JSON(http.StatusOK, response)
	}
}



func (a *OperatorController) AddNewZone(c *gin.Context) {
	var data addZoneRequest
	var response = sendOtpResponse{Status: false}
	c.BindJSON(&data)
	if len(data.Name) < 3 {
		response.Message = "Name is required"
		c.JSON(http.StatusOK, response)
	} else {
		var polyString = ""
		for i := 0; i < len(data.Polygon); i++ {
			if i != 0 {
				polyString += ","
			}
			polyString += FloatToString(data.Polygon[i].Lat) + " " + FloatToString(data.Polygon[i].Lng)
		}
		var count = 0
		//check if zone is within the city
		database.Db.Model(&models.Operator{}).Where("id = ? AND ST_Contains(polygon,ST_GeometryFromText('POLYGON(("+polyString+"))'))", data.LocationID).Count(&count)
		if count == 0 {
			response.Message = "Some part of the zone is outside the city."

		} else {
			//check if zone is intersecting other zone in the city
			var intersectLocation models.Zone
			var res = database.Db.Where("ST_Intersects(polygon,ST_GeometryFromText('POLYGON((" + polyString + "))'))").First(&intersectLocation)
			log.Println("count = ", res.RowsAffected)
			if intersectLocation.ID == 0 {
				var newZone models.Zone
				var newLocationAddResponse = database.Db.Raw("INSERT INTO zones (name,operator_id, polygon,is_active) VALUES ('" + data.Name + "','" + strconv.Itoa(int(data.LocationID)) + "',ST_GeometryFromText('POLYGON((" + polyString + "))'),true) RETURNING id;").Scan(&newZone)
				if newLocationAddResponse.Error != nil {
					response.Message = newLocationAddResponse.Error.Error()
				} else {
					response.Message = "Zone added successfully"
					response.Status = true

					for _, item := range data.PickupPoints {
						item.IsActive = true
						item.ZoneID = newZone.ID
						database.Db.Create(&item)
					}
				}

			} else {
				response.Message = "Zone intersects the previously created zone named " + intersectLocation.Name

			}
		}

		c.JSON(http.StatusOK, response)
	}
}
