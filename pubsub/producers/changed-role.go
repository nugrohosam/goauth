package producers

import "fmt"

// RoleTopic ..
const RoleTopic = "topic-role-behavior"

// DeletedRoleKey ..
const DeletedRoleKey = "deleted-role"

// ChangedRoleKey ..
const ChangedRoleKey = "changed-role"

// DeletedRole ..
func DeletedRole(ID int) {
	kafkaConfig.ConfigureKafka(url, RoleTopic)
	kafkaConfig.PushMessageKafka(DeletedRoleKey, fmt.Sprint(ID))
	kafkaConfig.CloseKafkaConn()
}

// ChangedRole ..
func ChangedRole(ID int) {
	kafkaConfig.ConfigureKafka(url, RoleTopic)
	kafkaConfig.PushMessageKafka(ChangedRoleKey, fmt.Sprint(ID))
	kafkaConfig.CloseKafkaConn()
}
