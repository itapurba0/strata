package organization

type Permission string

const (
	PermissionOrganizationRead   Permission = "organization:read"
	PermissionOrganizationUpdate Permission = "organization:update"

	PermissionMemberRead   Permission = "member:read"
	PermissionMemberManage Permission = "member:manage"

	PermissionRoleRead   Permission = "role:read"
	PermissionRoleManage Permission = "role:manage"

	PermissionJobCreate  Permission = "job:create"
	PermissionJobRead    Permission = "job:read"
	PermissionJobUpdate  Permission = "job:update"
	PermissionJobDelete  Permission = "job:delete"
	PermissionJobPublish Permission = "job:publish"

	PermissionApplicationRead   Permission = "application:read"
	PermissionApplicationUpdate Permission = "application:update"

	PermissionCandidateRead Permission = "candidate:read"

	PermissionInterviewRead     Permission = "interview:read"
	PermissionInterviewCreate   Permission = "interview:create"
	PermissionInterviewEvaluate Permission = "interview:evaluate"
)

type PermissionDefinition struct {
	Name        Permission
	Description string
}

var systemPermissions = []PermissionDefinition{
	{
		Name:        PermissionOrganizationRead,
		Description: "View organization information",
	},
	{
		Name:        PermissionOrganizationUpdate,
		Description: "Update organization information",
	},
	{
		Name:        PermissionMemberRead,
		Description: "View organization members",
	},
	{
		Name:        PermissionMemberManage,
		Description: "Manage organization members",
	},
	// ...
}
