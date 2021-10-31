package app

import (
	"alochym/domain"
	"alochym/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// Customer stubs
	repoCustomer := domain.NewCustomerRepositoryStub()
	svcRepoCustomer := service.NewCustomerServer(repoCustomer)
	hCustomer := NewCustomerHandler(svcRepoCustomer)

	// Customer routes
	router.HandleFunc("/customers", hCustomer.getAllCustomers).Methods(http.MethodGet)
	// Validation URL only contains numberics id - {id:[0-9]+}
	// router.HandleFunc("/customers/{id:[0-9]+}", hCustomer.getCustomer).Methods(http.MethodGet)

	fmt.Println("Start server with port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
