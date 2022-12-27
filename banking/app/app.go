package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"log"
	"net/http"
	"os"

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

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	//start server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
