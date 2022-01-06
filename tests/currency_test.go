package tests

import (
	"currency-converter/pkg/controllers"
	"currency-converter/pkg/model"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	t.Run("server health check", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/health", nil)
		response := httptest.NewRecorder()
		controllers.HealthCheck(response, request)
		expected := "OK"
		actual := response.Body.String()
		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})
}

func TestUnsupportedCurrencyConversionWebMethod(t *testing.T) {
	t.Run("unsupported web method", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/convert-ccy", nil)
		response := httptest.NewRecorder()
		controllers.GetExcahngeRate(response, request)
		expectedStatusCode := 405
		actualStatusCode := response.Result().StatusCode
		if actualStatusCode != expectedStatusCode {
			t.Errorf("expected %d, got %d", expectedStatusCode, actualStatusCode)
		}
	})
}

func TestCurrencyConversionWithoutParameters(t *testing.T) {
	t.Run("currency conversion without parameters", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/convert-ccy", nil)
		response := httptest.NewRecorder()
		controllers.GetExcahngeRate(response, request)
		expectedStatusCode := 400
		actualStatusCode := response.Result().StatusCode
		if actualStatusCode != expectedStatusCode {
			t.Errorf("expected %d, got %d", expectedStatusCode, actualStatusCode)
		}
	})
}

func TestUnsupportedCurrencyConversion(t *testing.T) {
	t.Run("unsupported currency paramters", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/convert-ccy?source=USD&target=GBP", nil)
		response := httptest.NewRecorder()
		controllers.GetExcahngeRate(response, request)
		expectedStatusCode := 400
		actualStatusCode := response.Result().StatusCode
		if actualStatusCode != expectedStatusCode {
			t.Errorf("response code; expected %d, got %d", expectedStatusCode, actualStatusCode)
		}
		actualResponseBody := response.Body.String()
		expectedResponseBody := "invalid currency: USD"
		if actualResponseBody != expectedResponseBody {
			t.Errorf("response body: expected %s, got %s", expectedResponseBody, actualResponseBody)
		}
	})
}

func TestSameCurrencyConversion(t *testing.T) {
	t.Run("same currency conversion", func(t *testing.T) {
		if err := model.SetupDatabase(true); err != nil {
			fmt.Printf("database error: %+v", err)
			panic(err)
		}
		request, _ := http.NewRequest("GET", "/convert-ccy?source=KSH&target=Ksh", nil)
		response := httptest.NewRecorder()
		controllers.GetExcahngeRate(response, request)
		expectedStatusCode := 200
		actualStatusCode := response.Result().StatusCode
		if actualStatusCode != expectedStatusCode {
			t.Errorf("response code; expected %d, got %d", expectedStatusCode, actualStatusCode)
		}
		actualResponseBody := response.Body.String()
		exceptedResponseBody := "rate: 1.000"
		if actualResponseBody != exceptedResponseBody {
			t.Errorf("response body; expected %s, got %s", exceptedResponseBody, actualResponseBody)
		}
	})
}

func TestSimpleConversion(t *testing.T) {
	t.Run("simple conversion test", func(t *testing.T) {
		if err := model.SetupDatabase(true); err != nil {
			fmt.Printf("database error: %+v", err)
			panic(err)
		}
		request, _ := http.NewRequest("GET", "/convert-ccy?source=NGN&target=KSH", nil)
		response := httptest.NewRecorder()
		controllers.GetExcahngeRate(response, request)
		expectedStatusCode := 200
		actualStatusCode := response.Result().StatusCode
		if actualStatusCode != expectedStatusCode {
			t.Errorf("response code; expected %d, got %d", expectedStatusCode, actualStatusCode)
		}
		actualResponseBody := response.Body.String()
		exceptedResponseBody := "rate: 0.270"
		if actualResponseBody != exceptedResponseBody {
			t.Errorf("response body; expected %s, got %s", exceptedResponseBody, actualResponseBody)
		}
	})
}
