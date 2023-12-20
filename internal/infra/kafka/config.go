package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func ConsumerConfigMap() *ckafka.ConfigMap {
	return &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}
}

func ProducerConfigMap() *ckafka.ConfigMap {
	return &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
	}
}
