package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	Id              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Account_Created time.Time          `json:"time,omitempty"`
	Username        string             `json:"username" binding:"required"`
	Email           string             `json:"email" binding:"required,email"`
	Password        string             `json:"password" binding:"required"`
	Profile         Profile            `json:"profile,omitempty"`
}

type Profile struct {
	Bio      string `json:"bio,omitempty"  binding:"max=60"`
	Location string `json:"location,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Age      int8   `json:"age" binding:"required,gte=12,lte=125"`
}

type SignIn struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ImagePost struct {
	Account_Id primitive.ObjectID `json:"accountid" binding:"required"`
	Type       string             `json:"type" binding:"required"`
	Title      string             `json:"title" binding:"min=2,max=10"`
	Date       time.Time          `json:"date,omitempty"`
	Repository string             `json:"repository" binding:"required"`
}

type Comments struct {
	File_Id    primitive.ObjectID `json:"image_id" binding:"required"`
	Account_Id primitive.ObjectID `json:"accountid" binding:"required"`
	Comments   string             `json:"comments" binding:"max=100"`
}

type Account_Log struct {
	Account_Id      primitive.ObjectID `json:"accountid" binding:"required"`
	Email_Edited    time.Time          `json:"time_email,omitempty"`
	Password_Edited time.Time          `json:"time_password,omitempty"`
	Profile_Edited  time.Time          `json:"time_profile,omitempty"`
}

// type UserToken struct {
// 	Account_id primitive.ObjectID `json:"accountid" binding:"required"`
// }
