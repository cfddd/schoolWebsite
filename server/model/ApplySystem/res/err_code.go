package res

type ErrorCode int // type ErrorCode int是一种类型声明。这行代码定义了一个新的类型ErrorCode，它是int类型的别名。

// 定义
const (
	SettingsError ErrorCode = 1001 // 系统错误
	ArgumentError ErrorCode = 1002 // 参数错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误",
	}
)
