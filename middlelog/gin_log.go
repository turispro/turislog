package middlelog

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/turispro/turislog/elastic_client"
	"github.com/turispro/turislog/model"
	"log"
	"os"
	"strings"
	"time"
)

const (
	DEBUG string = "DEBUG"
	INFO         = "INFO"
	WARN         = "WARN"
	ERROR        = "ERROR"
	FATAL        = "FATAL"
)

var client *elasticsearch.Client

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
	log.Println("client elastic ")
	client = elastic_client.Start()
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
			Username:   context.GetHeader("X-Usuario-Nombre"),
			Id:         context.GetHeader("X-Usuario-Id"),
			Sucursal:   context.GetHeader("X-Sucursal-Nombre"),
			SucursalId: context.GetHeader("X-Sucursal-Id"),
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
		jString, _ := json.Marshal(register)
		go sendBodyToElastic(string(jString))
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
	body, _ := json.Marshal(register)
	go sendBodyToElastic(string(body))
}

func sendBodyToElastic(body string) {
	id, _ := uuid.NewUUID()
	req := esapi.IndexRequest{
		Index:      "system-logs",
		DocumentID: id.String(),
		Body:       strings.NewReader(body),
		Refresh:    "true",
	}
	_, err := req.Do(context.TODO(), client)
	if err != nil {
		log.Println("Error en request: ", err)
	}
}
