package producers

// RolePermissionTopic ..
const RolePermissionTopic = "topic-role-permission-behavior"

// DeletedRolePermissionKey ..
const DeletedRolePermissionKey = "deleted-role-permission"

// ChangedRolePermissionKey ..
const ChangedRolePermissionKey = "changed-role-permission"

// DeletedRolePermission ..
func DeletedRolePermission(ID string) {
	kafkaConfig.ConfigureKafka(url, RolePermissionTopic)
	kafkaConfig.PushMessageKafka(DeletedRolePermissionKey, ID)
	kafkaConfig.CloseKafkaConn()
}

// ChangedRolePermission ..
func ChangedRolePermission(ID string) {
	kafkaConfig.ConfigureKafka(url, RolePermissionTopic)
	kafkaConfig.PushMessageKafka(ChangedRolePermissionKey, ID)
	kafkaConfig.CloseKafkaConn()

}
