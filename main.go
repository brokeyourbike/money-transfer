package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/goremit/money-transfer/api/handlers"
	"github.com/goremit/money-transfer/db"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)

	if err := run(log); err != nil {
		log.Fatalf("%s\n", err)
		os.Exit(1)
	}
}

func run(*logrus.Logger) error {
	orm, err := gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		return err
	}

	// repos
	trxRepo := db.NewTransactionsRepo(orm)

	// handlers
	customers := handlers.NewCustomers(trxRepo)

	r := chi.NewRouter()
	r.Route("customers", func(r chi.Router) {
		r.Route("transactions", func(r chi.Router) {
			r.Post("/", customers.HandleCreateTransaction)
			r.Route("{id}", func(r chi.Router) {
				r.Get("/", customers.HandleShowTransaction)
				r.Post("/fund", customers.HandleFundTransaction)
				r.Delete("/", customers.HandleCancelTransaction)
			})
		})
	})

	return http.ListenAndServe("127.0.0.1:9090", r)
}
