package services

import (
	"context"
	"time"

	"github.com/jcarvallo/golang-api-rest-with-mongodb/db"
	"github.com/jcarvallo/golang-api-rest-with-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = db.DBConnection("users")
var ctx = context.Background()

func GetUsers() (models.Users, error) {

	var users models.Users

	cur, err := collection.Find(ctx, bson.D{})

	for cur.Next(ctx) {
		var user models.User
		cur.Decode(&user)
		users = append(users, &user)
	}

	return users, err
}

func GetUser(id string) (models.User, error) {
	var user models.User

	oid, _ := primitive.ObjectIDFromHex(id)

	cur, err := collection.Find(ctx, bson.D{{Key: "_id", Value: oid}})

	for cur.Next(ctx) {
		cur.Decode(&user)
	}

	return user, err
}

func CreateUser(body models.User) (*mongo.InsertOneResult, error) {

	res, err := collection.InsertOne(ctx, body)

	return res, err
}

func UpdateUser(user models.User, id string) (*mongo.UpdateResult, error) {

	oid, _ := primitive.ObjectIDFromHex(id)

	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_at": time.Now(),
		},
	}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, update)

	return res, err
}

func DeleteUser(id string) error {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return err
	}

	return nil
}
