package sqlite3

import (
	"database/sql"
	"log"
	"reflect"
	"testing"

	"forum/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

func TestSqlite_CheckTheCookie(t *testing.T) {
	dbPath := "database.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	type fields struct {
		db      *sql.DB
		configs interface{}
	}
	type args struct {
		cookie   string
		livetime int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name:    "hui",
			fields:  fields{db: db, configs: ""},
			args:    args{cookie: "123", livetime: 1},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sqlite{
				db:      tt.fields.db,
				configs: tt.fields.configs,
			}
			got, err := s.CheckTheCookie(tt.args.cookie, tt.args.livetime)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sqlite.CheckTheCookie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sqlite.CheckTheCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
