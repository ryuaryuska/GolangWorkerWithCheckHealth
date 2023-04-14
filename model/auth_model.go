package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string             `json:"access_token"`
	RefreshToken string             `json:"refresh_token"`
	UserID       primitive.ObjectID `json:"user_id"`
	Username     string             `json:"username"`
	RoleID       primitive.ObjectID `json:"role_id"`
	RoleName     string             `json:"role_name"`
}

type JwtPayload struct {
	UserID      primitive.ObjectID `json:"user_id"`
	Username    string             `json:"username"`
	RoleID      primitive.ObjectID `json:"role_id"`
	RoleName    string             `json:"role_name"`
	AccessUuid  string             `json:"access_uuid"`
	RefreshUuid string             `json:"refresh_uuid"`
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type UpdateLoginFailed struct {
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
}
