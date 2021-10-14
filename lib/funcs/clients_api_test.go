package funcs

import (
	"database/sql/driver"
	"github.com/guark/guark/app"
	"github.com/guark/guark/log"
	"pokos/lib/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestDeleteClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	database.Db = &database.Database{
		Db: db,
	}
	defer func() {
		database.Db = nil
	}()

	collumns := []string{"id"}

	mock.ExpectQuery("SELECT id FROM kkm_registers (.+)").WithArgs([]driver.Value{1}).WillReturnRows(sqlmock.NewRows(collumns).AddRow(driver.Value(1)))

	logger := log.New("test")
	ctx := app.Context{
		App: &app.App{
			Log: logger,
		},
		Params: app.Params{
			"id": 1.0,
		},
	}

	result, err := DeleteClient(ctx)
	if err != nil {
		t.Logf("error is: %s", err)
		t.Fail()
		return
	}
	t.Logf("result: %+v", result)
}
