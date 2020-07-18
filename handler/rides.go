package handler

import (
	"net/http"

	"github.com/Glitchfix/golagobar/models/constants"

	"github.com/Glitchfix/golagobar/models"
	"github.com/Glitchfix/golagobar/services"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// RdieEstimateHandler - Create ride
func RdieEstimateHandler(c *gin.Context) {
	role, _ := c.Get("role")
	if role != constants.RoleRider {
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, services.FareCalculate())
}

// NearbyRidesHandler - Create ride
func NearbyRidesHandler(c *gin.Context) {
	role, _ := c.Get("role")
	if role != constants.RoleRider {
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}
	var coordinates models.Coordinates
	err := c.Bind(&coordinates)
	if nil != err || len(coordinates) != 2 {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}
	print(coordinates)
	locations, err := services.NearbyRide(c.Request.Context(), coordinates)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusInternalServerError)
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, locations)
}

// CreateRideHandler - Create ride
func CreateRideHandler(c *gin.Context) {
	var newRide models.Ride

	phone, _ := c.Get("phone")
	role, _ := c.Get("role")
	if role != constants.RoleRider {
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}
	err := c.Bind(&newRide)
	logrus.Infoln(newRide)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}
	rider := phone.(string)
	newRide.ID = new(int)
	newRide.Rider = &rider
	newRide.Provider = new(string)
	_, err = govalidator.ValidateStruct(newRide)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}
	err = services.CreateRide(c.Request.Context(), &newRide)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusInternalServerError)
		c.Next()
		return
	}

	newRide.Provider = nil
	c.AbortWithStatusJSON(http.StatusOK, newRide)
}

// AcceptRideHandler - Accept ride
func AcceptRideHandler(c *gin.Context) {
	var ride models.Ride

	phone, _ := c.Get("phone")
	role, _ := c.Get("role")
	if role != constants.RoleProvider {
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}

	err := c.Bind(&ride)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}

	provider := phone.(string)
	ride.Provider = &provider
	ride.Status = constants.RideAccepted
	err = services.EditRide(c.Request.Context(), &ride)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusInternalServerError)
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, ride)
}

// CancelRideHandler - Cancel ride
func CancelRideHandler(c *gin.Context) {
	var ride models.Ride

	phone, _ := c.Get("phone")
	role, _ := c.Get("role")
	if role.(int) != constants.RoleProvider && role.(int) != constants.RoleRider {
		c.Set("status", http.StatusUnauthorized)
		c.Next()
		return
	}

	err := c.Bind(&ride)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}

	provider := phone.(string)
	ride.Provider = &provider
	ride.Rider = &provider
	ride.Status = constants.RideCancelled
	ride.Pickup = new(models.Location)
	ride.Drop = ride.Pickup
	_, err = govalidator.ValidateStruct(ride)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}
	err = services.EditRide(c.Request.Context(), &ride)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusUnprocessableEntity)
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, ride)
}

// CompleteRideHandler - Completed ride
func CompleteRideHandler(c *gin.Context) {
	var ride models.Ride

	phone, _ := c.Get("phone")
	role, _ := c.Get("role")
	if role != constants.RoleProvider {
		c.Set("status", http.StatusUnauthorized)
		c.Next()
		return
	}

	err := c.Bind(&ride)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}

	provider := phone.(string)
	ride.Rider = new(string)
	ride.Provider = &provider
	ride.Status = constants.RideCompleted
	ride.Pickup = new(models.Location)
	ride.Drop = ride.Pickup
	_, err = govalidator.ValidateStruct(ride)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}
	err = services.EditRide(c.Request.Context(), &ride)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusInternalServerError)
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, ride)
}
