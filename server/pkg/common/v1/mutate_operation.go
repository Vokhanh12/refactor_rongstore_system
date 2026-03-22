package commonv1

// MutateOperationEnum tương ứng với proto MutateOperation.MutateOperationEnum
type MutateOperationEnum int32

const (
	MutateOperationEnum_UNSPECIFIED MutateOperationEnum = 0
	MutateOperationEnum_CREATE      MutateOperationEnum = 1
	MutateOperationEnum_EDIT        MutateOperationEnum = 2
	MutateOperationEnum_DELETE      MutateOperationEnum = 3
)

// Map string ↔ int32 để marshal/unmarshal JSON nếu cần
var MutateOperationEnum_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "CREATE",
	2: "EDIT",
	3: "DELETE",
}

var MutateOperationEnum_value = map[string]int32{
	"UNSPECIFIED": 0,
	"CREATE":      1,
	"EDIT":        2,
	"DELETE":      3,
}

// Optional helper function
func (op MutateOperationEnum) String() string {
	if s, ok := MutateOperationEnum_name[int32(op)]; ok {
		return s
	}
	return "UNKNOWN"
}
