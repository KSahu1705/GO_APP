
package response

// import (
// 	"GO_APP/internal/models/entity"
// )

// JSONRespHeader is a JSON structure for user response header
type JSONRespHeader struct {
	Message   string `json:"message,omitempty"`
	Reason    string `json:"reason,omitempty"`
	ErrorCode int64  `json:"error_code,omitempty"`
	Code      int    `json:"code,omitempty"`
}

// JSONResponse is a JSON structure for user response
type JSONResponse struct {
	Header      JSONRespHeader `json:"header"`
	Data        interface{}    `json:"data"`
	StatusCode  int            `json:"-"`
	ErrorString string         `json:"error,omitempty"`
	Log         string         `json:"-"`
}