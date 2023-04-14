package model

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"errors"`
}

type GeneralError struct {
	General string `json:"general"`
}

type NotValidImage struct {
	Image string `json:"image"`
}

const DateTimeLayout = "02-01-2006 15:04"
