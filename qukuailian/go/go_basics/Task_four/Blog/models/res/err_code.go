package res

const (
	SettingsError int = 1001 //系统错误
	ArgumentError int = 1002 //参数错误
)

var (
	ErrorMap = map[int]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误",
	}
)
