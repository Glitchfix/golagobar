package handler

import (
	"net/http"

	"github.com/Glitchfix/golagobar/models"
	"github.com/Glitchfix/golagobar/services"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoginHandler - Login Handler
func LoginHandler(c *gin.Context) {
	var userLogin models.Login

	err := c.Bind(&userLogin)
	logrus.Infoln(userLogin)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}
	_, err = govalidator.ValidateStruct(userLogin)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}
	print("still")
	result, err := services.LoginService(c.Request.Context(), userLogin)
	if nil != err {
		c.Set("status", http.StatusUnauthorized)
		c.Next()
		return
	}

	token, err := services.GenerateToken(c.Request.Context(), result)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusUnauthorized)
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
		"role":    result.Role,
	})
}

// RegisterHandler - Register Handler
func RegisterHandler(c *gin.Context) {
	var userRegister models.Profile

	err := c.Bind(&userRegister)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}

	_, err = govalidator.ValidateStruct(userRegister)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}

	err = services.UserAlreadyExists(c.Request.Context(), userRegister)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusUnprocessableEntity)
		c.Next()
		return
	}

	result, err := services.RegisterService(c.Request.Context(), userRegister)
	logrus.Infoln(result)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusUnprocessableEntity)
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": true,
		"text":    "Registration successful",
	})
}

// ProfileHandler - Fetch user profile
func ProfileHandler(c *gin.Context) {
	claimAuthor, _ := c.Get("phone")
	claimRole, _ := c.Get("role")

	profile, err := services.ProfileService(c.Request.Context(), claimAuthor.(string), claimRole.(int))
	if nil != err {
		c.Set("status", http.StatusUnprocessableEntity)
		c.Next()
		return
	}
	profile.Password = nil
	profile.Location = nil
	c.AbortWithStatusJSON(http.StatusOK, profile)
}

// SetLocation - Fetch user profile
func SetLocation(c *gin.Context) {
	claimAuthor, _ := c.Get("phone")
	var coordinates models.Coordinates
	err := c.BindJSON(&coordinates)
	if nil != err || len(coordinates) != 2 {
		logrus.Errorln(err)
		c.Set("status", http.StatusBadRequest)
		c.Next()
		return
	}

	err = services.SetLocation(c.Request.Context(), claimAuthor.(string), coordinates)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusUnprocessableEntity)
		c.Next()
		return
	}
	c.AbortWithStatus(http.StatusOK)
}

// EditPassword - Edit user password
func EditPassword(c *gin.Context) {
	claimAuthor, _ := c.Get("phone")
	claimRole, _ := c.Get("role")
	var editPass struct {
		OldPassword string `json:"old,omitempty"`
		NewPassword string `json:"new,omitempty"`
	}
	err := c.BindJSON(&editPass)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusUnprocessableEntity)
		c.Next()
		return
	}

	err = services.EditPassword(c.Request.Context(), claimAuthor.(string), int32(claimRole.(int)), editPass.OldPassword, editPass.NewPassword)
	if nil != err {
		logrus.Errorln(err)
		c.Set("status", http.StatusUnprocessableEntity)
		c.Next()
		return
	}
	c.AbortWithStatus(http.StatusOK)
}
