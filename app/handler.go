package app

import (
	"alochym/service"
	"fmt"
	"net/http"
)

type CustomerHandler struct {
	svcRepo service.CustomerServiceRepository
}

func (c CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applocation/json")
	cs, _ := c.svcRepo.GetAllCustomers()
	fmt.Fprint(w, cs)

}

// func (c CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	w.Header().Add("Content-Type", "applocation/json")
// 	cs, err := c.svcRepo.GetCustomer(params["id"])
// 	if err != nil {
// 		w.Header().Add("Content-Type", "applocation/json")
// 		fmt.Fprint(w, err.Error())
// 		return
// 	}
// 	w.Header().Add("Content-Type", "applocation/json")
// 	fmt.Fprint(w, cs)

// }
func NewCustomerHandler(svcRepo service.CustomerServiceRepository) CustomerHandler {
	return CustomerHandler{svcRepo: svcRepo}
}
