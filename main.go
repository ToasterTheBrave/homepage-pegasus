package main

import (
	"github.com/truppert/pegasus"
	"github.com/Shopify/sarama"
	"fmt"
	"flag"
	"os"
)

var (
	brokers = flag.String("brokers", "kafka:9092", "Comma separated string of kafka brokers.  Ex: kafka:9092,127.0.0.1:9092")
	topic = flag.String("topic", "", "Kafka topic to watch")
	dockercloud_user = flag.String("dockercloud_user", "", "Docker Cloud Username")
	dockercloud_apikey = flag.String("dockercloud_apikey", "", "Docker Cloud API Key")
)

func main() {

	flag.Parse()

	p := pegasus.NewPegasus(*brokers)
	processor, err := NewHomepageProcessor(*topic, sarama.OffsetNewest, *dockercloud_user, *dockercloud_apikey)
	if(err != nil) {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	p.StartProcessor(processor)
}
