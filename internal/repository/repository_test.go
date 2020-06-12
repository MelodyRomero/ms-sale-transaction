package repository_test

import (
	"fmt"
	"ms-sale-transaction/internal/models"
	"ms-sale-transaction/internal/repository"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {

	t.Run("Escenario: Repositorio responde", func(t *testing.T) {
		// Open new mock database
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			fmt.Println("failed to open sqlmock database:", err)
		}
		defer db.Close()

		repo := repository.NewSqlSaleRepository(db)
		// columns to be used for result
		rows := sqlmock.NewRows([]string{"id", "fecha", "producto", "total"}).
			AddRow("1", "enero", "auto", 0.98)

		// expect transaction begin
		mock.ExpectBegin()
		db.Begin()
		mock.ExpectQuery("SELECT * FROM sales WHERE sales.fecha = enero").WillReturnRows(rows)

		response, err := repo.FindAllSalesByMonth("enero")

		if err != nil {
			fmt.Println(err)
		}
		// ensure all expectations have been met
		if err = mock.ExpectationsWereMet(); err != nil {
			fmt.Printf("unmet expectation error: %s", err)
		}

		assert.Equal(t, models.Sales(models.Sales{&models.Sale{ID: "1", Date: "enero", Product: "auto", Total: 0.98}}), *response)
		assert.NotNil(t, response)
		assert.Nil(t, err)

	})
}
