package postgres

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func Test_queries_prepareStatement(t *testing.T) {
	first := "SELECT 1"
	mockdb, m, err := sqlmock.New()
	if err != nil {
		panic(err) 
	}
	defer mockdb.Close()
	sqlxDB := sqlx.NewDb(mockdb, "sqlmock")

	type args struct {
		db      *sqlx.DB
		queries string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "tc_fa_success",
			args: args{db: sqlxDB, queries: first},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.ExpectPrepare(tt.args.queries)

			q := &queries{}
			got, err := q.prepareStatement(tt.args.db, tt.args.queries)
			fmt.Println(got, err)
			if err != nil {
				t.Errorf("[queries][prepareStatement] error test case: %v", err)
			}
			if got == nil {
				t.Errorf("[queries][prepareStatement] error test case: %d, function should not return nil", i)
			}
		})
	}
}

// go test queries_test.go queries.go const.go

// ok      command-line-arguments  1.519s


