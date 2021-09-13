package types

type ServerConf struct {
	Bech32AccPrefix  string
	KafkaNodeurls    []string
	KafkaTopic       string
	KafkaGroupId     string
	KafkaUsername    string
	KafkaPassword    string
	KafkaResetOffset int64
}

type Config struct {
	Address  []string
	Topics   []string
	GroupID  string
	Username string
	Password string
	Offset   int64
}

type OffsetConfig struct {
	Address     []string
	Username    string
	Password    string
	Topic       string
	GroupID     string
	Partition   int32
	ResetOffset int64
}

//
//var (
//	SvrConf *ServerConf
//
//	bech32AccPrefix  = "iaa"
//	kafkaTopic       = "BLOCKS_GLOBAL"
//	kafkaGroupID     = "explorer-sync-consumer"
//	kafkaUsername    = "kafkaUser"
//	kafkaPassword    = "kafkaPassword"
//	kafkaResetOffset = ""
//	kafkaNodeurls    = []string{"10.0.0.49:9094", "10.0.0.21:9094", "10.0.0.8:9094"}
//)
//
//// get value of env var
//func init() {
//	if v, ok := os.LookupEnv(constant.EnvNameBech32AccPrefix); ok {
//		bech32AccPrefix = v
//	}
//	if v, ok := os.LookupEnv(constant.EnvNameKafkaGroupID); ok {
//		kafkaGroupID = v
//	}
//	if v, ok := os.LookupEnv(constant.EnvNameKafkaUserName); ok {
//		kafkaUsername = v
//	}
//	if v, ok := os.LookupEnv(constant.EnvNameKafkaPassword); ok {
//		kafkaPassword = v
//	}
//	if v, ok := os.LookupEnv(constant.EnvNameKafkaTopic); ok {
//		kafkaTopic = v
//	}
//	if v, ok := os.LookupEnv(constant.EnvNameKafkaNodeUrls); ok {
//		kafkaNodeurls = strings.Split(v, ",")
//	}
//	if v, ok := os.LookupEnv(constant.EnvNameKafkaResetOffset); ok {
//		kafkaResetOffset = v
//	}
//
//	SvrConf = &ServerConf{
//		Bech32AccPrefix:  bech32AccPrefix,
//		KafkaTopic:       kafkaTopic,
//		KafkaGroupId:     kafkaGroupID,
//		KafkaUsername:    kafkaUsername,
//		KafkaPassword:    kafkaPassword,
//		KafkaNodeurls:    kafkaNodeurls,
//		KafkaResetOffset: kafkaResetOffset,
//	}
//
//	logger.Debug("print server config", logger.String("serverConf", utils.MarshalJsonIgnoreErr(SvrConf)))
//}
