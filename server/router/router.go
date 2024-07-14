package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thanakize/lab_skill_api/server/controllers"
	"github.com/thanakize/lab_skill_api/usecase"
)

func InitRoute(r *gin.Engine, usecase usecase.IUseCase) {

	router := r.Group("/api/v1/skills")
	
	router.GET("/", controllers.GetController(usecase))
	router.GET("/:key", controllers.GetKeyController(usecase))
	router.POST("/", controllers.PostController(usecase))
	router.PUT("/:key", controllers.PutController(usecase))
	router.DELETE("/:key", controllers.DeleteController(usecase))
	router.PATCH("/:key/actions/name",controllers.PatchNameController(usecase))
	router.PATCH("/:key/actions/description",controllers.PatchDescriptionController(usecase))
	router.PATCH("/:key/actions/logo",controllers.PatchLogoController(usecase))
	router.PATCH("/:key/actions/tags",controllers.PatchTagsController(usecase))
}