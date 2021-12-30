package app

import (
	"banking/domain"
	"banking/logger"
	"banking/service"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" {
		panic("SERVER_ADDRESS NOT DEFINED")
	}
}

func migrateFile() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	address := os.Getenv("SERVER_ADDRESS")
	port := "3306"
	connectionString :=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", dbUser, dbPassword, address, port, dbName)
	logger.Info("Migrating " + connectionString)
	db, _ := sql.Open("mysql", connectionString)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migration/",
		"mysql",
		driver,
	)

	if err != nil {
		logger.Info("Error migration " + err.Error())
	}

	if m == nil {
		logger.Info("M is nil")
	}

	err = m.Up()

	if err != nil {
		logger.Info("Error in migration " + err.Error())
	}
}

func Start() {
	logger.Info("Starting our application")
	sanityCheck()
	migrateFile()

	mux := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	ah := AccountHandlers{service.NewAccountServiceImpl(domain.NewAccountRepositoryDb(), domain.NewTransactionRepositoryDb())}
	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/api/time", getTime).Methods(http.MethodGet)
	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	mux.HandleFunc("/account", ah.createAccount).Methods(http.MethodPost)
	mux.HandleFunc("/transaction", ah.createTransaction).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info("Server starting")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), mux))
}
