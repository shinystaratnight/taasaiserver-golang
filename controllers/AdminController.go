package controllers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"taxi/models"
	"taxi/shared/config"
	"taxi/shared/database"
	"time"
)

type AdminController struct {
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type authRequest struct {
	Email    string
	Password string
}
type authResponse struct {
	Status  bool
	Message string
	IsAdmin bool
	Token   string
}

func (a *AdminController) Authenticate(c *gin.Context) {
	var data authRequest
	var response = authResponse{Status: false}

	c.BindJSON(&data)
	if !validateEmail(data.Email) {
		response.Message = "Email is not valid"
		c.JSON(http.StatusOK, response)
		return
	} else if len(data.Password) < 3 {
		response.Message = "Password length should be minimum 6"
		c.JSON(http.StatusOK, response)
		return
	} else {
		var admin models.Admin
		database.Db.Where("email = ? AND is_active = true", data.Email).First(&admin)
		if admin.ID != 0 {
			if checkPasswordHash(data.Password, admin.Password) {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, config.JwtClaims{
					(admin.ID),
					"admin",
					jwt.StandardClaims{
						ExpiresAt: time.Now().Add(time.Hour * 24 * 2).Unix(),
						Issuer:    "taasai",
					},
				})
				tokenString, err := token.SignedString(config.JwtSecretKey)
				if err != nil {
					response.Message = err.Error()
					response.Status = false
				} else {
					database.Db.Model(&admin).UpdateColumn("auth_token", tokenString)
					response.Status = true
					response.IsAdmin = true
					response.Token = tokenString
					response.Message = "User verified successfully"
				}
				c.JSON(http.StatusOK, response)
				return
			} else {
				response.Message = "Email or Password incorrect"
				c.JSON(http.StatusOK, response)
				return
			}
		}  else {
			var operator models.Operator
			database.Db.Where("email = ? AND is_active = true", data.Email).First(&operator)
			if operator.ID != 0{
				if checkPasswordHash(data.Password, operator.Password) {
					token := jwt.NewWithClaims(jwt.SigningMethodHS256, config.JwtClaims{
						(operator.ID),
						"operator",
						jwt.StandardClaims{
							ExpiresAt: time.Now().Add(time.Hour * 24 * 2).Unix(),
							Issuer:    "taasai",
						},
					})
					tokenString, err := token.SignedString(config.JwtSecretKey)
					if err != nil {
						response.Message = err.Error()
						response.Status = false
					} else {
						database.Db.Model(&operator).UpdateColumn("auth_token", tokenString)
						response.Status = true
						response.Token = tokenString
						response.Message = "User verified successfully"
					}
					c.JSON(http.StatusOK, response)
					return
				} else {
					response.Message = "Email or Password incorrect"
					c.JSON(http.StatusOK, response)
					return
				}
			}else{
				response.Message = "Email or Password incorrect"
				c.JSON(http.StatusOK, response)
				return
			}


		}
	}
}

type addNewAdminResponse struct {
	Status  bool
	Message string
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}
func (a *AdminController) AddNewAdmin(c *gin.Context) {
	var data models.Admin
	var response = addNewAdminResponse{Status: false}
	_ = c.BindJSON(&data)
	if len(data.Name) < 3 {
		response.Message = "Name must contain 3 characters"
		c.JSON(http.StatusOK, response)
		return
	} else if !validateEmail(data.Email) {
		response.Message = "Email is not valid"
		c.JSON(http.StatusOK, response)
		return
	} else if len(data.Password) < 3 {
		response.Message = "Password length should be minimum 6"
		c.JSON(http.StatusOK, response)
		return
	} else {
		hashedPassword, err := hashPassword(data.Password)
		if err == nil {
			data.Password = hashedPassword
			data.IsActive = true
			result := database.Db.Create(&data)
			if result.Error == nil {
				response.Status = true
				response.Message = "Admin user created successfully"
				c.JSON(http.StatusOK, response)
				return
			} else {
				response.Message = result.Error.Error()
				c.JSON(http.StatusOK, response)
				return
			}
		} else {
			response.Message = "Something went wrong"
			c.JSON(http.StatusOK, response)
			return
		}
	}
}
