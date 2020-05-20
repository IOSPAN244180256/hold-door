package utils

var _kafkaLog *kafkaLog = &kafkaLog{producer: nil}

type kafkaLog struct {
	producer *sarama.SyncProducer
}

func (kl *kafkaLog)Write(p []byte) (n int, err error){
	msg := &sarama.ProducerMessage{}
	msg.Topic = lk.Topic
	msg.Value = sarama.ByteEncoder(p)
	_, _, err = kl.Producer.SendMessage(msg)
	if err != nil {
		return
	}
	return
}

func NewKafkaLog() (kl *kafkaLog,err error){
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

		_kafkaLog.producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
		return  _kafkaLog,err

	}else{
		return  _kafkaLog,nil
	}
}