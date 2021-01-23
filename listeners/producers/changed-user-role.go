package producers

import "fmt"

// UserRoleTopic ..
const UserRoleTopic = "topic-user-role-behavior"

// DeletedUserRoleKey ..
const DeletedUserRoleKey = "deleted-user-role"

// ChangedUserRoleKey ..
const ChangedUserRoleKey = "changed-user-role"

// DeletedUserRole ..
func DeletedUserRole(ID int) {
	kafkaConfig.ConfigureKafka(url, UserRoleTopic)
	kafkaConfig.PushMessageKafka(DeletedUserRoleKey, fmt.Sprint(ID))
	kafkaConfig.CloseKafkaConn()
}

// ChangedUserRole ..
func ChangedUserRole(ID int) {
	kafkaConfig.ConfigureKafka(url, UserRoleTopic)
	kafkaConfig.PushMessageKafka(ChangedUserRoleKey, fmt.Sprint(ID))
	kafkaConfig.CloseKafkaConn()
}
