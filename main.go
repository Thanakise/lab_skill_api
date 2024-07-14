package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/thanakize/lab_skill_api/repository"
	"github.com/thanakize/lab_skill_api/server"
	"github.com/thanakize/lab_skill_api/usecase"
)


func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	
	if err != nil {
		log.Fatal("Connect to database error", err)
	}


	skillRepo := repository.InitSkill(db)
	skillUsecase := usecase.InitUseCase(skillRepo)
	defer skillRepo.CloseDB()

	// server.InitServerGO(skillUsecase)
	goServer := server.InitServerGO(skillUsecase)
	goServer.StartServer()
}
