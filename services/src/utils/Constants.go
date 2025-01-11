package utils

import "errors"

//env
const (
	LOCAL = "local"
	DEV   = "dev"
	PROD  = "prod"
)

const (
	JWT_NAME = "name"
	JWT_ID   = "id"
	JWT_EXP  = "exp"
	JWT_IAT  = "iat"
	JWT_SRC  = "src"
)

const (
	TOKEN_TYPE = "Bearer"
)

const (
	STATUS_INPROGRESS = "in-progress"
	STATUS_PENDING    = "pending"
	STATUS_COMPLETED  = "completed"
)

var (
	NotFoundErr = errors.New("not found")
)
