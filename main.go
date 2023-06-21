package main

import (
	"fmt"
	"os"

	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/infrastructure/db"
	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/infrastructure/db/transaction"
	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/internal/helpers/pagination"
	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/internal/helpers/profileconverter"
	httpprotocol "github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/internal/http-protocol"
	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/internal/logs"
	httphand "github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/src/handlers/http"
	profilerepo "github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/src/modules/profile/repository"
	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/src/modules/profile/services"
)

func main() {
	// time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	// panic(time)
	dir, _ := os.Getwd()
	dir = fmt.Sprintf("%s/internal/logs", dir)
	logs.InitLogger(logs.Config{
		ConsoleLoggingEnabled: true,
		EncodeLogsAsJson:      true,
		FileLoggingEnabled:    true,
		Directory:             dir,
		Filename:              "logging.log",
		MaxSize:               200000000,
		MaxBackups:            2000,
		MaxAge:                2000,
	})

	mysqlDB := db.NewMysqlConnection()
	transaction := transaction.NewTransactionImpl(mysqlDB)
	profileCVT := profileconverter.NewProfileConverterImpl()
	paginate := pagination.NewPaginationImpl()

	profileRepo := profilerepo.NewProfileRepositoryImpl()
	profileService := services.NewProfileServiceImpl(mysqlDB, transaction, profileRepo, profileCVT, paginate)
	httpHandler := httphand.NewHttpHandlerImpl(profileService)
	http := httpprotocol.NewHttpImpl(httpHandler)
	http.Listen()
}
