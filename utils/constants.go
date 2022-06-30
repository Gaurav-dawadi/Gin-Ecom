package utils

import (
	"time"
)

/*
TODO: Save secret_key in env
fetch from .env and save to some variable
then save to const
*/

const (
	UNAUTHORIZED              string        = "Unauthorized"
	TOKEN_EXPIRED             string        = "Token expired"
	BAD_REQUEST               string        = "Bad Request"
	AUTHORIZATION             string        = "Authorization"
	ACCESS_TOKEN_EXPIRY_TIME  time.Duration = 9 * time.Hour
	REFRESH_TOKEN_EXPIRY_TIME time.Duration = 24 * time.Hour
	SECRET_KEY                string        = "gin-gonic-insecure-(wqfn-ll!0)eq7tzl2fko4)jazq(qi*t@fca0z_a8!n8s&f^8f"
	USER_ID                   string        = "User_Id"
	USER_NAME                 string        = "Username"
)
