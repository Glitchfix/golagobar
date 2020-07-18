package services

import (
	"context"
	"math/rand"
	"time"

	"github.com/Glitchfix/golagobar/models/constants"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Glitchfix/golagobar/models"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

// nextRideID - get the next ride ID
func nextRideID(ctx context.Context) (sequenceID *int, err error) {
	var sequence models.RidesSequence
	filter := bson.M{
		"rideID": "rideID",
	}
	update := bson.M{
		"$inc": bson.M{
			"sequence": 1,
		},
	}

	err = ridesCollection.FindOneAndUpdate(ctx, filter, update).Decode(&sequence)
	sequenceID = &sequence.Sequence
	return
}

// NearbyRide - create a new ride
func NearbyRide(ctx context.Context, location models.Coordinates) (locations []models.Location, err error) {
	filter := bson.M{
		"near":          bson.M{"type": "Point", "coordinates": location},
		"distanceField": "distance",
		"minDistance":   5 * constants.KMPerMile,
		"spherical":     true,
		"query":         bson.M{"role": constants.RoleProvider},
	}
	locations = []models.Location{}
	cursor, err := userCollection.Aggregate(
		ctx,
		[]bson.M{
			{
				"$geoNear": filter,
			},
			{"$limit": 10},
			{
				"$project": bson.M{
					"_id":         0,
					"coordinates": "$location.coordinates",
				},
			},
		})
	if nil != err {
		logrus.Errorln(err)
		return
	}
	err = cursor.All(ctx, &locations)
	return
}

// CreateRide - create a new ride
func CreateRide(ctx context.Context, ride *models.Ride) (err error) {
	ride.ID, err = nextRideID(ctx)
	if nil != err {
		logrus.Errorln(err)
		return
	}
	ride.StartTime = time.Now()
	ride.ModifiedTime = time.Now()
	ride.Fare = FareCalculate()
	ride.Status = constants.OpenRide
	_, err = ridesCollection.InsertOne(ctx, ride)
	return
}

// EditRide - Edit a new ride
func EditRide(ctx context.Context, ride *models.Ride) (err error) {
	filter := bson.M{
		"_id": *ride.ID,
	}
	update := bson.M{
		"$set": bson.M{
			"modified": time.Now(),
			"status":   ride.Status,
		},
	}

	switch ride.Status {
	case constants.RideAccepted:
		filter["status"] = constants.OpenRide
		update["$set"].(bson.M)["provider"] = ride.Provider
	case constants.RideCancelled:
		filter["status"] = constants.OpenRide
		filter["$or"] = []bson.M{
			{
				"rider": ride.Rider,
			},
			{
				"provider": ride.Provider,
			},
		}
	case constants.RideCompleted:
		filter["status"] = constants.RideAccepted
		filter["provider"] = ride.Provider
	}

	logrus.Infoln(filter)
	logrus.Infoln(update)

	return ridesCollection.FindOneAndUpdate(
		ctx,
		filter,
		update,
		&options.FindOneAndUpdateOptions{
			ReturnDocument: &returnAfter,
		}).Decode(&ride)
}

// FareCalculate - get fare for ride
func FareCalculate() (fare models.Fare) {
	rand.Seed(time.Now().UnixNano())
	return models.Fare{
		ChargesPerKM:  float64(rand.Intn((20 - 7 + 1) + 7)),
		ChargesPerMin: float64(rand.Intn((3 - 1 + 1) + 1)),
	}
}
