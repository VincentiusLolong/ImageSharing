package services

import (
	"context"
	"mestorage/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) Add(a context.Context, file models.ImagePost, account_id primitive.ObjectID) (*mongo.InsertOneResult, error) {
	newfile := &models.ImagePost{
		Account_Id: account_id,
		Type:       file.Title,
		Title:      file.Title,
		Date:       time.Now(),
		Repository: file.Repository,
	}
	result, err := s.db.MongoImage().InsertOne(a, newfile)
	return result, err
}

func (s *service) Get(a context.Context, accountid primitive.ObjectID) ([]models.ImagePost, error) {
	var allfile []models.ImagePost
	cursor, err := s.db.MongoImage().Find(a, bson.M{"accountid": accountid})
	if err != nil {
		return nil, err
	}

	err = cursor.All(a, &allfile)
	if err != nil {
		return nil, err
	}

	return allfile, nil
}
