package app

import (
	service "banking/service"
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.ICustomerService
}

func (ch *CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomer()

	writeResponse(w, r, http.StatusOK, customers)
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		writeResponse(w, r, err.Code, err.AsMessage())
	} else {
		writeResponse(w, r, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		if err := xml.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	}
}
