package controllers

import (
	"currency-converter/pkg/services"
	"fmt"
	"net/http"
)

// GET /convert-ccy?source={ccy1}&target={ccy2}
//
// Perform exchange rate conversion from source to target currencies
func GetExcahngeRate(resp http.ResponseWriter, req *http.Request) {
	// permits only GET requests
	switch req.Method {
	case "GET":
		// retrieve currencies from query parameters
		sourceCcy := req.URL.Query().Get("source")
		targetCcy := req.URL.Query().Get("target")
		// process currency conversion
		retval, err := services.ProcessConversion(sourceCcy, targetCcy)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			resp.Write([]byte(fmt.Sprintf("%+v", err)))
			return
		}
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte(fmt.Sprintf("rate: %+v", retval)))
	default:
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("method not allowed"))
	}
}

// GET /health
//
// Peform health information
func HealthCheck(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("OK"))
}
