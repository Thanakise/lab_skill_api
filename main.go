package main

import (
	"github.com/thanakize/lab_skill_api/repository"
	"github.com/thanakize/lab_skill_api/server"
	"github.com/thanakize/lab_skill_api/usecase"
)


func main() {
	skillRepo := repository.InitSkill()
	skillUsecase := usecase.InitUseCase(skillRepo)
	defer skillRepo.CloseDB()

	// server.InitServerGO(skillUsecase)
	goServer := server.InitServerGO(skillUsecase)
	goServer.StartServer()
}
