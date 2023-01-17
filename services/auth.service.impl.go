package services

import (
	"context"

	"exmaple.com/job-x-apis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewAuthService(userCollection *mongo.Collection, ctx context.Context) AuthService {
	return &AuthServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *AuthServiceImpl) Login(credentials *models.Credentials) (*models.User, error) {
	var user *models.User
	query := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "email", Value: credentials.Email}},
				bson.D{{Key: "password", Value: credentials.Password}},
			},
		},
	}
	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *AuthServiceImpl) CheckStatus() (string) {
	return ""
}
