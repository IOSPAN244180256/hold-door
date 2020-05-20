package utils

import (
	"github.com/Shopify/sarama"
)

var _kafkaLog *kafkaLog = &kafkaLog{producer: nil}

type kafkaLog struct {
	producer sarama.SyncProducer
	topic    string
}

func (kl *kafkaLog) Write(p []byte) (n int, err error) {
	msg := &sarama.ProducerMessage{
		Topic: kl.topic,
		Value: sarama.ByteEncoder(p),
	}
	_, _, err = (kl.producer).SendMessage(msg)
	if err != nil {
		return
	}
	return
}

func NewKafkaLog(topic string) (kl *kafkaLog, err error) {
	_kafkaLog.topic = topic

	if _kafkaLog.producer == nil {

		// 设置日志输入到Kafka的配置
		config := sarama.NewConfig()
		//等待服务器所有副本都保存成功后的响应
		config.Producer.RequiredAcks = sarama.WaitForAll
		//随机的分区类型
		config.Producer.Partitioner = sarama.NewRandomPartitioner
		//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
		config.Producer.Return.Successes = true
		config.Producer.Return.Errors = true

		producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
		if err != nil {
			_kafkaLog.producer = producer
		}
		return _kafkaLog, err

	} else {
		return _kafkaLog, nil
	}
}
