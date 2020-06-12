package service_test

import (
	"ms-sale-transaction/internal/models"
	"ms-sale-transaction/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {

	salestotest := []*models.Sale{
		{
			ID:      "1",
			Date:    "enero",
			Product: "laptop",
			Total:   3000000,
		},
		{
			ID:      "2",
			Date:    "febrero",
			Product: "laptop",
			Total:   120302131232,
		},
	}

	t.Run("Escenario: GetSales con el parametro month = enero responde correctamente ", func(t *testing.T) {
		svc := service.NewSalesService(salestotest)

		respuesta := svc.GetSales("enero")

		assert.NotNil(t, respuesta)
		assert.Equal(t, "1", respuesta[0].ID)
		assert.Equal(t, "laptop", respuesta[0].Product)
	})

	t.Run("Escenario: GetSales con el parametro month = enero no responde correctamente ", func(t *testing.T) {
		svc := service.NewSalesService(salestotest)

		respuesta := svc.GetSales("enero")

		assert.NotNil(t, respuesta)
		assert.Equal(t, "1", respuesta[0].ID)
		assert.NotEqual(t, "microfono", respuesta[0].Product)
	})
}
