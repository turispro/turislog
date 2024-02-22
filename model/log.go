package model

import "time"

type User struct {
	Username         string `json:"username" mongo:"username"`
	Id               string `json:"id,omitempty" mongo:"id"`
	TourOperatorName string `json:"tourOperatorName,omitempty" mongo:"tourOperatorName,omitempty"`
	TourOperatorId   string `json:"tourOperatorId,omitempty" mongo:"tourOperatorId,omitempty"`
}

type Log struct {
	Timestamp time.Time   `json:"timestamp" mongo:"timestamp"`
	Level     string      `json:"level" mongo:"level"`
	Service   string      `json:"service" mongo:"service"`
	User      User        `json:"user" mongo:"user"`
	Message   string      `json:"message" mongo:"message"`
	Request   LogRequest  `json:"request" mongo:"request"`
	Response  LogResponse `json:"response" mongo:"response"`
}

type LogRequest struct {
	Method string `json:"method" mongo:"method"`
	Body   any    `json:"body,omitempty" mongo:"body,omitempty"`
	Url    string `json:"url" mongo:"url"`
}

type LogResponse struct {
	Status       int   `json:"status" mongo:"status"`
	ResponseTime int64 `json:"responseTime" mongo:"responseTime"`
}
