package producers

import "fmt"

// UserTopic ..
const UserTopic = "topic-user-behavior"

// DeletedUserKey ..
const DeletedUserKey = "deleted-user"

// ChangedUserKey ..
const ChangedUserKey = "changed-user"

// DeletedUser ..
func DeletedUser(ID int) {
	kafkaConfig.ConfigureKafka(url, UserTopic)
	kafkaConfig.PushMessageKafka(DeletedUserKey, fmt.Sprint(ID))
	kafkaConfig.CloseKafkaConn()
}

// ChangedUser ..
func ChangedUser(ID int) {
	kafkaConfig.ConfigureKafka(url, UserTopic)
	kafkaConfig.PushMessageKafka(ChangedUserKey, fmt.Sprint(ID))
	kafkaConfig.CloseKafkaConn()
}
