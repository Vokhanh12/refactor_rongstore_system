package commonv1

// LocaleEnum tương ứng với proto Locale.LocaleEnum
type LocaleEnum int32

const (
	LocaleEnum_UNSPECIFIED LocaleEnum = 0
	LocaleEnum_VI          LocaleEnum = 1
	LocaleEnum_EN          LocaleEnum = 2
)

// Map string để marshal/unmarshal JSON nếu cần
var LocaleEnum_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "VI",
	2: "EN",
}

var LocaleEnum_value = map[string]int32{
	"UNSPECIFIED": 0,
	"VI":          1,
	"EN":          2,
}
