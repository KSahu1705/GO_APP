package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type queries struct {
	insertUserData, 
	insertAddrData, 
	getUserByID, 
	getUserAddressByID, 
	getUserAddressByIDAddrID, 
	getAllUser, 
	updateUserByID, 
	updateUserAddress, 
	updateUserStatus, 
	deleteUser, 
	deleteUserAddr, 
	deleteAddresses *sqlx.Stmt
}

//newQueries is a method for instantiation
func newQueries(db *sqlx.DB) (*queries, error) {
	
	queries := &queries{}

	constructUserData, err := queries.prepareStatement(db, queryInsertUserData)
	fmt.Println("err1",err)

	if err != nil {
		return nil, err
	}

	constructAddrData, err := queries.prepareStatement(db, queryInsertAddrData)
	fmt.Println("err2",err)

	if err != nil {
		return nil, err
	}


	// userByID, err := queries.prepareStatement(db, queryFindUser)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("err3",err)

	// userAddressByID, err := queries.prepareStatement(db, queryFindUserAddr)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("err4",err)

	// userAddressByIDAddrID, err := queries.prepareStatement(db, queryFilterUserAddressWid)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("err5",err)


	// getAllUser, err := queries.prepareStatement(db, queryAlluser)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("err6",err)


	// updateUserByID, err := queries.prepareStatement(db, queryUpdateUser)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("err7",err)


	// updateUserAddress, err := queries.prepareStatement(db, queryUpdateUserAddr)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("err8",err)

	// updateUserStatus, err := queries.prepareStatement(db, queryUpdateUserStatus)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("err9",err)


	// deleteUser, err := queries.prepareStatement(db, queryDeleteUser)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("err10",err)


	// deleteUserAddr, err := queries.prepareStatement(db, queryDeleteUserAddr)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("err11", err)


	// deleteAddresses, err := queries.prepareStatement(db, queryDeleteAddresses)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("err12", err)


	queries.insertUserData = constructUserData
	queries.insertAddrData = constructAddrData
	// queries.getUserByID = userByID
	// queries.getUserAddressByID = userAddressByID
	// queries.getUserAddressByIDAddrID = userAddressByIDAddrID
	// queries.getAllUser = getAllUser
	// queries.updateUserByID = updateUserByID
	// queries.updateUserAddress = updateUserAddress
	// queries.updateUserStatus = updateUserStatus
	// queries.deleteUser = deleteUser
	// queries.deleteUserAddr = deleteUserAddr
	// queries.deleteAddresses = deleteAddresses

	return queries, err
}

//prepareStatement is a method for preparing sqlx statement
func (queries *queries) prepareStatement(db *sqlx.DB, query string) (*sqlx.Stmt, error) {
	stmt, err := db.Preparex(query) //https://go.dev/doc/database/prepared-statements
	if err != nil {
		return nil, err
	}

	return stmt, err
}