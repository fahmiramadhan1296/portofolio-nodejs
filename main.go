package main

import (
	"log"
	"my-portofolio/mongoDB/config"
)

/*
	@Author: DevProblems(Sarang Kumar)
	@YTChannel: https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ
*/

// var mongoUrl = "mongodb+srv://my_portofolio_2022:Fahmi120296@cluster0.akwywmv.mongodb.net/?retryWrites=true&w=majority"

func init() {
}

func main() {
	server, ctx, mongoclient, uc := config.Init()
	basepath := server.Group("/v1")
	defer mongoclient.Disconnect(ctx)
	uc.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":9090"))
}