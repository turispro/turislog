package model

import "time"

type Log struct {
	Timestamp time.Time   `json:"timestamp"`
	Level     string      `json:"level"`
	Service   string      `json:"service"`
	User      string      `json:"user"`
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
