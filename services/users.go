package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Glitchfix/golagobar/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// LoginService - Login service golang
func LoginService(ctx context.Context, userLogin models.Login) (result models.Login, err error) {
	filter := bson.M{
		"phone": *userLogin.Phone,
		"pass":  *userLogin.Password,
		"role":  *userLogin.Role,
	}
	err = userCollection.FindOne(ctx, filter).Decode(&result)
	return
}

// GenerateToken - Generate token service golang
func GenerateToken(ctx context.Context, profile models.Login) (token string, err error) {
	token = uuid.New().String()
	data := bson.M{
		"token": token,
		"phone": profile.Phone,
		"role":  profile.Role,
	}
	_, err = authCollection.InsertOne(ctx, data)
	return
}

// RegisterService - Login service golang
func RegisterService(ctx context.Context, newUser models.Profile) (result *mongo.InsertOneResult, err error) {
	result, err = userCollection.InsertOne(ctx, newUser)
	if nil != err {
		return
	}
	return
}

// FindUser - Find user
func FindUser(ctx context.Context, newUser models.Profile) (user models.Profile, err error) {
	err = userCollection.FindOne(ctx, bson.M{
		"phone": *newUser.Phone,
		"role":  *newUser.Role,
	}).Decode(&user)
	return
}

// UserAlreadyExists - Check if user already exists
func UserAlreadyExists(ctx context.Context, newUser models.Profile) (err error) {
	err = userCollection.FindOne(ctx, bson.M{
		"phone": *newUser.Phone,
		"role":  *newUser.Role,
	}).Err()
	if nil != err {
		return nil
	}
	err = errors.New(fmt.Sprint(newUser.Phone, " already exists"))
	return
}

// ProfileService - get profile information
func ProfileService(ctx context.Context, phone string, role int) (profile models.Profile, err error) {
	err = userCollection.FindOne(ctx, bson.M{
		"phone": phone,
		"role":  role,
	}).Decode(&profile)
	return
}

// SetLocation - set profile location
func SetLocation(ctx context.Context, phone string, coords models.Coordinates) (err error) {
	return userCollection.FindOneAndUpdate(
		ctx,
		bson.M{
			"phone": phone,
		},
		bson.M{
			"$set": bson.M{
				"location.coordinates": coords,
			},
		},
	).Err()
}

// EditPassword - edit profile password
func EditPassword(ctx context.Context, phone string, role int32, oldPass, newPass string) (err error) {
	update := bson.M{
		"phone": phone,
		"role":  role,
		"pass":  oldPass,
	}
	return userCollection.FindOneAndUpdate(
		ctx,
		update,
		bson.M{
			"$set": bson.M{
				"pass": newPass,
			},
		},
	).Err()
}
