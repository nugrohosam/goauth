package producers

// UserTopic ..
const UserTopic = "topic-user-behavior"

// DeletedUserKey ..
const DeletedUserKey = "deleted-user"

// ChangedUserKey ..
const ChangedUserKey = "changed-user"

// DeletedUser ..
func DeletedUser(ID string) {
	kafkaConfig.ConfigureKafka(url, UserTopic)
	kafkaConfig.PushMessageKafka(DeletedUserKey, ID)
	kafkaConfig.CloseKafkaConn()
}

// ChangedUser ..
func ChangedUser(ID string) {
	kafkaConfig.ConfigureKafka(url, UserTopic)
	kafkaConfig.PushMessageKafka(ChangedUserKey, ID)
	kafkaConfig.CloseKafkaConn()
}
