package common

import (
	"Agent-Demo/src/main/models"
	"errors"
	"time"
)

func NewSuccessResp(data any) models.RespModel {
	return models.RespModel{
		Code: 1,
		Val:  data,
		Time: time.Now().String(),
	}
}

func NewErrorResp(msg string) models.RespModel {
	return models.RespModel{
		Code: 0,
		Val:  errors.New(msg),
		Time: time.Now().String(),
	}
}
