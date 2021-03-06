package code

import (
	"encoding/json"
)

type ServiceError struct {
	Id     string
	Code   int
	Detail string
	Status string
}

var errMap = map[string]uint32{
	"RECORD_OK": 0,
	//数据不存在
	"RECORD_NODATA": 4002,
	//参数错误
	"RECORD_PARAMERR": 4003,
	//服务内部错误
	"RECORD_SERVERERR": 4500,
	//微服务报错
	"RECORD_MICROSERVERERERR": 4501,
}

// GetCode 获取错误码
func GetCode(code string) uint32 {
	if status, ok := errMap[code]; ok {
		return status
	}
	return 0
}

//处理service 错误信息
func HandleServiceError(err error) ServiceError {
	var temp ServiceError
	jsonErr := json.Unmarshal([]byte(err.Error()), &temp)
	if jsonErr != nil {
		return temp
	}
	return temp
}
