package controllers

import (
	"net/http"
	"taxi/models"
	"taxi/shared/database"
	"taxi/utils"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
}
type addNewCompanyRequest struct {
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
	Commission      float64
	Locations       []uint
}
type addNewCompanyResponse struct {
	Status  bool
	Message string
}

func (a *CompanyController) AddNewCompany(c *gin.Context) {
	var data addNewCompanyRequest
	var response = addNewCompanyResponse{Status: false}
	c.BindJSON(&data)
	if data.Password != data.ConfirmPassword {
		response.Message = "Passwords doesn't match"
		c.JSON(http.StatusOK, response)
		return
	} else if !validateEmail(data.Email) {
		response.Message = "Email is not valid"
		c.JSON(http.StatusOK, response)
		return
	} else if len(data.Locations) == 0 {
		response.Message = "Location is not valid"
		c.JSON(http.StatusOK, response)
		return
	} else if data.Commission < 0 {
		response.Message = "Commission must be greater than or equal to 0"
		c.JSON(http.StatusOK, response)
		return
	} else {
		hashedPassword, err := hashPassword(data.Password)
		if err == nil {
			var newCompany = models.Company{}
			newCompany.Password = hashedPassword
			newCompany.IsActive = true
			newCompany.Email = data.Email
			newCompany.Commission = data.Commission
			newCompany.Name = utils.Capitalize(data.Name)
			result := database.Db.Create(&newCompany)
			if result.Error == nil {
				response.Status = true
				response.Message = "Company created successfully"
				for _, location := range data.Locations {
					var newCompanyLocationAssignment = models.CompanyLocationAssignment{
						CompanyID:  newCompany.ID,
						LocationID: location,
						IsActive:   true,
					}
					database.Db.Create(&newCompanyLocationAssignment)
				}
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

func (a *CompanyController) GetCompanies(c *gin.Context) {
	type companyResponse struct {
		Name          string
		Email         string
		ID            uint
		IsActive      bool
		Locations     string
		Commission    float64
		DriversCount  int64
		VehiclesCount int64
	}
	var list []companyResponse
	//database.Db.Raw("SELECT companies.name,companies.commission,companies.email,companies.id,companies.is_active,(SELECT COUNT(*) FROM drivers WHERE drivers.company_location_assignment_id = company_location_assignments.id) as drivers_count,(SELECT COUNT(*) FROM vehicles WHERE vehicles.company_location_assignment_id = company_location_assignments.id) as vehicles_count,array_agg(locations.name) As locations FROM companies INNER JOIN company_location_assignments ON  company_location_assignments.company_id = companies.id INNER JOIN locations ON company_location_assignments.location_id = locations.id GROUP BY companies.id, company_location_assignments.id ").Scan(&list)
	database.Db.Raw("SELECT companies.name,companies.commission,companies.email,companies.id,companies.is_active,array_agg(locations.name) As locations FROM companies INNER JOIN company_location_assignments ON  company_location_assignments.company_id = companies.id INNER JOIN locations ON company_location_assignments.location_id = locations.id GROUP BY companies.id").Scan(&list)
	c.JSON(http.StatusOK, list)
}

type statusChangeResponse struct {
	Status  bool
	Message string
}

func (a *CompanyController) EnableCompany(c *gin.Context) {
	var response = statusChangeResponse{Status: true}
	database.Db.Model(&models.Company{}).Where("id = ?", c.Param("companyId")).UpdateColumn("is_active", true)
	c.JSON(http.StatusOK, response)
}

func (a *CompanyController) DisableCompany(c *gin.Context) {
	var response = statusChangeResponse{Status: true}
	database.Db.Model(&models.Company{}).Where("id = ?", c.Param("companyId")).UpdateColumn("is_active", false)
	c.JSON(http.StatusOK, response)
}
