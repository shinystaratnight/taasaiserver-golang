package main

import (
	"fmt"
	"log"
	"net/http"
	"taxi/controllers"
	"taxi/models"
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
	operatorController   = controllers.OperatorController{}
	vehicleTypeController = controllers.VehicleTypeController{}
	driverController      = controllers.DriverController{}
	fareController        = controllers.FareController{}
	rideBookingController = controllers.RideBookingController{}
	adminController       = controllers.AdminController{}
	mqttController        = controllers.MqttController{}
	rideController        = controllers.RideController{}
	companyController     = controllers.CompanyController{}
	zoneFareController    = controllers.ZoneFareController{}
	dashboardController   = controllers.DashboardController{}
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
							database.Db.Model(&models.Driver{}).Where("id = ? AND auth_token = ? ", claims.UserID, tokenString).Count(&count)
						} else if claims.UserType == "admin" {
							database.Db.Model(&models.Admin{}).Where("id = ? AND auth_token = ? AND is_active = true", claims.UserID, tokenString).Count(&count)

						}
					}else if claims.UserType == "operator" {
						database.Db.Model(&models.Operator{}).Where("id = ? AND auth_token = ?", claims.UserID, tokenString).Count(&count)
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
		adminRoutePrivate.POST("/addNewDriver", driverController.AddNewDriver)
		adminRoutePrivate.POST("/approveDriver", driverController.ApproveDriver)
		adminRoutePrivate.GET("/getDrivers", driverController.GetDrivers)
		adminRoutePrivate.GET("/getDriverDetailsWithDoc/:id", driverController.GetDriverDetailsWithDoc)
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

		adminRoutePrivate.GET("/getUnAssignedVehicleType/:operatorId", vehicleTypeController.GetUnAssignedVehicleTypes)

		adminRoutePrivate.GET("/getFares", fareController.GetActiveFare)
		adminRoutePrivate.GET("/getFareByID/:id", fareController.GetFareByID)
		adminRoutePrivate.GET("/GetOperatorByID/:id", operatorController.GetOperatorByID)
		adminRoutePrivate.GET("/getZoneFares/:zoneId", zoneFareController.GetActiveZoneFare)
		adminRoutePrivate.GET("/getFares/:operatorId", fareController.GetActiveFareForLocation)
		adminRoutePrivate.PUT("/disableFare", fareController.DisableFare)
		adminRoutePrivate.PUT("/editFare", fareController.EditFare)
		adminRoutePrivate.PUT("/disableZoneFare", zoneFareController.DisableZoneFare)
		adminRoutePrivate.POST("/addNewFare", fareController.AddNewFare)
		adminRoutePrivate.POST("/addNewZoneFare", zoneFareController.AddNewZoneFare)

		adminRoutePrivate.POST("/addNewOperator", operatorController.AddNewOperator)
		adminRoutePrivate.POST("/addOperatorDoc", operatorController.AddOperatorDoc)
		adminRoutePrivate.POST("/getOperatorDocs", operatorController.OperatorDocs)

		adminRoutePrivate.POST("/addNewZone", operatorController.AddNewZone)

		adminRoutePrivate.GET("/getDriversForCompany/:companyId", driverController.GetDriversForCompany)


		adminRoutePrivate.GET("/getZones/:locationId", operatorController.GetZones)
		adminRoutePrivate.GET("/getLocations", operatorController.GetOperators)
		adminRoutePrivate.GET("/getActiveLocations", operatorController.GetActiveOperators)
		adminRoutePrivate.GET("/getActiveLocationsForCompany/:companyId", operatorController.GetActiveLocationsForCompany)
		adminRoutePrivate.GET("/getLocation/:locationId", operatorController.GetOperatorById)
		adminRoutePrivate.GET("/getCoordinates/:locationId", operatorController.GetCoordinates)
		adminRoutePrivate.PUT("/enableLocation/:locationId", operatorController.EnableLocation)
		adminRoutePrivate.PUT("/disableLocation/:locationId", operatorController.DisableLocation)

		adminRoutePrivate.PUT("/enableZone/:locationId", operatorController.EnableZone)
		adminRoutePrivate.PUT("/disableZone/:locationId", operatorController.DisableZone)
		adminRoutePrivate.POST("/getRides", rideController.GetRides)
		adminRoutePrivate.GET("/getRideDetail/:rideId", rideController.GetRideDetail)
		adminRoutePrivate.GET("/getRideLocations/:rideId", rideController.GetRideLocations)
		adminRoutePrivate.PUT("/enableVehicleCategory/:categoryId", vehicleTypeController.EnableVehicleTypeCategory)
		adminRoutePrivate.PUT("/disableVehicleCategory/:categoryId", vehicleTypeController.DisableVehicleTypeCategory)
		adminRoutePrivate.POST("/editVehicleTypeCategory", vehicleTypeController.EditVehicleTypeCategory)

		adminRoutePrivate.PUT("/enableDriver/:driverId", driverController.EnableDriver)
		adminRoutePrivate.PUT("/disableDriver/:driverId", driverController.DisableDriver)

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

	customerRoutePublic := router.Group("/customer")
	{
		customerRoutePublic.POST("/sendOtp", passengerController.SendOtp)
		customerRoutePublic.POST("/verifyOtp", passengerController.VerifyOtp)
	}

	customerRoutePrivate := router.Group("/customer")
	customerRoutePrivate.Use(tokenAuthMiddleware("passenger"))
	{
		customerRoutePrivate.POST("/completeProfile", passengerController.AddCustomerBasicDetails)
		customerRoutePrivate.POST("/updateFcm", passengerController.UpdateFcm)
		customerRoutePrivate.POST("/getVehicleTypes", rideBookingController.GetEstimatedFare)
		customerRoutePrivate.POST("/bookRide", rideBookingController.BookRide)
		customerRoutePrivate.POST("/getRides", rideBookingController.GetBookingHistory)
		customerRoutePrivate.POST("/cancelRide", rideBookingController.CancelRide)
		customerRoutePrivate.POST("/checkIsOnRide", passengerController.CheckIsOnRide)
		customerRoutePrivate.POST("/getNearestDrivers", passengerController.GetNearByDrivers)
		customerRoutePrivate.POST("/getRideDetails/:rideId", rideController.GetRideDetailsForMobile)
		customerRoutePrivate.POST("/rateDriver", rideController.RateDriver)
		customerRoutePrivate.POST("/getRideTimeline/:id", rideController.GetRideTimeline)


	}

	driverRoutePublic := router.Group("/driver")
	{
		driverRoutePublic.POST("/sendOtp", driverController.SendOtp)
		driverRoutePublic.POST("/verifyOtp", driverController.VerifyOtp)
		driverRoutePublic.POST("/createDriverAccount", driverController.CreateDriverAccount)
		driverRoutePublic.GET("/getActiveOperators", operatorController.GetActiveOperators)
		driverRoutePublic.GET("/getActiveVehicleTypes", vehicleTypeController.GetActiveVehicleTypes)

	}

	driverRoutePrivate := router.Group("/driver")
	driverRoutePrivate.Use(tokenAuthMiddleware("driver"))
	{
		driverRoutePrivate.GET("/getDriverDocs/:id", operatorController.GetDriverDocs)

		driverRoutePrivate.POST("/goOnline", driverController.GoOnline)
		driverRoutePrivate.POST("/cancelRide", rideBookingController.CancelRideDriver)
		driverRoutePrivate.POST("/getDriverDetails", driverController.GetDriverDetails)
		driverRoutePrivate.POST("/uploadDriverDocument", driverController.UploadDriverDocument)
		driverRoutePrivate.POST("/submitForApproval", driverController.SubmitForApproval)
		driverRoutePrivate.POST("/goOffline", driverController.GoOffline)
		driverRoutePrivate.POST("/acceptRide", rideController.RideAccept)
		driverRoutePrivate.POST("/driverArrived/:rideId", rideController.DriverArrived)
		driverRoutePrivate.POST("/startTrip/:rideId", rideController.StartTrip)
		driverRoutePrivate.POST("/stopTrip/:rideId", rideController.StopTrip)
		driverRoutePrivate.POST("/checkQueuedRide", rideController.CheckRideQueue)
		driverRoutePrivate.POST("/updateRideLocations", rideController.UpdateRideOperators)
		driverRoutePrivate.POST("/getRides", rideBookingController.GetDriverBookingHistory)
		driverRoutePrivate.POST("/getRideDetails/:rideId", rideController.GetRideDetailsForMobile)
		driverRoutePrivate.POST("/ratePassenger", rideController.RatePassenger)
		driverRoutePrivate.POST("/getRideTimeline/:id", rideController.GetRideTimeline)

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
