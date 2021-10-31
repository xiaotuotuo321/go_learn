package main

// 1.初识
/*
sarama用于处理kafka的纯Go客户端，它包括一个易于生成和使用消息的高级API，以及一个在高级API不足时控制线路上字节的低级API。
内嵌提供了高级API的用法,示例及其完整文档
*/

// 2.生产者：有几种类型的生产者
// sarama.NewSyncProducer() 同步发送者
// sarama.NewAsyncProducer() 异步发送者

// 2.1.同步模式
//func main() {
//	config := sarama.NewConfig() // 实例化sarama的config
//	config.Producer.Return.Successes = true // 是否开启消息发送成功后通知 successes channel1
//	config.Producer.Partitioner = sarama.NewRandomPartitioner	// 随机分机器
//	client, err := sarama.NewClient([]string{"127.0.0.1:9092"}, config)	// 初始化客户端
//	defer client.Close()
//
//	if err != nil {panic(err)}
//
//	producer, err := sarama.NewSyncProducerFromClient(client)
//
//	if err != nil {panic(err)}
//
//	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
//		Topic: "liangtian_topic",
//		Key: nil,
//		Value: sarama.StringEncoder("hahaha"),
//	})
//
//	if err != nil {
//		log.Fatal("unable to produce message: %q", err)
//	}
//	fmt.Println("partition", partition)
//	fmt.Println("offset", offset)
//}

// 2.2.异步模式：produce一个message之后不等待发送完成返回；真阳调用者可以继续做其他的工作
//func main() {
//	config := sarama.NewConfig()
//	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
//
//	if err != nil{
//		log.Fatal("unable to create kafka client: %q", err)
//	}
//
//	producer, err := sarama.NewAsyncProducerFromClient(client)
//
//	if err != nil{
//		log.Fatal("unable to create kafka producer: %q", err)
//	}
//
//	defer producer.Close()
//	text := fmt.Sprintf("message: %08d", 1)
//
//	producer.Input() <- &sarama.ProducerMessage{Topic:"test001", Key: nil, Value: sarama.StringEncoder(text)}
//
//	select {
//	case err := <- producer.Errors():
//		log.Println("Produced message failure ", err)
//	default:
//		log.Println("Produced message default")
//	}
//}

// 2.3.关于异步producer有一个地方是需要注意的，异步模式produce一个消息后，缺省并不会报告成功状态
// 结论就是说配置：config.Producer.Return.Success = true和操作 <- producer.success必须是配套使用的。配置成true，那么就去读取
// Success,如果配置成false则不去读取successes。

// 3.消费者
// 3.1.使用消费组消费：每一个topic只能被一个消费者所消费。一个消费者可以同时消费对个topic

