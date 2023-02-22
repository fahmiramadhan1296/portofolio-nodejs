package models

/*
	@Author: DevProblems(Sarang Kumar)
	@YTChannel: https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ
*/

type IdModel string

type User struct {
	Id IdModel `json:"_id" bson:"_id"`
	Name 	string  `json:"name" bson:"name"`
	Email   string  `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	AuthKey string `json:"authKey" bson:"authKey"`
	IsAuth bool `json:"isAuth" bson:"isAuth"`
}

