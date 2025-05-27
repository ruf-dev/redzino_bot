package permissions

type Permission int64

const (
	PermissionAddVideo Permission = 1 << 0 // 0001
	//PermissionWrite    = 1 << 1 // 0010
	//PermissionDelete   = 1 << 2 // 0100
	//PermissionAdmin    = 1 << 3 // 1000
)
