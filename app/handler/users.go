package handler

import (
	"GO_APP/app/model"
	"GO_APP/app/queries"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// getUserOr404 gets a User instance if exists, or respond the 404 error otherwise
func getUserOr404(db *sqlx.DB, id int, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	err := db.Get(&user, queries.QueryFindUser, id)
	if (err!=nil){
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &user
}

// getUserAddressOr404 gets a User Address instance if exists, or respond the 404 error otherwise
func getUserAddressOr404(db *sqlx.DB, id int, w http.ResponseWriter, r *http.Request) []model.UserAddress {
	userAddress := []model.UserAddress{}
	err :=  db.Select(&userAddress, queries.QueryFindUserAddr, id)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return userAddress
}

func CreateUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	curtime := time.Now()
	user.CreatedAt = curtime
	user.UpdatedAt = curtime
	tx := db.MustBegin()
	tx.NamedExec(queries.QueryInsertUserData, &user)	
    err:=tx.Commit()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, user)
}

func CreateUserAddress(db *sqlx.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return 
	}

	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	userAddress := model.UserAddress{UserId: id}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userAddress); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	curtime := time.Now()
	userAddress.CreatedAt = curtime
	userAddress.UpdatedAt = curtime
	tx := db.MustBegin()
	tx.NamedExec(queries.QueryInsertAddrData, &userAddress)	
    err=tx.Commit()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, userAddress)
}

func GetAllUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	err := db.Select(&users, queries.QueryAlluser)
	fmt.Println(users)
	if (err!=nil){
		respondError(w, http.StatusNotFound, err.Error())
	}

	for i, ar := range users {

		addresses := []model.UserAddress{}
		err = db.Select(&addresses, queries.QueryFilterUserAddress, ar.ID)
		users[i].Addrs = addresses
	}

	respondJSON(w, http.StatusOK, users)
}

func GetUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	addresses := []model.UserAddress{}
	err = db.Select(&addresses, queries.QueryFilterUserAddress, id)
	user.Addrs = addresses

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	respondJSON(w, http.StatusOK, user)
}

func GetUserAddress(db *sqlx.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	userAddress := getUserAddressOr404(db, id, w, r)
	if userAddress == nil {
		return
	}

	respondJSON(w, http.StatusOK, userAddress)
}

func UpdateUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}

	user.UpdatedAt = time.Now()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	tx := db.MustBegin()
	tx.NamedExec(queries.QueryUpdateUser, &user)	
	err = tx.Commit()

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func UpdateUserAddress(db *sqlx.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err1 := strconv.Atoi(ps.ByName("id"))
	addr_id, err2 := strconv.Atoi(ps.ByName("addr_id"))

	if err1 != nil {
		respondError(w, http.StatusInternalServerError, err1.Error())
	}
	if err2 != nil {
		respondError(w, http.StatusInternalServerError, err2.Error())
	}
	userAddress := model.UserAddress{}
	tx := db.MustBegin()
	tx.Get(&userAddress, queries.QueryFilterUserAddressWid, id, addr_id)
	err := tx.Commit()
	userAddress.UpdatedAt = time.Now()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userAddress); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	tx = db.MustBegin()
	tx.NamedExec(queries.QueryUpdateUserAddr, &userAddress)
	err = tx.Commit()

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, userAddress)
}

func DisableUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	user.Disable()

	tx := db.MustBegin()
	tx.NamedExec(queries.QueryUpdateUserStatus, &user)
	err = tx.Commit()

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func EnableUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	user.Enable()

	tx := db.MustBegin()
	tx.NamedExec(queries.QueryUpdateUserStatus, &user)
	err = tx.Commit()

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, user)
}


func DeleteUser(db *sqlx.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}

	user.DeletedAt = time.Now()
	tx := db.MustBegin()

	tx.MustExec(queries.QueryDeleteAddresses, id)
    tx.MustExec(queries.QueryDeleteUser, id)	
	err = tx.Commit()

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// respondJSON(w, http.StatusNoContent, nil)
}

func DeleteUserAddress(db *sqlx.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err1 := strconv.Atoi(ps.ByName("id"))
	addr_id, err2 := strconv.Atoi(ps.ByName("addr_id"))
	if err1 != nil {
		respondError(w, http.StatusInternalServerError, err1.Error())
	}
	if err2 != nil {
		respondError(w, http.StatusInternalServerError, err2.Error())
	}

	userAddress := model.UserAddress{}

	tx := db.MustBegin()
	tx.Get(&userAddress, queries.QueryFilterUserAddressWid, id, addr_id)
	err := tx.Commit()
	userAddress.DeletedAt = time.Now()

	tx = db.MustBegin()
	// tx.NamedExec("UPDATE user_addresses SET deleted_at=:deleted_at WHERE id=:id AND user_id=:user_id;", &userAddress)
	tx.NamedExec(queries.QueryDeleteUserAddr, &userAddress)

	err = tx.Commit()

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// respondJSON(w, http.StatusOK, userAddress)
}

// {
// 	"Id":  1,
// 	"UserId": 240,
// 	"AddressLine1": "abcdf",
// 	"AddressLine2": "lucknow",
// 	"City": "kanpur",
// 	"PostalCode": "226021",
// 	"Country": "India",
// 	"Phone": "9818476950",
// 	"Telephone": "783232"
// }

// {
//     "Id":2140,
//     "Username":"Kriti",
//     "Password":"",
//     "ModifiedAt":"",
//     "Addrs":null
// }
