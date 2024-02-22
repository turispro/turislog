package middlelog

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/turispro/turislog/internal/backend"
	"github.com/turispro/turislog/model"
	"log"
	"os"
	"time"
)

const (
	DEBUG string = "DEBUG"
	INFO         = "INFO"
	WARN         = "WARN"
	ERROR        = "ERROR"
	FATAL        = "FATAL"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
}

func logLevel(status int) string {
	switch {
	case status < 400:
		return INFO
	case status >= 400 && status < 500:
		return ERROR
	case status >= 500:
		return FATAL
	}
	return INFO
}

func GinLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Next()
		var body any
		_ = json.NewDecoder(context.Request.Body).Decode(&body)
		user := model.User{
			Username:         context.GetHeader("X-User-Name"),
			Id:               context.GetHeader("X-User-Id"),
			TourOperatorId:   context.GetHeader("X-Tour-Operator-Id"),
			TourOperatorName: context.GetHeader("X-Tour-Operator-Name"),
		}
		register := model.Log{
			Timestamp: time.Now(),
			Message:   "request",
			Level:     logLevel(context.Writer.Status()),
			Service:   os.Getenv("SERVICE_NAME"),
			User:      user,
			Request: model.LogRequest{
				Method: context.Request.Method,
				Body:   body,
				Url:    context.Request.URL.String(),
			},
			Response: model.LogResponse{
				ResponseTime: time.Since(t).Microseconds(),
				Status:       context.Writer.Status(),
			},
		}
		go sendToBackend(register)
	}
}

func Debug(message string) {
	sendMessage(message, DEBUG, nil)
}

func Warn(message string) {
	sendMessage(message, WARN, nil)
}

func Error(message string) {
	sendMessage(message, ERROR, nil)
}

func Fatal(message string) {
	sendMessage(message, FATAL, nil)
}

func Info(message string) {
	sendMessage(message, INFO, nil)
}

func DebugWithUser(message string, user *model.User) {
	sendMessage(message, DEBUG, user)
}

func WarnWithUser(message string, user *model.User) {
	sendMessage(message, WARN, user)
}

func ErrorWithUser(message string, user *model.User) {
	sendMessage(message, ERROR, user)
}

func FatalWithUser(message string, user *model.User) {
	sendMessage(message, FATAL, user)
}

func InfoWithUser(message string, user *model.User) {
	sendMessage(message, INFO, user)
}

func sendMessage(message, level string, user *model.User) {
	var logUser model.User
	if user == nil {
		logUser = model.User{}
	} else {
		logUser = *user
	}
	register := model.Log{
		Timestamp: time.Now(),
		Level:     level,
		Service:   os.Getenv("SERVICE_NAME"),
		User:      logUser,
		Message:   message,
	}
	go sendToBackend(register)
}

func sendToBackend(message model.Log) {
	be := backend.NewBackend(os.Getenv("BACKEND"))
	be.Register(message)
}
