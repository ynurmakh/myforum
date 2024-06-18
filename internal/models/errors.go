package models

import "errors"

var (
	CreateUser_NotUniqEmail    error = errors.New("err: users email not uniq")
	CreateUser_NotUniqNickName error = errors.New("err: users email not uniq")
)
