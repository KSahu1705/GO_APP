package postgres
const (
	createDB = `
		CREATE TABLE IF NOT EXISTS users (
			id        int primary key,
			created_at timestamp,
			updated_at timestamp,
			deleted_at timestamp,
			username  text,
			password  text,
			first_name text,
			last_name  text,
			phone     text,
			status bool
		);

		CREATE TABLE IF NOT EXISTS user_addresses (
			id        int primary key,	   
			created_at timestamp,
			updated_at timestamp,
			deleted_at timestamp, 
			user_id       int,	
			address_line1 text, 
			address_line2 text,	
			city         text,	
			postal_code   text,	
			country      text,	
			phone        text,
			telephone    text,
			FOREIGN KEY (user_id) REFERENCES users(id)				
		);
	`
	// queryInsertUserData = `
	// 	INSERT INTO users (id, created_at, updated_at, deleted_at, username, password, first_name, last_name, phone, status) VALUES(:id, :created_at, :updated_at, :deleted_at, :username, :password, :first_name, :last_name, :phone, :status)
	// `

	queryInsertUserData = `
		INSERT INTO users (id, created_at, updated_at, deleted_at, username, password, first_name, last_name, phone, status) VALUES($1, now(), now(), $2, $3, $4, $5, $6, $7, $8)
	`

	queryInsertAddrData = `
		INSERT INTO user_addresses (id, created_at, updated_at, deleted_at, user_id, address_line1, address_line2, city, postal_code, country, phone, telephone) VALUES($1, now(), now(), $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	queryFindUser = `SELECT * FROM users WHERE id=$1`

	queryFindUserAddr = `SELECT * FROM user_addresses WHERE user_id=$1`

	queryAlluser = `SELECT * FROM users`

	queryFilterUserAddress = `SELECT * FROM user_addresses WHERE user_id=$1`
	
	queryUpdateUser = `UPDATE users SET updated_at=:updated_at, username=:username, password=:password, first_name=:first_name, last_name=:last_name, phone=:phone, status=:status WHERE id=:id;`

	queryFilterUserAddressWid = `SELECT * FROM user_addresses WHERE id=$2 AND user_id=$1;`

	queryUpdateUserAddr = `UPDATE user_addresses SET updated_at=:updated_at, address_line1=:address_line1, address_line2=:address_line2, city=:city, postal_code=:postal_code, country=:country, phone=:phone, telephone=:telephone WHERE id=:id AND user_id=:user_id;`

	queryUpdateUserStatus = `UPDATE users SET status=:status WHERE id=:id;`

	queryDeleteUser = `DELETE FROM users WHERE id=$1;`

	queryDeleteUserAddr = `DELETE FROM user_addresses WHERE id=:id AND user_id=:user_id;`

	queryDeleteAddresses = `DELETE FROM user_addresses WHERE user_id=$1;`

)
