package main

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/thanakize/lab_skill_api/sharedinterface"
	"github.com/thanakize/lab_skill_api/views"
)

func TestGetAllToDo(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`
		CREATE TABLE IF NOT EXISTS todos (
				id SERIAL PRIMARY KEY,
				title TEXT,
				status TEXT
		);
		INSERT INTO
			todos (title, status)
		VALUES
			(
				'write code',
				'active',
			)

	`).WithArgs(nil)
	mock.ExpectQuery("SELECT id, title, status FROM todos").WillReturnRows(sqlmock.NewRows([]string{}))


	q := `
		CREATE TABLE IF NOT EXISTS todos (
				id SERIAL PRIMARY KEY,
				title TEXT,
				status TEXT
		);
		INSERT INTO
			todos (title, status)
		VALUES
			(
				'write code',
				'active',
			)

	`
	if _, err := db.Exec(q) ; err != nil{
		t.Error(err.Error())
	}



	expectTodo := []sharedinterface.Todo{
		{
			ID: 1,
			Title: "write code",
			Status: "active",
		},
	}
	todos, err := views.GetTodos(db)

	if err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(todos, expectTodo) {
		t.Errorf("need: %v got : %v",expectTodo, todos)
	}

}