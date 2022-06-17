package utils

import (
	"time"
)

const (
	UNAUTHORIZED              string        = "Unauthorized"
	TOKEN_EXPIRED             string        = "Token expired"
	BAD_REQUEST               string        = "Bad Request"
	AUTHORIZATION             string        = "Authorization"
	ACCESS_TOKEN_EXPIRY_TIME  time.Duration = 1 * time.Minute
	REFRESH_TOKEN_EXPIRY_TIME time.Duration = 24 * time.Hour
	SECRET_KEY                string        = "your-256-bit-secret"
)
