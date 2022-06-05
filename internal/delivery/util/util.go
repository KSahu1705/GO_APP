package util

import (
	// "context"
	"encoding/json"
	// "fmt"
	"net/http"
	// "net/url"
	// "strings"
	// "GO_APP/internal/model/response"
	// "GO_APP/internal/model/entity"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))

	return nil
}

// //WriteResponse write a response from handler
// func WriteResponse(ctx context.Context, handlerName string, w http.ResponseWriter, resp interface{}) {


// 	// w.Header().Set(entity.AllowOriginKey, entity.AllowOriginValue)
// 	// w.Header().Set(entity.ContentTypeKey, entity.ContentTypeValueJSON)

// 	/*resp must be json response*/
// 	byteResult, err := json.Marshal(resp)
// 	if err != nil {
// 		log.ErrorWithFields("[WriteResponse]", log.KV{"error": err, "handler": handlerName})
// 		WriteErrorResponse(ctx, w, http.StatusInternalServerError, "Internal Server Error")
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(byteResult))
// }

// //WriteResponseRedirection write a response from handler
// func WriteResponseRedirection(ctx context.Context, handlerName string, w http.ResponseWriter, resp interface{}) {
// 	span, ctx := tracer.StartSpanFromContext(ctx, "Util.WriteResponseRedirection")
// 	defer span.Finish()

// 	w.Header().Set(entity.AllowOriginKey, entity.AllowOriginValue)
// 	w.Header().Set(entity.ContentTypeKey, entity.ContentTypeValueJSON)

// 	/*resp must be json response*/
// 	byteResult, err := json.Marshal(resp)
// 	if err != nil {
// 		log.ErrorWithFields("[WriteResponse]", log.KV{"error": err, "handler": handlerName})
// 		WriteErrorResponse(ctx, w, http.StatusInternalServerError, "Internal Server Error")
// 	}

// 	http.Error(w, string(byteResult), http.StatusNotFound)
// }

// //WriteErrorResponse write an error response from handler
// func WriteErrorResponse(ctx context.Context, w http.ResponseWriter, statusCode int, errorMessage string) {
// 	span, _ := tracer.StartSpanFromContext(ctx, "Util.WriteErrorResponse")
// 	defer span.Finish()

// 	w.Header().Set(entity.AllowOriginKey, entity.AllowOriginValue)
// 	w.Header().Set(entity.ContentTypeKey, entity.ContentTypeValueJSON)

// 	r := &response.Response{
// 		Header: response.JSONRespHeader{
// 			Message: errorMessage,
// 			Code:    statusCode,
// 		},
// 		Errors: []response.Error{
// 			{
// 				Code:   fmt.Sprintf("HTTP%d", statusCode),
// 				Title:  http.StatusText(statusCode),
// 				Detail: errorMessage,
// 			},
// 		},
// 	}
// 	b, err := json.Marshal(r)
// 	if err != nil {
// 		log.ErrorWithFields("[WriteErrorResponse]", log.KV{"error": err, "message": errorMessage})
// 		statusCode = http.StatusInternalServerError
// 	}

// 	http.Error(w, string(b), statusCode)
// }
