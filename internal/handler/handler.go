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
    // Ограничение размера multipart (например, 10 МБ)
    router.MaxMultipartMemory = 10 << 20

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.Static("/uploads", "./uploads")
	router.StaticFile("/map", "./map.html")

	user := router.Group("/user")
	{
		user.POST("/sign-up", h.SignUp)
		user.POST("/sign-in", h.SignIn)
		user.GET("/leaders", h.GetUsersSortedByPoints)
		user.GET("/change-password", h.ChangePasswordByMail)
	}

	authorized := router.Group("/authorized", h.UserIdentity)
	{
		authorized.GET("/user", h.UserByID)

		task := authorized.Group("/task")
		{
			task.POST("/create", h.CreateTask)
			task.POST("/take", h.TakeTask)
			task.POST("/complete", h.CompleteTask)
		}

		message := authorized.Group("/message")
		{
			message.GET("", h.GetMessageByUserID)
			message.GET("/read", h.ReadMessageByID)
		}

		authorized.GET("/report-by-id", h.GetReportById)
		authorized.POST("/task-photo", h.UploadTaskPhoto)
		authorized.POST("/report-photo", h.UploadReportPhoto)
		authorized.POST("/answers", h.CheckAnswers)
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
		global.GET("/question", h.GetQuestionById)
		global.GET("/test", h.GenerateTest)
		global.GET("/analise", h.GetTasksForAnalise)
		global.GET("/tasks-map", h.GetTasksWithCoordinates)
		global.GET("/map-points", h.GetMapPoints)
		global.GET("/points-summary", h.GetPointsSummary)
	}

	violation := router.Group("/violation")
	{
		violation.GET("/all", h.GetAllViolations)
		violation.GET("", h.GetViolationByCategory)
		violation.GET("/by-id", h.GetViolationByID)
	}
	return router
}
