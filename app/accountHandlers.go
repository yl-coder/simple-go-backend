package app

import (
	"banking/domain"
	"banking/errs"
	"banking/logger"
	"banking/service"
	"net/http"
	"strconv"
)

type AccountHandlers struct {
	service service.AccountService
}

const ErrorMsgProvideCustomerIdAccountType = "Please provide query param customerId and accountType"
const ErrorMessageParsingCustomerId = "Error parsing customerId"
const ErrorMessageProvideAccountIdTransactionType = "Please provide query param accountId and transactionType"
const ErrorMsgParsingAccountIdAmount = "Error parsing accountId or amount"

func (ah *AccountHandlers) createAccount(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Has("customerId") == false || r.URL.Query().Has("accountType") == false {

		writeResponseHelper(w, errs.NewBadRequestError(ErrorMsgProvideCustomerIdAccountType), nil)
	}

	customerId, err := strconv.Atoi(r.URL.Query().Get("customerId"))

	if err != nil {

		writeResponseHelper(w, errs.NewBadRequestError(ErrorMessageParsingCustomerId), nil)
	}


	accountType := r.URL.Query().Get("accountType")

	appError := ah.service.CreateAccount(int64(customerId), domain.AccountType(accountType))

	if appError != nil {
		writeResponseHelper(w, appError, nil)
	}

	logger.Info("Created account")
	writeResponseHelper(w, nil, nil)
}

func (ah *AccountHandlers) createTransaction(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Has("transactionType") == false || r.URL.Query().Has("accountId") == false || r.URL.Query().Has("amount") == false {

		writeResponseHelper(w, errs.NewBadRequestError(ErrorMessageProvideAccountIdTransactionType), nil)
	}

	accountId, err := strconv.Atoi(r.URL.Query().Get("accountId"))

	amount, err := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)

	if err != nil {
		writeResponseHelper(w, errs.NewBadRequestError(ErrorMsgParsingAccountIdAmount), nil)
	}
	transactionType := r.URL.Query().Get("transactionType")

	appError := ah.service.CreateTransaction(int64(accountId), amount, domain.TransactionType(transactionType))

	if appError != nil {
		writeResponseHelper(w, appError, nil)
		return
	}

	logger.Info("Created transaction")
	writeResponseHelper(w, nil, nil)
}


