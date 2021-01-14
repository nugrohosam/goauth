package producers

// RoleTopic ..
const RoleTopic = "topic-role-behavior"

// DeletedRoleKey ..
const DeletedRoleKey = "deleted-role"

// ChangedRoleKey ..
const ChangedRoleKey = "changed-role"

// DeletedRole ..
func DeletedRole(ID string) {
	kafkaConfig.ConfigureKafka(url, RoleTopic)
	kafkaConfig.PushMessageKafka(DeletedRoleKey, ID)
	kafkaConfig.CloseKafkaConn()
}

// ChangedRole ..
func ChangedRole(ID string) {
	kafkaConfig.ConfigureKafka(url, RoleTopic)
	kafkaConfig.PushMessageKafka(ChangedRoleKey, ID)
	kafkaConfig.CloseKafkaConn()
}
