package usecase

import (
	"github.com/thanakize/lab_skill_api/repository"
	"github.com/thanakize/lab_skill_api/sharedinterface"
)

type SkillUseCase struct{
	skillRepo repository.ISkillRepository
}

type IUseCase interface{
	GetSkill(string) (sharedinterface.Skill, error)
	GetSkills() ([]sharedinterface.Skill, error)
	InsertSkill(sharedinterface.Skill) (sharedinterface.Skill, error)
	UpdateSkill(sharedinterface.Skill, string) (sharedinterface.Skill, error)
	DeleteSkill(string) error
	PatchSkillName(string, string) (sharedinterface.Skill, error)
	PatchSkillDescription(string, string) (sharedinterface.Skill, error)
	PatchSkillLogo(string, string) (sharedinterface.Skill, error)
	PatchSkillTags([]string, string) (sharedinterface.Skill, error)
}

func InitUseCase (skillRepo repository.ISkillRepository) IUseCase{
	return SkillUseCase{
		skillRepo: skillRepo,
	}
}

func (usecase SkillUseCase) GetSkills() ([]sharedinterface.Skill, error) {
	return usecase.skillRepo.GetSkills()
}
func (usecase SkillUseCase) GetSkill(key string) (sharedinterface.Skill, error) {
	return usecase.skillRepo.GetSkill(key)
}
func (usecase SkillUseCase) InsertSkill(skill sharedinterface.Skill) (sharedinterface.Skill, error) {
	return usecase.skillRepo.InsertSkill(skill)
}
func (usecase SkillUseCase) UpdateSkill(skill sharedinterface.Skill, key string) (sharedinterface.Skill, error) {
	return usecase.skillRepo.UpdateSkill(skill, key)
}
func (usecase SkillUseCase) DeleteSkill(key string) error {
	return usecase.skillRepo.DeleteSkill(key)
}
func (usecase SkillUseCase) PatchSkillName(name string, key string) (sharedinterface.Skill, error) {
	return usecase.skillRepo.PatchSkillName(name, key)
}
func (usecase SkillUseCase) PatchSkillDescription(description string, key string) (sharedinterface.Skill, error) {
	return usecase.skillRepo.PatchSkillDescription(description, key)
}
func (usecase SkillUseCase) PatchSkillLogo(logo string, key string) (sharedinterface.Skill, error) {
	return usecase.skillRepo.PatchSkillLogo(logo, key)
}
func (usecase SkillUseCase) PatchSkillTags(tags []string, key string) (sharedinterface.Skill, error) {
	return usecase.skillRepo.PatchSkillTags(tags, key)
}
