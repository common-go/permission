package permission

const (
	ActionNone    int32 = 0
	ActionRead    int32 = 1
	ActionAdd     int32 = 2
	ActionEdit    int32 = 4
	ActionPatch   int32 = 8
	ActionDelete  int32 = 16
	ActionApprove int32 = 32
	ActionAll     int32 = 2147483647
)
