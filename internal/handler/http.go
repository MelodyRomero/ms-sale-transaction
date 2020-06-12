package handler

import (
	"encoding/json"
	"ms-sale-transaction/internal/models"
	"ms-sale-transaction/internal/service"
	"net/http"
	"strconv"
)

type SalesHandler interface {
	GetSalesHandler(string) http.HandlerFunc
	GetSalesByProductHandler(string) http.HandlerFunc
	GetLatestSalesHandler(string) http.HandlerFunc
}
type salesHandler struct {
	svc service.SalesService
}

func NewSalesHandler(repo []*models.Sale) SalesHandler {
	return &salesHandler{
		svc: service.NewSalesService(repo),
	}
}

func (s salesHandler) GetSalesHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		month := (r.URL.Query().Get("month"))
		if month == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)
			return
		}
		service := s.svc.GetSales(month)

		raw, err := json.Marshal(service)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(raw))
	}
}

func (s salesHandler) GetLatestSalesHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q, err := strconv.Atoi(r.URL.Query().Get("quantity"))
		if err != nil || q < 1 {
			http.NotFound(w, r)
			return
		}

		service := s.svc.GetLatestSales(q)
		raw, err := json.Marshal(service)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(raw))
	}
}
func (s salesHandler) GetSalesByProductHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		product := (r.URL.Query().Get("product"))

		if product == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)
			return
		}
		service := s.svc.GetSalesByProduct(product)

		raw, err := json.Marshal(service)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(raw))
	}
}
