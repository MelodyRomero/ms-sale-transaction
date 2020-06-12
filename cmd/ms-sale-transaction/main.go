package main

import (
	"log"
	"ms-sale-transaction/internal/handler"
	"ms-sale-transaction/internal/models"
	mux "ms-sale-transaction/internal/router"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	var repo []*models.Sale

	repo = append(repo,
		&models.Sale{ID: "1", Date: "enero", Product: "auto", Total: 0.0},
		&models.Sale{ID: "2", Date: "febrero", Product: "auto", Total: 0.0},
	)

	h := handler.NewSalesHandler(repo)
	router.GET("/sales/product/", h.GetSalesByProductHandler(""))
	router.GET("/sales/month/", h.GetSalesHandler(""))
	router.GET("/sales/latest/", h.GetLatestSalesHandler(""))
	log.Fatal(http.ListenAndServe(":8080", router))
}
