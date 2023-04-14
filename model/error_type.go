package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	NOT_BLANK_ERR_TYPE              = "NOT_BLANK"
	MUST_NUMBER_ERR_TYPE            = "MUST_NUMBER"
	MUST_STRING_ERR_TYPE            = "MUST_STRING"
	NOT_VALID_ERR_TYPE              = "NOT_VALID"
	NOT_MATCH_ERR_TYPE              = "NOT_MATCH"
	NOT_FOUND_ERR_TYPE              = "NOT_FOUND"
	USERNAME_NOT_FOUND              = "USERNAME_NOT_FOUND"
	ALREADY_EXIST_ERR_TYPE          = "ALREADY_EXIST"
	FORBIDDEN_ERR_TYPE              = "FORBIDDEN"
	AUTHENTICATION_FAILURE_ERR_TYPE = "AUTHENTICATION_FAILURE"
	USER_BLOCKED                    = "USER_BLOCKED"
	INTERNAL_ERROR_ERR_TYPE         = "INTERNAL_ERROR"
)

func Min(x int) string {
	return fmt.Sprintf("MIN_%v", x)
}

func Max(x int) string {
	return fmt.Sprintf("MAX_%v", x)
}

var NilID, _ = primitive.ObjectIDFromHex("000000000000000000000000")
