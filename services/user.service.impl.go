package services

import (
	"context"
	"errors"

	"exmaple.com/job-x-apis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserSevice(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.userCollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.userCollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_address", Value: user.Address}, bson.E{Key: "user_age", Value: user.Age}}}}

	result, _ := u.userCollection.UpdateOne(u.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("NO MATCH FOUND")
	}
	return nil
}
func (u *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	result, _ := u.userCollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("NO MATCH FOUND")
	}
	return nil
}
