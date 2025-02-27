package main

import (
	"fmt"
	"log"
	"peanut/config"
	"peanut/infra"
	"peanut/middleware"
	"peanut/pkg/i18n"

	"gorm.io/gorm"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	fmt.Println("---- Hello world! ----")

	config.Setup()
	i18n.SetupI18n()
	middleware.ValidateFunction()

	dbClient := dbConnect()
	server := infra.SetupServer(dbClient)
	infra.Migration(dbClient)

	err := server.Router.Run(":8080")
	if err != nil {
		log.Fatal("Start fail")
	}
}

func dbConnect() *gorm.DB {
	db, err := infra.PostgresOpen()
	if err != nil {
		log.Fatal("[main] DB connect error: ", err)
	}
	return db
}
