package middlewares

import (
	"net/http"
	"strings"

	"github.com/Glitchfix/golagobar/db"
	"github.com/Glitchfix/golagobar/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var authCollection *mongo.Collection

// Initialize config
func Init() {
	authCollection = db.GetCollection("auth")
}

// AuthMiddleware - authenticate endpoints-
func AuthMiddleware(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		c.AbortWithStatusJSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"success": false,
			"text":    "Token missing",
		})
		return
	}

	reqToken = strings.TrimSpace(splitToken[1])
	// Token filter
	filter := bson.M{
		"token": reqToken,
	}
	var result models.Profile
	err := authCollection.FindOne(c.Request.Context(), filter).Decode(&result)
	if nil != err {
		logrus.Errorln(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"text":    "Unauthorized request",
		})
		return
	}
	c.Set("phone", *result.Phone)
	c.Set("role", *result.Role)
	logrus.Infoln("phone=", *result.Phone)

	c.Next()
}
