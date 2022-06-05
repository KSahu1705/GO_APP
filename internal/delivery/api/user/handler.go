package user

import (
	// "GO_APP/internal/delivery/util"
	// "GO_APP/internal/model/entity"
	"context"
	"fmt"
	"net/http"
	// "strconv"
	// "strings"
	// "time"
	// "encoding/json"

	// "GO_APP/internal/model/response"

	"github.com/julienschmidt/httprouter"
	// "GO_APP/internal/repository"
)

// // getUserOr404 gets a User instance if exists, or respond the 404 error otherwise
// func (user *API) getUserOr404(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// // func getUserOr404(db *sqlx.DB, id int, w http.ResponseWriter, r *http.Request) (*model.User, error) {
// 	user := model.User{}
// 	err := db.Get(&user, queries.QueryFindUser, id)
// 	// if (err!=nil){
// 	// 	respondError(w, http.StatusNotFound, err.Error())
// 	// 	return nil
// 	// }
// 	return &user, err
// }

//DetailHandler is a handler for GetCategoryDetail
func (user *API) DetailHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println(params, "@@@@")
	// ctx := r.Context()
	// ctx, cancel := context.WithTimeout(r.Context(), 1*time.Millisecond)
	// defer cancel()
	// decoder := json.NewDecoder(r.Body)

	// var data entity.User
	ctx := context.Background()

	ctx = context.WithValue(ctx, "id", 1)
	// ctx = context.WithValue(ctx, "")

	// _ = decoder.Decode(&data)
	// ctx := 
	// context.WithValue(r.Context(), "id", params)
	// fmt.Println("$$$",data, "\n#####")
	fmt.Println(r.Body)
	fmt.Println("****", ctx,"****")

	// ctx = r.WithContext(ctx)
	fmt.Println("line1111111")
	b,a := user.UserCase.PutUserData(ctx)
	fmt.Println(a,"==>",  b)
}



