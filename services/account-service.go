package services

import (
	"context"
	"errors"
	"fmt"
	"mestorage/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) CreateAccount(a context.Context, account models.Account) (string, error) {
	var mails map[string]interface{}

	err := s.db.MongoAccount().FindOne(a, bson.M{"email": account.Email}).Decode(&mails)
	if err == nil {
		str := fmt.Sprintf("Email (%v) Already Registered", mails["email"])
		return "", errors.New(str)
	}

	_, errs := s.db.MongoAccount().InsertOne(a, account)
	if errs != nil {
		return "", errs
	}
	userid := fmt.Sprintf("%v", mails["_id"])
	return userid, errs
}

func (s *service) LoginAccount(a context.Context, userlogin models.SignIn) (*models.Account, error) {
	var myaccount *models.Account
	cursor := s.db.MongoAccount().FindOne(a, bson.M{"email": userlogin.Email}).Decode(&myaccount)

	return myaccount, cursor
}
