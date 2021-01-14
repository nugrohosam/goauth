package producers

// UserRoleTopic ..
const UserRoleTopic = "topic-user-role-behavior"

// DeletedUserRoleKey ..
const DeletedUserRoleKey = "deleted-user-role"

// ChangedUserRoleKey ..
const ChangedUserRoleKey = "changed-user-role"

// DeletedUserRole ..
func DeletedUserRole(ID string) {
	kafkaConfig.ConfigureKafka(url, UserRoleTopic)
	kafkaConfig.PushMessageKafka(DeletedUserRoleKey, ID)
	kafkaConfig.CloseKafkaConn()
}

// ChangedUserRole ..
func ChangedUserRole(ID string) {
	kafkaConfig.ConfigureKafka(url, UserRoleTopic)
	kafkaConfig.PushMessageKafka(ChangedUserRoleKey, ID)
	kafkaConfig.CloseKafkaConn()
}
