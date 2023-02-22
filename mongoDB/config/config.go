package config

import (
	"context"
	"fmt"
	"log"
	"my-portofolio/controllers"
	"my-portofolio/services"
	"net/url"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*
	@Author: DevProblems(Sarang Kumar)
	@YTChannel: https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ
*/
var (
	server      *gin.Engine
	us          services.UserService
	uc          controllers.UserController
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)


func Init() (*gin.Engine, context.Context, *mongo.Client, controllers.UserController){
	username := "my_portofolio_2022"
	password := "Fahmi120296"
	cluster := "cluster0.akwywmv.mongodb.net"
	retryWrites  := "true"
	w := "majority"

	mongoUrl := "mongodb+srv://" + url.QueryEscape(username) + ":" + 
		url.QueryEscape(password) + "@" + cluster + 
		"/?retryWrites=" + retryWrites +
		"&w=" + w
	ctx = context.TODO()
	fmt.Println("mongoUrl", mongoUrl)

	mongoconn := options.Client().ApplyURI(mongoUrl)
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("my_portofolio_2022").Collection("users")
	us = services.NewUserService(userc, ctx)
	uc = controllers.New(us)
	server = gin.Default()

	return server, ctx, mongoclient, uc
}
