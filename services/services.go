package services

import (
	"context"
	"mestorage/conf/database"
	"mestorage/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PrivateServices interface {
	Add(a context.Context, file models.ImagePost, account_id primitive.ObjectID) (*mongo.InsertOneResult, error)
	Get(a context.Context, title primitive.ObjectID) ([]models.ImagePost, error)
	CreateAccount(a context.Context, account models.Account) (string, error)
	LoginAccount(a context.Context, userlogin models.SignIn) (*models.Account, error)
}

// type PubilcSerivces interface {
// }

type service struct {
	db database.Dbs
}

func PrivateSerivce(db database.Dbs) PrivateServices {
	return &service{
		db: db,
	}
}

// func PubilcSerivce(db database.Dbs) PubilcSerivces {
// 	return &service{
// 		db: db,
// 	}
// }
