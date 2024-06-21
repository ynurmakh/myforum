package models

import "errors"

var (
	CreateUser_NotUniqEmail    error = errors.New("Select another email")
	CreateUser_NotUniqNickName error = errors.New("Select another Name")
)
