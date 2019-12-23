package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taxi/models"
	"taxi/shared/config"
	"taxi/shared/database"
)

type EmergencyContactController struct{

}

func (e *EmergencyContactController) AddNewPassengerContact(c *gin.Context){
	type responseFormat struct{
		Status bool
	}
	var response = responseFormat{Status:true}
	var newContact models.EmergencyContact
	c.BindJSON(&newContact)
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	newContact.UserID = userData.UserID
	newContact.IsPassenger = true
	newContact.IsActive = true
	database.Db.Create(&newContact)
	c.JSON(http.StatusOK,response)
}


func (e *EmergencyContactController) GetPassengerContacts(c *gin.Context){
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var list []models.EmergencyContact
	database.Db.Where("is_passenger = true AND user_id = ?",userData.UserID).Find(&list)
	c.JSON(http.StatusOK,list)
}

func (e *EmergencyContactController) AddNewDriverContact(c *gin.Context){
	type responseFormat struct{
		Status bool
	}
	var response = responseFormat{Status:true}
	var newContact models.EmergencyContact
	c.BindJSON(&newContact)
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	newContact.UserID = userData.UserID
	newContact.IsPassenger = false
	newContact.IsActive = true
	database.Db.Create(&newContact)
	c.JSON(http.StatusOK,response)
}


func (e *EmergencyContactController) GetDriverContacts(c *gin.Context){
	var userData = c.MustGet("jwt_data").(*config.JwtClaims)
	var list []models.EmergencyContact
	database.Db.Where("is_passenger = false AND user_id = ?",userData.UserID).Find(&list)
	c.JSON(http.StatusOK,list)
}