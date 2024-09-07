package models

import "net/http"

type Result struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

func Success(data any) Result {
	return Result{
		Code: http.StatusOK,
		Data: data,
	}
}

func Error(msg string) Result {
	return Result{
		Code: http.StatusOK,
		Data: msg,
	}
}
