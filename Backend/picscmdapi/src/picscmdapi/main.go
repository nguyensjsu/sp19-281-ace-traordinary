/*
	picscmdapi REST API (Version 1)
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/glog"
	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	// kafka
	kafkaBrokerUrl     string
	kafkaVerbose       bool
	kafkaTopic         string
	kafkaConsumerGroup string
	kafkaClientId      string
)

func main() {
	flag.StringVar(&kafkaBrokerUrl, "kafka-brokers", "http://35.163.127.180:9092", "Kafka brokers in comma separated value")
	flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaTopic, "kafka-topic", "picassa", "Kafka topic. Only one topic per worker.")
	flag.StringVar(&kafkaConsumerGroup, "kafka-consumer-group", "consumer-group", "Kafka consumer group")
	flag.StringVar(&kafkaClientId, "kafka-client-id", "my-client-id", "Kafka client id")

	Urls := []string{"35.163.127.180:9092"}
	flag.Parse()
	kafkaProducer, err := Configure(Urls, kafkaClientId, kafkaTopic)
	fmt.Println(kafkaProducer)
	if err != nil {
		glog.Error("error unable to configure kafka", err)
		return
	}
	defer kafkaProducer.Close()
	router := mux.NewRouter()
	headersOk := h.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := h.AllowedOrigins([]string{"*"})
	methodsOk := h.AllowedMethods([]string{"HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	router.HandleFunc("/images", UploadPictureHandler).Methods("POST")
	router.HandleFunc("/images/:imageid", UpdatePictureHandler).Methods("PUT")
	router.HandleFunc("/images/{imageid}", DeletePictureHandler).Methods("DELETE")
	router.HandleFunc("/ping", PingHandler).Methods("GET")
	fmt.Println("Starting server on port 8001...")
	log.Fatal(http.ListenAndServe(":8001", h.CORS(headersOk, methodsOk, originsOk)(router)))
}
