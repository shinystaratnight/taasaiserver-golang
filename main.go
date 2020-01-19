package main

import (
	"fmt"
	"log"
	"net/http"
	"taxi/controllers"
	"taxi/models"
	"taxi/paymentGateway/razorPay"
	"taxi/shared/config"
	"taxi/shared/database"
	"taxi/shared/googleMap"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

var (
	passengerController   = controllers.PassengerController{}
	locationController    = controllers.LocationController{}
	vehicleTypeController = controllers.VehicleTypeController{}
	vehicleController     = controllers.VehicleController{}
	driverController      = controllers.DriverController{}
	fareController        = controllers.FareController{}
	rideBookingController = controllers.RideBookingController{}
	adminController       = controllers.AdminController{}
	mqttController        = controllers.MqttController{}
	rideController        = controllers.RideController{}
	companyController     = controllers.CompanyController{}
	zoneFareController    = controllers.ZoneFareController{}
	dashboardController   = controllers.DashboardController{}
	emergencyController   = controllers.EmergencyContactController{}
	razorPayController    = razorPay.RazorPay{}
)

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.AbortWithStatus(code)
}

func tokenAuthMiddleware(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		fmt.Println("token = " + tokenString)
		if tokenString == "" {
			respondWithError(401, "API token required", c)
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &config.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return config.JwtSecretKey, nil
		})
		if err == nil {
			if claims, ok := token.Claims.(*config.JwtClaims); ok && token.Valid {
				var count = 0
				fmt.Println("%+v", claims)
				fmt.Println("id", claims.UserID)
				if err == nil {
					if claims.UserType == userType {
						if claims.UserType == "passenger" {
							database.Db.Model(&models.Passenger{}).Where("id = ? AND auth_token = ? AND is_active = true", claims.UserID, tokenString).Count(&count)
						} else if claims.UserType == "driver" {
							database.Db.Model(&models.Driver{}).Where("id = ? AND auth_token = ? AND is_active = true", claims.UserID, tokenString).Count(&count)
						} else if claims.UserType == "admin" {
							database.Db.Model(&models.Admin{}).Where("id = ? AND auth_token = ? AND is_active = true", claims.UserID, tokenString).Count(&count)
						}
					}

				}
				if count == 0 {
					respondWithError(401, "Invalid API token", c)
					return
				} else {
					c.Set("jwt_data", claims)
				}
			} else {
				fmt.Println("token invalid")
				respondWithError(401, "Invalid API token", c)
				return
			}
		} else {
			fmt.Println(err.Error())
			respondWithError(401, "Invalid API token", c)
			return
		}

		c.Next()
	}
}

func setupRouter() http.Handler {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Static("/public", "./public")
	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Server running")
	})

	adminRoutePrivatePublic := router.Group("/admin")
	{
		adminRoutePrivatePublic.POST("/login", adminController.Authenticate)
		adminRoutePrivatePublic.POST("/new", adminController.AddNewAdmin)
	}

	adminRoutePrivate := router.Group("/admin")
	adminRoutePrivate.Use(tokenAuthMiddleware("admin"))
	{
		adminRoutePrivate.POST("/addNewVehicle", vehicleController.AddNewVehicle)
		adminRoutePrivate.GET("/getVehicles", vehicleController.GetVehicles)
		adminRoutePrivate.GET("/getVehiclesOfCompany/:companyId", vehicleController.GetVehiclesOfCompany)
		adminRoutePrivate.POST("/addNewDriver", driverController.AddNewDriver)
		adminRoutePrivate.GET("/getVehicleAssignments/:driverId", driverController.GetVehicleAssignmentsForID)
		adminRoutePrivate.POST("/addNewVehicleAssignment", driverController.AddNewVehicle)
		adminRoutePrivate.GET("/getDrivers", driverController.GetDrivers)
		adminRoutePrivate.POST("/addNewVehicleType", vehicleTypeController.AddNewVehicleType)
		adminRoutePrivate.POST("/editVehicleType", vehicleTypeController.EditVehicleType)
		adminRoutePrivate.POST("/addNewVehicleTypeCategory", vehicleTypeController.AddNewVehicleTypeCategory)
		adminRoutePrivate.GET("/getVehicleTypes", vehicleTypeController.GetVehicleTypes)
		adminRoutePrivate.GET("/getVehicleTypeWithID/:vehicleTypeId", vehicleTypeController.GetVehicleTypeWithID)
		adminRoutePrivate.GET("/getVehicleTypes/:categoryId", vehicleTypeController.GetVehicleTypesForCategory)
		adminRoutePrivate.GET("/getVehicleTypeCategories", vehicleTypeController.GetVehicleTypeCategories)
		adminRoutePrivate.GET("/getActiveVehicleTypeCategories", vehicleTypeController.GetActiveVehicleTypeCategories)
		adminRoutePrivate.GET("/getVehicleTypesWithFare/:locationId", vehicleTypeController.GetVehicleTypesWithFare)
		adminRoutePrivate.GET("/getActiveVehicleTypes", vehicleTypeController.GetActiveVehicleTypes)
		adminRoutePrivate.PUT("/enableVehicleType/:vehicleTypeId", vehicleTypeController.EnableVehicleType)
		adminRoutePrivate.PUT("/disableVehicleType/:vehicleTypeId", vehicleTypeController.DisableVehicleType)
		adminRoutePrivate.GET("/getUnAssignedVehicleTypeForZone/:locationId", vehicleTypeController.GetUnAssignedVehicleTypesForZone)

		adminRoutePrivate.GET("/getUnAssignedVehicleType/:locationId", vehicleTypeController.GetUnAssignedVehicleTypes)

		adminRoutePrivate.GET("/getFares", fareController.GetActiveFare)
		adminRoutePrivate.GET("/getZoneFares/:zoneId", zoneFareController.GetActiveZoneFare)
		adminRoutePrivate.GET("/getFares/:locationId", fareController.GetActiveFareForLocation)
		adminRoutePrivate.PUT("/disableFare", fareController.DisableFare)
		adminRoutePrivate.PUT("/disableZoneFare", zoneFareController.DisableZoneFare)
		adminRoutePrivate.POST("/addNewFare", fareController.AddNewFare)
		adminRoutePrivate.POST("/addNewZoneFare", zoneFareController.AddNewZoneFare)
		adminRoutePrivate.POST("/addNewLocation", locationController.AddNewLocation)
		adminRoutePrivate.POST("/addNewZone", locationController.AddNewZone)
		adminRoutePrivate.POST("/addNewCompany", companyController.AddNewCompany)
		adminRoutePrivate.GET("/getCompanies", companyController.GetCompanies)

		adminRoutePrivate.GET("/getDriversForCompany/:companyId", driverController.GetDriversForCompany)
		adminRoutePrivate.GET("/getVehiclesForCompany/:companyId", vehicleController.GetVehiclesForCompany)

		adminRoutePrivate.PUT("/enableCompany/:companyId", companyController.EnableCompany)
		adminRoutePrivate.PUT("/disableCompany/:companyId", companyController.DisableCompany)

		adminRoutePrivate.GET("/getZones/:locationId", locationController.GetZones)
		adminRoutePrivate.GET("/getLocations", locationController.GetLocations)
		adminRoutePrivate.GET("/getActiveLocations", locationController.GetActiveLocations)
		adminRoutePrivate.GET("/getActiveLocationsForCompany/:companyId", locationController.GetActiveLocationsForCompany)
		adminRoutePrivate.GET("/getLocation/:locationId", locationController.GetLocationById)
		adminRoutePrivate.GET("/getCoordinates/:locationId", locationController.GetCoordinates)
		adminRoutePrivate.PUT("/enableLocation/:locationId", locationController.EnableLocation)
		adminRoutePrivate.PUT("/disableLocation/:locationId", locationController.DisableLocation)

		adminRoutePrivate.PUT("/enableZone/:locationId", locationController.EnableZone)
		adminRoutePrivate.PUT("/disableZone/:locationId", locationController.DisableZone)
		adminRoutePrivate.POST("/getRides", rideController.GetRides)
		adminRoutePrivate.GET("/getRideDetail/:rideId", rideController.GetRideDetail)
		adminRoutePrivate.GET("/getRideLocations/:rideId", rideController.GetRideLocations)
		adminRoutePrivate.PUT("/enableVehicleCategory/:categoryId", vehicleTypeController.EnableVehicleTypeCategory)
		adminRoutePrivate.PUT("/disableVehicleCategory/:categoryId", vehicleTypeController.DisableVehicleTypeCategory)
		adminRoutePrivate.POST("/editVehicleTypeCategory", vehicleTypeController.EditVehicleTypeCategory)

		adminRoutePrivate.PUT("/enableVehicle/:vehicleId", vehicleController.EnableVehicle)
		adminRoutePrivate.PUT("/disableVehicle/:vehicleId", vehicleController.DisableVehicle)

		adminRoutePrivate.PUT("/enableDriver/:driverId", driverController.EnableDriver)
		adminRoutePrivate.PUT("/disableDriver/:driverId", driverController.DisableDriver)

		adminRoutePrivate.PUT("/enableDriverAssignment/:id", driverController.EnableAssignment)
		adminRoutePrivate.PUT("/disableDriverAssignment/:id", driverController.DisableAssignment)
		adminRoutePrivate.GET("/getDataCount", dashboardController.GetDataCount)
		adminRoutePrivate.GET("/getAllPassengers", passengerController.GetAllPassengers)
		adminRoutePrivate.GET("/getRidesForPassenger/:passengerId", rideController.GetRidesForPassenger)
	}
	return router

}

func setupMobileAppRouter() http.Handler {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Static("/public", "./public")
	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Server running")
	})
	router.POST("/razorPay/webhook", razorPayController.Webhook)
	customerRoutePublic := router.Group("/customer")
	{
		customerRoutePublic.POST("/sendOtp", passengerController.SendOtp)
		customerRoutePublic.POST("/verifyOtp", passengerController.VerifyOtp)
	}

	customerRoutePrivate := router.Group("/customer")
	customerRoutePrivate.Use(tokenAuthMiddleware("passenger"))
	{
		customerRoutePrivate.POST("/completeProfile", passengerController.AddCustomerBasicDetails)
		customerRoutePrivate.POST("/getVehicleTypes", rideBookingController.GetEstimatedFare)
		customerRoutePrivate.POST("/bookRide", rideBookingController.BookRide)
		customerRoutePrivate.POST("/getRides", rideBookingController.GetBookingHistory)
		customerRoutePrivate.POST("/cancelRide", rideBookingController.CancelRide)
		customerRoutePrivate.POST("/checkIsOnRide", passengerController.CheckIsOnRide)
		customerRoutePrivate.POST("/getNearestDrivers", passengerController.GetNearByDrivers)
		customerRoutePrivate.POST("/getRideDetails/:rideId", rideController.GetRideDetailsForMobile)
		customerRoutePrivate.POST("/rateDriver", rideController.RateDriver)

		customerRoutePrivate.POST("/addNewPassengerContact", emergencyController.AddNewPassengerContact)
		customerRoutePrivate.POST("/getPassengerEmergencyContacts", emergencyController.GetPassengerContacts)

	}

	driverRoutePublic := router.Group("/driver")
	{
		driverRoutePublic.POST("/sendOtp", driverController.SendOtp)
		driverRoutePublic.POST("/verifyOtp", driverController.VerifyOtp)
	}

	driverRoutePrivate := router.Group("/driver")
	driverRoutePrivate.Use(tokenAuthMiddleware("driver"))
	{
		driverRoutePrivate.POST("/getVehicleAssignments", driverController.GetVehicleAssignments)
		driverRoutePrivate.POST("/goOnline", driverController.GoOnline)
		driverRoutePrivate.POST("/goOffline", driverController.GoOffline)
		driverRoutePrivate.POST("/acceptRide", rideController.RideAccept)
		driverRoutePrivate.POST("/driverArrived/:rideId", rideController.DriverArrived)
		driverRoutePrivate.POST("/startTrip/:rideId", rideController.StartTrip)
		driverRoutePrivate.POST("/stopTrip/:rideId", rideController.StopTrip)
		driverRoutePrivate.POST("/updateRideLocations", rideController.UpdateRideLocations)
		driverRoutePrivate.POST("/getRides", rideBookingController.GetDriverBookingHistory)
		driverRoutePrivate.POST("/getRideDetails/:rideId", rideController.GetRideDetailsForMobile)
		driverRoutePrivate.POST("/ratePassenger", rideController.RatePassenger)
		driverRoutePrivate.POST("/addNewDriverContact", emergencyController.AddNewDriverContact)
		driverRoutePrivate.POST("/getDriverEmergencyContacts", emergencyController.GetDriverContacts)

	}

	return router

}

func setupPrivateRouter() http.Handler {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/auth", mqttController.HandleMqttAuth)
	router.POST("/acl", mqttController.HandleMqttAuthorization)
	router.POST("/webhook", mqttController.WebHook)
	router.POST("/clientGone", mqttController.ClientGone)
	return router

}

func main() {
	database.SetupDb()
	googleMap.SetupClient()
	gin.SetMode(gin.ReleaseMode)
	adminRouter := setupRouter()
	mobileAppRouter := setupMobileAppRouter()
	privateRouter := setupPrivateRouter()

	publicServer := &http.Server{
		Addr:         ":4001",
		Handler:      adminRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	mobileAppServer := &http.Server{
		Addr:         ":4002",
		Handler:      mobileAppRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	privateServer := &http.Server{
		Addr:         ":4000",
		Handler:      privateRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return publicServer.ListenAndServe()
	})
	g.Go(func() error {
		return mobileAppServer.ListenAndServe()
	})
	g.Go(func() error {
		return privateServer.ListenAndServe()
	})
	mqttController.Connect()
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
