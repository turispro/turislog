package elastic_client

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"os"
)

var es *elasticsearch.Client

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
