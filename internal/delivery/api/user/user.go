package user

import (
	"github.com/julienschmidt/httprouter"
	"GO_APP/internal/delivery"
	"GO_APP/internal/model/config"
	"GO_APP/internal/usecase"

)

// // App has router and db instances
// type App struct {
// 	Router *httprouter.Router//mux.Router
// 	DB     *sqlx.DB
// }

// // App initialize with predefined configuration
// func (a *App) Initialize(config *config.Config) {
// 	dbURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		config.DB.Host,
// 		config.DB.Port,
// 		config.DB.User,
// 		config.DB.Password,
// 		config.DB.DBname,
// 	)

// 	db, err := sqlx.Connect(config.DB.Dialect, dbURI)
// 	if err != nil {
// 		log.Fatal("Could not connect database")
// 	} else {
// 		fmt.Printf("Connected to database\n")
// 	}

// 	a.DB = model.DBMigrate(db)
// 	a.Router = httprouter.New()
// 	a.setRouters()
// }

// // https://github.com/gin-gonic/gin/issues/1681
// // Set all required routers
// func (a *App) setRouters() {
// 	router := a.Router
// 	// Routing for handling the projects
// 	router.GET("/users", a.GetAllUser)
// 	router.GET("/users/:id", a.GetUser)
// 	router.GET("/users/:id/address", a.GetUserAddress)
// 	router.POST("/users", a.CreateUser)
// 	router.POST("/users/:id/add_address", a.CreateUserAddress)
// 	router.PUT("/users/:id/update_user", a.UpdateUser)
// 	router.PUT("/users/:id/update_address/:addr_id", a.UpdateUserAddress)
// 	router.PUT("/users/:id/disable", a.DisableUser)
// 	router.PUT("/users/:id/enable", a.EnableUser)
// 	router.DELETE("/users/:id", a.DeleteUser)
// 	router.DELETE("/users/:id/del/:addr_id", a.DeleteUserAddress)
// }


//API is a struct
type API struct {
	UserCase     usecase.UserCase
	// Metric       *metrics.Client
	Config       *config.Config
	Environment  string
}

//NewAPI returns struct of API
func NewAPI(cfg *config.Config, userUsecase usecase.UserCase,  environment string) delivery.API {
	return &API{
		Config:       cfg,
		UserCase: 	  userUsecase,
		// Metric:       metric,
		Environment:  environment,
	}
}

//RegisterRoutes is a method for router
func (user *API) RegisterRoutes(r *httprouter.Router) {
	r.POST("/users/:id", user.DetailHandler)
	// r.GET("/category/v1/detail/:id_identifier", panics.CaptureHTTPRouterHandler(category.DetailHandler))
	// r.GET("/category/v1/list/:id_identifier", panics.CaptureHTTPRouterHandler(category.ListHandler))
	// r.GET("/category/v1/tree/:id_identifier", panics.CaptureHTTPRouterHandler(category.TreeHandler))
	// r.GET("/category/v1/lite", panics.CaptureHTTPRouterHandler(category.LiteHandler))
	// r.GET("/category/v1/tokonow/tree/:id_identifier", panics.CaptureHTTPRouterHandler(category.TokoNowTreeHandler))
}

