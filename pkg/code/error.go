package code

var errMap = map[string]int{
	"RECORD_OK": 0,
	//数据不存在
	"RECORD_NODATA": 4002,
	//参数错误
	"RECORD_PARAMERR": 4003,
	//服务内部错误
	"RECORD_SERVERERR": 4500,
}

// GetCode 获取错误码
func GetCode(code string) uint32 {
	if code, err := errMap[code]; !err {
		return uint32(code)
	}
	return 0
}
