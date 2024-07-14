package repository

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/thanakize/lab_skill_api/sharedinterface"
)

type SkillReposiory struct{
	DB *sql.DB
}

type ISkillRepository interface{
	CloseDB()
	GetSkill(key string) (sharedinterface.Skill, error)
	GetSkills() ([]sharedinterface.Skill, error)
	InsertSkill(skill sharedinterface.Skill) (sharedinterface.Skill, error)
	UpdateSkill(skill sharedinterface.Skill, key string) (sharedinterface.Skill, error)
	DeleteSkill(key string) error
	PatchSkillName(name string, key string) (sharedinterface.Skill, error)
	PatchSkillDescription(description string, key string) (sharedinterface.Skill, error)
	PatchSkillLogo(logo string, key string) (sharedinterface.Skill, error)
	PatchSkillTags(tags []string, key string) (sharedinterface.Skill, error)
}

func InitSkill(db *sql.DB) ISkillRepository{
	return SkillReposiory{
		DB: db,
	}
}

func (repo SkillReposiory) CloseDB(){
	repo.DB.Close()
}


func (repo SkillReposiory) GetSkills() ([]sharedinterface.Skill, error) {
	var skills []sharedinterface.Skill
	rows, err := repo.DB.Query("SELECT key, name, description, logo, tags FROM skill")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var skill sharedinterface.Skill
		err := rows.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
		if err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}
	return skills, nil
}
func (repo SkillReposiory) GetSkill(key string) (sharedinterface.Skill, error) {
	var skill sharedinterface.Skill
	row := repo.DB.QueryRow("SELECT key, name, description, logo, tags FROM skill where key=$1", key)
	
	err := row.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
	if err != nil {
		return sharedinterface.Skill{}, err
	}
	return skill, nil
}
func (repo SkillReposiory) InsertSkill(skill sharedinterface.Skill) (sharedinterface.Skill, error) {
	q := "INSERT INTO skill (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5) RETURNING key, name, description, logo, tags"
	row := repo.DB.QueryRow(q, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))
	var newSkill sharedinterface.Skill
	err := row.Scan(&newSkill.Key, &newSkill.Name, &newSkill.Description, &newSkill.Logo, pq.Array(&newSkill.Tags))
	if err != nil {
	
		return sharedinterface.Skill{}, err
	}
	return newSkill, nil
}
func (repo SkillReposiory) UpdateSkill(skill sharedinterface.Skill, key string) (sharedinterface.Skill, error) {
	q := "update skill set name = $2, description = $3, logo = $4, tags = $5 where key = $1 RETURNING key, name, description, logo, tags"
	// q := "INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id"
	row := repo.DB.QueryRow(q, key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))
	var newSkill sharedinterface.Skill 
	err := row.Scan(&newSkill.Key, &newSkill.Name, &newSkill.Description, &newSkill.Logo, pq.Array(&newSkill.Tags))

	if err != nil {

		return sharedinterface.Skill{}, err
	}
	return newSkill, nil
}
func (repo SkillReposiory) DeleteSkill(key string) error {
	row, err := repo.DB.Exec("DELETE FROM skill WHERE key = $1", key)
	if err != nil {
		return err
	}
	affect, err := row.RowsAffected() 
	if err != nil{	
		return err
	}

	if affect == 0 {
		return errors.New("afect row is 0")
	}
	return nil
}


func (repo SkillReposiory) PatchSkillDescription(description string, key string) (sharedinterface.Skill, error) {
	q := "update skill set description = $1 where key = $2 RETURNING key, name, description, logo, tags"
	row := repo.DB.QueryRow(q, description, key) 
	var newSkill sharedinterface.Skill 
	err := row.Scan(&newSkill.Key, &newSkill.Name, &newSkill.Description, &newSkill.Logo, pq.Array(&newSkill.Tags))

	if err != nil {

		return sharedinterface.Skill{}, err
	}

	return newSkill, nil
}
func (repo SkillReposiory) PatchSkillName(name string, key string) (sharedinterface.Skill, error) {
	q := "update skill set name = $1 where key = $2 RETURNING key, name, description, logo, tags"
	row := repo.DB.QueryRow(q, name, key) 
	var newSkill sharedinterface.Skill 
	err := row.Scan(&newSkill.Key, &newSkill.Name, &newSkill.Description, &newSkill.Logo, pq.Array(&newSkill.Tags))

	if err != nil {

		return sharedinterface.Skill{}, err
	}

	return newSkill, nil
}
func (repo SkillReposiory) PatchSkillTags(tags []string, key string) (sharedinterface.Skill, error) {
	q := "update skill set tags = $1 where key = $2 RETURNING key, name, description, logo, tags"
		row := repo.DB.QueryRow(q, pq.Array(tags), key) 
		var newSkill sharedinterface.Skill 
		err := row.Scan(&newSkill.Key, &newSkill.Name, &newSkill.Description, &newSkill.Logo, pq.Array(&newSkill.Tags))

		if err != nil {

			return sharedinterface.Skill{}, err
		}

		return newSkill, nil
}
func (repo SkillReposiory) PatchSkillLogo(logo string, key string) (sharedinterface.Skill, error) {
	q := "update skill set logo = $1 where key = $2 RETURNING key, name, description, logo, tags"
	row := repo.DB.QueryRow(q, logo, key) 
	var newSkill sharedinterface.Skill 
	err := row.Scan(&newSkill.Key, &newSkill.Name, &newSkill.Description, &newSkill.Logo, pq.Array(&newSkill.Tags))

	if err != nil {

		return sharedinterface.Skill{}, err
	}

	return newSkill, nil
}