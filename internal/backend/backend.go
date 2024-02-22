package backend

import (
	"github.com/turispro/turislog/internal/backend/elastic"
	"github.com/turispro/turislog/internal/backend/mongo"
	"github.com/turispro/turislog/model"
)

const (
	MONGO   = "mongo"
	ELASTIC = "elastic"
)

type Backend interface {
	Register(message model.Log)
}

func NewBackend(selection string) Backend {
	if selection == MONGO {
		return mongo.NewMongoBackend()
	} else {
		return elastic.NewElasticBackend()
	}
}
