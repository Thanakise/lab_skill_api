package main

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/thanakize/lab_skill_api/repository"
)

func TestGetAllToDo(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	skillRepo := repository.InitSkill(db)
	// skillRepo.CreateDataTest()
	// defer skillRepo.DeleteDataTest()
	
	
	expectedSQL := "SELECT key, name, description, logo, tags FROM skill"

	skill := sqlmock.NewRows([]string{"key", "name", "description", "logo", "tags"})

	mock.ExpectQuery(expectedSQL).WillReturnRows(skill)
	
	// expectSkill := []sharedinterface.Skill{
	// 	{
	// 		Key : "go",
	// 		Name : "go",
	// 		Description : "Go is an open source programming...",
	// 		Logo : "",
	// 		Tags : []string{"go", "golang"},
	// 	},
	// }
	_, err = skillRepo.GetSkills()

	if err != nil {
		t.Error(err.Error())
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("unmet expectation error: %s", err)
	}
	

}