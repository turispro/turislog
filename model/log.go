package model

import "time"

type User struct {
	Username         string `json:"username"`
	Id               string `json:"id,omitempty"`
	TourOperatorName string `json:"tourOperatorName,omitempty"`
	TourOperatorId   string `json:"tourOperatorId,omitempty"`
}

type Log struct {
	Timestamp time.Time   `json:"timestamp"`
	Level     string      `json:"level"`
	Service   string      `json:"service"`
	User      User        `json:"user"`
	Message   string      `json:"message"`
	Request   LogRequest  `json:"request"`
	Response  LogResponse `json:"response"`
}

type LogRequest struct {
	Method string `json:"method"`
	Body   any    `json:"body,omitempty"`
	Url    string `json:"url"`
}

type LogResponse struct {
	Status       int   `json:"status"`
	ResponseTime int64 `json:"responseTime"`
}
