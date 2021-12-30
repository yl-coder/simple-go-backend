package app

import (
	"banking/dto"
	"banking/errs"
	"banking/mocks/service"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Given_ValidCustomer_When_MethodGetAllCustomers_Then_Status200(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockCustomerService(ctrl)
	customers := []dto.CustomerResponse{
		{"1001", "Daniel", "Singapore", "12345", "13-11-2020", "Ok"},
	}

	mockService.EXPECT().GetAllCustomer(sql.NullString{}).Return(customers, nil)

	ch := CustomerHandlers{service: mockService}

	router := mux.NewRouter()

	router.HandleFunc("/customers", ch.getAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Should be statusOK")
	}
}



func Test_Given_InvalidDbState_When_MethodGetAllCustomers_Then_Status500(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockCustomerService(ctrl)

	mockService.EXPECT().GetAllCustomer(sql.NullString{}).Return(nil, errs.NewUnexpectedError("Unexpected error"))

	ch := CustomerHandlers{service: mockService}

	router := mux.NewRouter()

	router.HandleFunc("/customers", ch.getAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Should be StatusInternalServerError")
	}
}
