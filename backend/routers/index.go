package routers

import(
	"github.com/gin-gonic/gin"
	"github.com/waranyafeen/Projectwaranrat/controller"
	"github.com/waranyafeen/Projectwaranrat/middlewares"
)

func SetupRouter() *gin.Engine {
	return gin.Default()
}

func InitRouter(route *gin.Engine) {
	route.Use(middlewares.CORS())

	authRouter := route.Group("/")
	initRequiredAuthRouter(authRouter)

}

func initRequiredAuthRouter(route *gin.RouterGroup) {
	route.Use(middlewares. Authorizes())

	user := middlewares. Authorizes()

	// employee system
	route.GET("/employees", controllers.GetAllEmployees)
	route.GET("/employees/:id", controllers.GetEmployees)
	route.POST("/employees", controllers.CreateEmployees)
	route.PUT("/employee/:id", controllers.UpdateEmployees)
	route.DELETE("/employee/:id", controllers.DeleteEmployees)

	route.GET("/employee/position", user, controllers.GetAllPositions)
	route.GET("/employee/precede", controllers.GetAllPrecedes)
	route.GET("/employee/gender", controllers.GetAllGendersForEmployee)

	// payment system
	route.GET("/payments", controllers.GetAllPayments)
	route.GET("/payments/:id", controllers.GetPayments)
	route.POST("/payments", controllers.CreatePayments)
	route.PUT("/payment/:id", controllers.UpdatePayments)
	route.DELETE("/payment/:id", controllers.DeletePayments)

	// ticket system
	route.GET("/tickets", controllers.GetAllTickets)
	route.GET("/tickets/:id", controllers.GetTickets)
	route.POST("/tickets", controllers.CreateTickets)
	route.PUT("/tickets/:id", controllers.UpdateTickets)
	route.DELETE("/tickets/:id", controllers.DeleteTickets)

	route.GET("/tickets/car", controllers.GetAllCars)
	route.GET("/tickets/departure", controllers.GetAllDepartures)

	// user system
	route.GET("/users", controllers.GetAllUsers)
	route.GET("/users/:id", controllers.GetUsers)
	route.POST("/users", controllers.CreateUsers)
	route.PUT("/users/:id", controllers.UpdateUsers)
	route.DELETE("/users/:id", controllers.DeleteUsers)

	route.GET("/users/role", controllers.GetAllRoles)
	route.GET("/employee/gender", controllers.GetAllGendersForUser)
}
