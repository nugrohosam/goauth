package producers

import "fmt"

// RolePermissionTopic ..
const RolePermissionTopic = "topic-role-permission-behavior"

// DeletedRolePermissionKey ..
const DeletedRolePermissionKey = "deleted-role-permission"

// ChangedRolePermissionKey ..
const ChangedRolePermissionKey = "changed-role-permission"

// DeletedRolePermission ..
func DeletedRolePermission(ID int) {
	kafkaConfig.ConfigureKafka(url, RolePermissionTopic)
	kafkaConfig.PushMessageKafka(DeletedRolePermissionKey, fmt.Sprint(ID))
	kafkaConfig.CloseKafkaConn()
}

// ChangedRolePermission ..
func ChangedRolePermission(ID int) {
	kafkaConfig.ConfigureKafka(url, RolePermissionTopic)
	kafkaConfig.PushMessageKafka(ChangedRolePermissionKey, fmt.Sprint(ID))
	kafkaConfig.CloseKafkaConn()

}
