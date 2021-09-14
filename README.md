# kafka-sdk-go

### Kafka Producer Client

```go
topic := []string{"test"}
groupId := "group-2"
address := []string{"127.0.0.1:9092"}
gcfg := types.Config{}
gcfg.Address = address
gcfg.Topics = topic
gcfg.GroupID = groupId
gcfg.Username = ""
gcfg.Password = ""

s.ProducerClient = kafka_sdk_go.NewProducerClient(gcfg)

//produce message

err := s.ProducerClient.Produce("test", []byte("i am a programmer"), 0)

```

### Kafka Consumer Client

```go
topic := []string{"test"}
groupId := "group-2"
address := []string{"127.0.0.1:9092"}
gcfg := types.Config{}
gcfg.Address = address
gcfg.Topics = topic
gcfg.GroupID = groupId
gcfg.Username = ""
gcfg.Password = ""

offsetCfg := types.OffsetConfig{}
offsetCfg.Address = address
offsetCfg.Topic = topic[0]
offsetCfg.GroupID = groupId
offsetCfg.Username = ""
offsetCfg.Password = ""
offsetCfg.Partition = 0
// offsetCfg.ResetOffset = 2 //reset offset 
s.ConsumerClient = kafka_sdk_go.NewConsumerClient(gcfg, offsetCfg)

// consume message
s.ConsumerClient.Start(func(data []byte) error {
if len(data) == 0 {
return nil
}
fmt.Fprintf(os.Stdout, "%s\n",data)
return nil
})
```

