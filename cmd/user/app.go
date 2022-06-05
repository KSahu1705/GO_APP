package main

import (
	// "context"
	// "flag"
	"fmt"
	// "net"
	// "net/http"
	// "os"
	// "strings"
	// "sync"
	// "time"

	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq"
	"GO_APP/internal/delivery/api/user"
	"GO_APP/internal/repository/postgres"
	"GO_APP/internal/usecase"

	"GO_APP/internal/model/config"



	"GO_APP/internal/delivery"
	"GO_APP/internal/repository"
	"gopkg.in/tokopedia/grace.v1"

)

var (
	err                           error
	cfg                           *config.Config
	userRepository                repository.User
	userUsecase					  usecase.UserCase
	userAPI                   	  delivery.API
	// detectionAPI                  delivery.API
	host                          string
)

func init() {
	// initialize app configuration
	cfg = config.NewConfig()
	if err1 := cfg.ReadConfig(); err != nil {
		fmt.Errorf("[Init][Configuration] %v", err1)
	} 

	// host, err = os.Hostname()
	// if err != nil {
	// 	fmt.Errorf("[Init][Get Hostname]Failed to get the hostname %+v", err)
	// }

	// fmt.Println(cfg.Database)
    // for no, month := range cfg.Database {
    //     fmt.Print(no)
    //     fmt.Println("-" , month)
    // }
	// initialize postgres and repository
	if userRepository, err := postgres.NewPostgres(cfg.Database); err != nil {
		fmt.Println(userRepository)
		fmt.Println("@@@")
		fmt.Println(err.Error())

	} else {
		fmt.Println("###")
		fmt.Println("[Init][Postgres] successfully connected.")
	}
	fmt.Println("-->",err)
	userAPI = user.NewAPI(cfg, userUsecase, "ENV")

}

func main() {
	router := httprouter.New()
	userAPI.RegisterRoutes(router)
	// detectionAPI.RegisterRoutes(router)

	// if err = categoryRepository.ConstructCategory(); err != nil {
	// 	log.Fatalf("[Init][Construct-Category] %v", err)
	// } else {
	// 	log.Infof("[Init][Construct-Category] category successfully constructed")
	// }

	// if err = categoryRepository.ConstructCategoryImage(); err != nil {
	// 	log.Fatalf("[Init][Construct-Category-Image] %v", err)
	// } else {
	// 	log.Infof("[Init][Construct-Category-Image] category image successfully constructed")
	// }

	// if err = categoryRepository.GetCategoryRedirectionMap(); err != nil {
	// 	log.Fatalf("[Init][GetCategoryRedirectionMap] %v", err)
	// } else {
	// 	log.Infof("[Init][GetCategoryRedirectionMap] redirection map successfully constructed")
	// }

	grace.Serve(":"+cfg.Server.Port, router)
}

