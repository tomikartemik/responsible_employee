package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"responsible_employee/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.Static("/uploads", "./uploads")

	user := router.Group("/user")
	{
		user.POST("/sign-up", h.SignUp)
		user.POST("/sign-in", h.SignIn)
		user.GET("/leaders", h.GetUsersSortedByPoints)
	}

	authorized := router.Group("/authorized", h.UserIdentity)
	{
		authorized.GET("/user", h.UserByID)
		authorized.POST("/create-task", h.CreateTask)
		authorized.POST("/take-task", h.TakeTask)
		authorized.POST("/complete-task", h.CompleteTask)
		authorized.POST("/task-photo", h.UploadTaskPhoto)
		authorized.POST("/report-photo", h.UploadReportPhoto)
	}

	task := router.Group("/task", h.UserIdentity)
	{
		task.POST("", h.CreateTask)
		task.POST("/take", h.TakeTask)
		task.POST("/complete", h.CompleteTask)
	}

	global := router.Group("/global")
	{
		global.GET("/tasks", h.GetAllTasks)
		global.GET("/task", h.GetTaskById)
	}

	violation := router.Group("/violation")
	{
		violation.GET("/all", h.GetAllViolations)
		violation.GET("", h.GetViolationByCategory)
		violation.GET("/by-id", h.GetViolationByID)
	}
	return router
}
