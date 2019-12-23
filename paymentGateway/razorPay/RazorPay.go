package razorPay

import (
	"fmt"
	"net/http"
	"strings"
	"taxi/models"
	"taxi/shared/database"

	"github.com/gin-gonic/gin"
)

type RazorPay struct {
}

type request struct {
	Event    string   `json:"event"`
	Entity   string   `json:"entity"`
	Contains []string `json:"contains"`
	Payload  struct {
		Payment struct {
			Entity struct {
				ID               string      `json:"id"`
				Entity           string      `json:"entity"`
				Amount           int         `json:"amount"`
				Currency         string      `json:"currency"`
				Status           string      `json:"status"`
				AmountRefunded   int         `json:"amount_refunded"`
				RefundStatus     interface{} `json:"refund_status"`
				Method           string      `json:"method"`
				OrderID          string      `json:"order_id"`
				CardID           string      `json:"card_id"`
				Bank             interface{} `json:"bank"`
				Captured         bool        `json:"captured"`
				Email            string      `json:"email"`
				Contact          string      `json:"contact"`
				Description      string      `json:"description"`
				ErrorCode        interface{} `json:"error_code"`
				ErrorDescription interface{} `json:"error_description"`
				Fee              int         `json:"fee"`
				ServiceTax       int         `json:"service_tax"`
				International    bool        `json:"international"`
				Notes            struct {
					ReferenceNo string `json:"reference_no"`
				} `json:"notes"`
				Vpa    interface{} `json:"vpa"`
				Wallet interface{} `json:"wallet"`
			} `json:"entity"`
		} `json:"payment"`
		CreatedAt int `json:"created_at"`
	} `json:"payload"`
}

func (r *RazorPay) Webhook(c *gin.Context) {
	var data request
	var isPaid = false
	c.BindJSON(&data)
	var ride models.Ride
	var rideId = strings.Split(data.Payload.Payment.Entity.Description, "#")[1]
	fmt.Printf("%+v", data.Payload.Payment.Entity)
	database.Db.Where("id = ?", rideId).First(&ride)
	if data.Payload.Payment.Entity.Status == "authorized" {
		isPaid = true
	}
	if ride.ID != 0 {
		database.Db.Model(&ride).UpdateColumns(&models.Ride{
			IsPaid:        true,
			TransactionID: data.Payload.Payment.Entity.ID,
		})
	}

	c.JSON(http.StatusOK, isPaid)
}
