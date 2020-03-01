package controllers

import (
	"net/http"
	"strconv"
	"taxi/models"
	"taxi/shared/config"
	"taxi/shared/database"

	"github.com/gin-gonic/gin"
)

type FleetController struct {
}

type AddFleetRequest struct {
	ID              uint   `json:"id"`
	OperatorID      string `json:"operator_id"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email           string `json:"email"`
	IsActive        bool   `json:"is_active"`
}

type getFleetResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	OperatorID   int64  `json:"operator_id"`
	OperatorName string `json:"operator_name"`
	Email        string `json:"email"`
	IsActive     bool   `json:"is_active"`
}

type addFleetResponse struct {
	Status  bool
	Message string
	id      uint
}

func (a *FleetController) GetFleets(c *gin.Context) {
	var list []getFleetResponse
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	if userData.UserType == "admin" {
		database.Db.Raw("SELECT F.*, O.name operator_name FROM fleets F INNER JOIN operators O ON O.id = F.operator_id;").Find(&list)

	} else {
		database.Db.Raw("SELECT F.*, O.name operator_name FROM fleets F INNER JOIN operators O ON O.id = F.operator_id AND O.id = " + strconv.Itoa(int(userData.UserID)) + ";").Find(&list)

	}
	c.JSON(http.StatusOK, list)
}

func (a *FleetController) GetFleetById(c *gin.Context) {
	var id = c.Param("id")
	var data models.Fleet
	database.Db.Where("id = ?", id).First(&data)
	c.JSON(http.StatusOK, data)
}

func (a *FleetController) AddNewFleet(c *gin.Context) {
	var data AddFleetRequest
	var response = addFleetResponse{Status: false}
	c.BindJSON(&data)
	if len(data.Name) < 3 {
		response.Message = "Name is required"
		c.JSON(http.StatusOK, response)
	} else if len(data.Password) < 6 {
		response.Message = "Password min length is 6"
		c.JSON(http.StatusOK, response)
		return
	} else if data.Password != data.ConfirmPassword {
		response.Message = "Passwords doesn't match"
		c.JSON(http.StatusOK, response)
		return
	} else if data.Password != data.ConfirmPassword {
		response.Message = "Passwords doesn't match"
		c.JSON(http.StatusOK, response)
		return
	} else if !validateEmail(data.Email) {
		response.Message = "Email is not valid"
		c.JSON(http.StatusOK, response)
		return
	} else {
		var count = 0
		database.Db.Model(&models.Fleet{}).Where("email = ? AND id != ?", data.Email, data.ID).Count(&count)
		if count != 0 {
			response.Message = "Email is duplicated. Try with another one."
			response.Status = false
			c.JSON(http.StatusOK, response)
			return
		}

		hashedPassword, err := hashPassword(data.Password)
		data.Password = hashedPassword
		if err == nil {
			var newRow models.Fleet = models.Fleet{}
			newRow.ID = data.ID
			newRow.OperatorID = data.OperatorID
			newRow.Email = data.Email
			newRow.Password = hashedPassword
			newRow.Name = data.Name

			database.Db.Model(&models.Fleet{}).Where("id = ?", newRow.ID).Count(&count)

			if count == 0 {
				var result = database.Db.Create(&newRow)
				// var insertedRow models.Fleet
				if result.Error != nil {
					response.Message = result.Error.Error()
					response.Status = false
				} else {
					// result.Last(&insertedRow)
					// fmt.Println(insertedRow)
					response.Message = "Fleet Manager added successfully"
					response.Status = true
					// response.id = insertedRow.ID
				}
				c.JSON(http.StatusOK, response)
				return
			} else {
				var result = database.Db.Model(&models.Fleet{}).Where("id = ?", data.ID).Update(&newRow)
				if result.Error != nil {
					response.Message = result.Error.Error()
					response.Status = false
				} else {
					response.Message = "Fleet Manager edited successfully"
					response.Status = true
				}
				c.JSON(http.StatusOK, response)
				return
			}
		}
	}
}
