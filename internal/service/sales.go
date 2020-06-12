package service

import "ms-sale-transaction/internal/models"

type SalesService interface {
	GetSales(string) []*models.Sale
	GetSalesByProduct(string) []*models.Sale
	GetLatestSales(int) []*models.Sale
}

type salesService struct {
	repositorio []*models.Sale
}

func NewSalesService(repo []*models.Sale) SalesService {

	return &salesService{
		repositorio: repo,
	}
}

func (s *salesService) GetSales(month string) []*models.Sale {
	var sales []*models.Sale
	for _, sale := range s.repositorio {
		if sale.Date == month {
			sales = append(sales, sale)
		}
	}
	return sales
}

func (s *salesService) GetLatestSales(quantity int) []*models.Sale {

	a := s.repositorio[len(s.repositorio)-quantity : len(s.repositorio)]

	return a
}

func (s *salesService) GetSalesByProduct(product string) []*models.Sale {
	var sales []*models.Sale
	for _, sale := range s.repositorio {
		if sale.Date == product {
			sales = append(sales, sale)
		}
	}
	return sales
}
