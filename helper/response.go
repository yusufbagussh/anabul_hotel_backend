package helper

import (
	"strings"
)

// Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type DataRes struct {
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"paginate"`
}

// EmptyObj object is used when data doesn't want to be null on json
type EmptyObj struct{}

// BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildResponsePage(status bool, message string, data interface{}, pagination interface{}) Response {
	type DataPagination struct {
	}
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data: DataRes{
			Data:       data,
			Pagination: pagination,
		},
	}
	return res
}

// BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

//func newUUID() (string, error) {
//	uuid := make([]byte, 16)
//	n, err := io.ReadFull(rand.Reader, uuid)
//	if n != len(uuid) || err != nil {
//		return "", err
//	}
//	// variant bits; see section 4.1.1
//	uuid[8] = uuid[8]&^0xc0 | 0x80
//	// version 4 (pseudo-random); see section 4.1.3
//	uuid[6] = uuid[6]&^0xf0 | 0x40
//	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
//}
