package controllers

import (
	"currency-converter/pkg/services"
	"fmt"
	"net/http"
)

func GetExcahngeRate(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		ccy1 := req.URL.Query().Get("source")
		ccy2 := req.URL.Query().Get("target")
		retval, err := services.ProcessConversion(ccy1, ccy2)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			resp.Write([]byte(fmt.Sprintf("%+v", err)))
			return
		}
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte(fmt.Sprintf("result: %+v", retval)))
	default:
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("method not allowed"))
	}
}
