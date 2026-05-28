package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	healthHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", rr.Code)
	}
}

func TestGetOrdersHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	rr := httptest.NewRecorder()

	getOrdersHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", rr.Code)
	}
}