package postgres

import (
	"errors"
	"fmt"
	// "sync"
	// "time"
	// "log"

	"github.com/jmoiron/sqlx"
	"GO_APP/internal/model/config"
	// "GO_APP/internal/model/entity"
)

//Postgres is a struct
type Postgres struct {
	driverName                             string
	cfg                                    map[string]*config.DBConfig
	dbMap                                  map[string]*sqlx.DB
	queries                                *queries
}

//NewPostgres is a method
func NewPostgres(config map[string]*config.DBConfig) (*Postgres, error) {
	var err error
	postgres := &Postgres{
		driverName: "postgres",
		cfg:        config,
		dbMap:      make(map[string]*sqlx.DB),
	}

	db, err := postgres.GetConnection("go_dummy")
	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, errors.New("failed connecting to category master DB")
	}

	if (db!=nil){
		postgres.queries, err = newQueries(db)
		if err != nil {
			return nil, err
		}
	}

	return postgres, nil
}

//GetConnection is a method
func (postgres *Postgres) GetConnection(connection string) (*sqlx.DB, error) {
	var (
		db  *sqlx.DB
		err error
	)

	// check for invalid db string (subconfig DB not present)
	if _, ok := postgres.cfg[connection]; !ok {
		return nil, errors.New("[postgres][GetConnection] invalid database name specified")
	}

	c := postgres.cfg[connection]

	db = postgres.dbMap[connection]
	if db != nil {
		err = db.Ping()
		if err == nil {
			return db, err
		}
	}
	

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",c.Host, c.Port, c.User, c.Password, c.DBname)
	// fmt.Println(conn)

	db, err = sqlx.Connect(postgres.driverName, conn)
	// fmt.Println(db)
	if err != nil {
		err = fmt.Errorf("[%s] %s", connection, err.Error())
		return db, err
	}

	// if err != nil {
	// 	log.Fatal("Could not connect database")
	// } else {
	// 	fmt.Printf("Connected to database\n")
	// }
	return db, err
}