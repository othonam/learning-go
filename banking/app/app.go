package app

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	//define routes
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//start server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
