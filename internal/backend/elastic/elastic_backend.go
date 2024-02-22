package elastic

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
	"github.com/turispro/turislog/model"
	"log"
	"os"
	"strings"
)

var es *elasticsearch.Client

type ElasticBackend struct {
}

func NewElasticBackend() ElasticBackend {
	es = Start()
	return ElasticBackend{}
}

func (eb *ElasticBackend) sendBodyToElastic(body string) {
	id, _ := uuid.NewUUID()
	req := esapi.IndexRequest{
		Index:      os.Getenv("ELASTICSEARCH_INDEX"),
		DocumentID: id.String(),
		Body:       strings.NewReader(body),
		Refresh:    "true",
	}
	_, err := req.Do(context.TODO(), es)
	if err != nil {
		log.Println("Error en request: ", err)
	}
}

func (eb ElasticBackend) Register(message model.Log) {
	body, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	eb.sendBodyToElastic(string(body))
}

func Start() *elasticsearch.Client {
	log.Println(os.Getenv("ELASTICSEARCH"))
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTICSEARCH"),
		},
	}
	var err error
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("error creating client %s", err)
	}
	return es
}
