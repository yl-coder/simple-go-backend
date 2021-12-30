package app

import (
	"banking/errs"
	"banking/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func getTime(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Has("tz") {

		tz := r.URL.Query().Get("tz")

		location, err := time.LoadLocation(tz)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("invalid timezone"))
			return
		}

		locationStr := time.Now().In(location).String()
		print(fmt.Fprintf(w, locationStr))

	} else {
		locationStr := time.Now().In(time.UTC).String()
		print(fmt.Fprintf(w, locationStr))
	}
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	statusFilter := sql.NullString{}

	if r.URL.Query().Has("status") {
		statusFilter.String = r.URL.Query().Get("status")
		statusFilter.Valid = true
	}
	customers, errs := ch.service.GetAllCustomer(statusFilter)
	writeResponseHelper(w, errs, customers)
}

func writeResponseHelper(w http.ResponseWriter, errs *errs.AppError, data interface{}) {
	if errs != nil {
		writeResponse(w, errs.Code, errs.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, data)
	}
}


func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customers, errs := ch.service.GetCustomer(vars["customer_id"])
	writeResponseHelper(w, errs, customers)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}


