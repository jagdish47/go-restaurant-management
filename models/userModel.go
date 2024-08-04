package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID 	`bson:"_id"`
	FirstName    *string             	`json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string             	`json:"last_name" validate:"required,min=2,max=100"`
	Password     *string             	`json:"password" validate:"required,min=6"`
	Email        *string             	`json:"email" validate:"required,email"`
	Avatar       *string             	`json:"avatar,omitempty"`
	Phone        *string             	`json:"phone,omitempty"`
	Token        *string             	`json:"token,omitempty"`
	RefreshToken *string             	`json:"refresh_token,omitempty"`
	CreatedAt    time.Time          	`json:"created_at"`
	UpdatedAt    time.Time          	`json:"updated_at"`
	UserID       string             	`json:"user_id"`
}
