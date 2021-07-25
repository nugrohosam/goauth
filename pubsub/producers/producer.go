package producers

import (
	"github.com/nugrohosam/gosampleapi/services/infrastructure"
	"github.com/spf13/viper"
)

var kafkaServerHost = viper.GetString("kafka.server-1.host")
var kafkaServerPort = viper.GetString("kafka.server-1.port")

var kafkaConfig infrastructure.KafkaConfig

var url = kafkaServerHost + ":" + kafkaServerPort
