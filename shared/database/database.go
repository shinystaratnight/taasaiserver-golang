package database

import (
	"log"
	"os"
	"taxi/models"

	"github.com/jinzhu/gorm"
)

var (
	Db *gorm.DB
)

func SetupDb() {
	var err error
	Db, err = gorm.Open("postgres", "host=35.193.214.177 port=25060 user=test dbname=taasai password=test#123 ")
	if err != nil {
		panic(err.Error())
		os.Exit(1)
	} else {
		log.Println("PostgresDb Connected Successfully")
	}
	Db.AutoMigrate(
		&models.Otp{},
		&models.Admin{},
		&models.VehicleCategory{},
		&models.VehicleType{},
		&models.Location{},
		&models.Fare{},
		&models.Zone{},
		&models.ZoneFare{},
	)
	Db.Exec("CREATE EXTENSION postgis;")
	Db.Exec("ALTER TABLE locations ADD COLUMN polygon geometry;")
	Db.Exec("ALTER TABLE zones ADD COLUMN polygon geometry;")

	Db.Exec("ALTER TABLE vehicles ADD COLUMN latlng geometry;")
	Db.AutoMigrate(
		&models.Company{},
		&models.CompanyLocationAssignment{},
		&models.Driver{},
	)
	Db.AutoMigrate(
		&models.Vehicle{},
		&models.DriverVehicleAssignment{},
		&models.Passenger{},
		&models.Ride{},
		&models.SentRideRequest{},
		&models.RideLocation{},
		&models.EmergencyContact{},
	)
	Db.Exec("ALTER TABLE ride_locations ADD COLUMN latlng geometry;")
	Db.Exec("CREATE INDEX ride_locations_latlng_idx ON ride_locations USING gist (latlng);")

	Db.LogMode(true)
}
