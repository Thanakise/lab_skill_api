package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanakize/lab_skill_api/sharedinterface"
	"github.com/thanakize/lab_skill_api/usecase"
)

func GetController(usecase usecase.IUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) { 
		todos, err := usecase.GetSkills()
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, sharedinterface.SkillsResponse{
			Status: "success",
			Data: todos,
		})
	}
}

func GetKeyController(usecase usecase.IUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) { 
		key := ctx.Param("key")
		skill, err := usecase.GetSkill(key)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: "Skill not found",
			})
			return
		}
		ctx.JSON(http.StatusOK, sharedinterface.SkillResponse{
			Status: "success",
			Data: skill,
		})
	}
}
func PostController(usecase usecase.IUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) { 
		var skill sharedinterface.Skill
		if err := ctx.ShouldBindJSON(&skill); err != nil {
			ctx.JSON(http.StatusBadRequest, sharedinterface.ErrorResponse{
				Status: "error",
				Message: err.Error(),
			})
			return
		}
		newSkill, err := usecase.InsertSkill(skill)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: "Skill already exists",
			})
			return
		}
		ctx.JSON(http.StatusOK, sharedinterface.SkillResponse{
			Status: "success",
			Data: newSkill,
		})
	}
}
func PutController(usecase usecase.IUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) { 
		key := ctx.Param("key")
		var skill sharedinterface.Skill
		if err := ctx.ShouldBindJSON(&skill); err != nil {
			log.Println(err.Error())

			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: err.Error(),
			})
			return
		}
		newSkill, err := usecase.UpdateSkill(skill, key)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: "not be able to update skill",
			})
			return
		}
		ctx.JSON(http.StatusOK, sharedinterface.SkillResponse{
			Status: "success",
			Data: newSkill,
		})
	}
}
func DeleteController(usecase usecase.IUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) { 
		key := ctx.Param("key")
		err := usecase.DeleteSkill(key)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: "not be able to delete skill",
			})
			return
		}
		ctx.JSON(http.StatusOK, sharedinterface.DeleteResponse{
			Status: "success",
			Message : "Skill deleted",
		})
	}
}
func PatchNameController(usecase usecase.IUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) { 
		key := ctx.Param("key")
		var skill sharedinterface.Skill
		if err := ctx.ShouldBindJSON(&skill); err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: err.Error(),
			})
			return
		}
		newSkill, err := usecase.PatchSkillName(skill.Name, key)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: "not be able to update skill name",
			})
			return
		}
		ctx.JSON(http.StatusOK, sharedinterface.SkillResponse{
			Status: "success",
			Data: newSkill,
		})
	}
}
func PatchDescriptionController(usecase usecase.IUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) { 
		key := ctx.Param("key")
		var skill sharedinterface.Skill
		if err := ctx.ShouldBindJSON(&skill); err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: err.Error(),
			})
			return
		}
		newSkill, err := usecase.PatchSkillDescription(skill.Description, key)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: "not be able to update skill description",
			})
			return
		}
		ctx.JSON(http.StatusOK, sharedinterface.SkillResponse{
			Status: "success",
			Data: newSkill,
		})
	}
}
func PatchLogoController(usecase usecase.IUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) { 
		key := ctx.Param("key")
		var skill sharedinterface.Skill
		if err := ctx.ShouldBindJSON(&skill); err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: err.Error(),
			})
			return
		}
		newSkill, err := usecase.PatchSkillLogo(skill.Logo, key)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: "not be able to update skill logo",
			})
			return
		}
		ctx.JSON(http.StatusOK, sharedinterface.SkillResponse{
			Status: "success",
			Data: newSkill,
		})
	}
}
func PatchTagsController(usecase usecase.IUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) { 
		key := ctx.Param("key")
		var skill sharedinterface.Skill
		if err := ctx.ShouldBindJSON(&skill); err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: err.Error(),
			})
			return
		}
		newSkill, err := usecase.PatchSkillTags(skill.Tags, key)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(400, sharedinterface.ErrorResponse{
				Status: "error",
				Message: "not be able to update skill tags",
			})
			return
		}
		ctx.JSON(http.StatusOK, sharedinterface.SkillResponse{
			Status: "success",
			Data: newSkill,
		})
	}
}