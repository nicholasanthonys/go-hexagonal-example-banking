package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nicholasanthonys/hexagonal-banking/domain"
	"github.com/nicholasanthonys/hexagonal-banking/service"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")

	// starting a server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}
