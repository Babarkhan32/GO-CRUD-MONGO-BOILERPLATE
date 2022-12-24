package main

import (
	"context"
	"fmt"
	"log"

	controllers "exmaple.com/job-x-apis/contorllers"
	"exmaple.com/job-x-apis/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	usercontroller controllers.UserController
	userservice    services.UserService
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)

	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Db Connection established")
	usercollection = (*mongo.Collection)(mongoclient.Database("db_jobX").Collection("users"))
	userservice = services.NewUserSevice(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)
	basePath := server.Group("/v1")
	usercontroller.RegisterRoutes(basePath)
	log.Fatal(server.Run(":9090"))
}
