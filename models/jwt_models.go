package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTClaim struct {
	Email string `json:"email,omitempty" validate:"required"`
	Name  string `json:"name,omitempty" validate:"required"`
	jwt.StandardClaims
}

type JWTRefreshClaim struct {
	jwt.StandardClaims
}

type UserToken struct {
	Id       primitive.ObjectID `json:"_id,omitempty"`
	Email    string             `json:"email,omitempty"`
	Username string             `json:"password,omitempty"`
}

type RefreshDataToken struct {
	Id primitive.ObjectID `json:"id,omitempty"`
}

type SessionData struct {
	Id          string
	AccessToken string
}
