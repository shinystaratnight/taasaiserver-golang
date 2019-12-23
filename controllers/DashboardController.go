package controllers

import (
	"net/http"
	"taxi/shared/database"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
}

func (d *DashboardController) GetDataCount(c *gin.Context) {
	type rideStatusCount struct {
		CompletedCount         int64
		DriverUnavailableCount int64
		CancelledCount         int64
	}
	type weeklyReportItem struct {
		Date       string
		Count      int64
		RideStatus int64
	}
	type responseData struct {
		RideCount         int64
		VehicleCount      int64
		DriverCount       int64
		CompanyCount      int64
		PassengerCount    int64
		LocationCount     int64
		VehicleTypeCount  int64
		TotalRideEarnings float64
		RideStatusCount   rideStatusCount
		WeeklyReport      []weeklyReportItem
	}

	var result responseData
	var statusCountData rideStatusCount
	var lastweekRideCount []weeklyReportItem

	database.Db.Raw("SELECT COUNT(*)as ride_count,SUM(rides.total_fare) as total_ride_earnings,(SELECT COUNT(*) as company_count FROM companies),(SELECT COUNT(*) as driver_count FROM drivers ),(SELECT COUNT(*)as vehicle_count FROM vehicles) ,(SELECT COUNT(*) as passenger_count FROM passengers ),(SELECT COUNT(*)as location_count FROM locations) ,(SELECT COUNT(*)as vehicle_type_count FROM vehicle_types) FROM rides").Scan(&result)
	database.Db.Raw("SELECT  COUNT(case when ride_status = 4 then id end) as completed_count,COUNT(case when ride_status = 5 then id end) as driver_unavailable_count,COUNT(case when ride_status = 6 then id end) as cancelled_count FROM rides ").Scan(&statusCountData)
	database.Db.Raw("select Count(*),ride_status,(created_at::date)::text AS date FROM rides where ride_status in (4,5,6) AND created_at >= date_trunc('week',current_date) group by date,ride_status").Scan(&lastweekRideCount)

	result.RideStatusCount = statusCountData
	result.WeeklyReport = lastweekRideCount

	c.JSON(http.StatusOK, result)
}
