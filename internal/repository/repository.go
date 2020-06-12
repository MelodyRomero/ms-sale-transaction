package repository

import (
	"database/sql"
	"ms-sale-transaction/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type SaleRepository interface {
	FindAllSalesByMonth(val string) (*models.Sales, error)
	FindLatestSales(quantity int) (*models.Sales, error)
}

type saleRepository struct {
	Con *sql.DB
}

//sql.Open("mysql", dataSource)
func NewSqlSaleRepository(con *sql.DB) SaleRepository {

	return &saleRepository{
		Con: con,
	}
}

func (r *saleRepository) FindAllSalesByMonth(val string) (*models.Sales, error) {

	var sales models.Sales
	sales = make(models.Sales, 0)

	querystring := `SELECT * FROM sales WHERE sales.date = ` + val
	rows, err := r.Con.Query(querystring)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Sale
		err := rows.Scan(&r.ID, &r.Date, &r.Product, &r.Total)
		if err != nil {
			return nil, err
		}
		sales = append(sales, &r)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &sales, nil
}

func (r *saleRepository) FindLatestSales(quantity int) (*models.Sales, error) {
	var sales *models.Sales

	rows, err := r.Con.Query("SELECT top ? * FROM sales ", quantity)

	if err := rows.Scan(&sales); err != nil {
		return nil, err
	}

	rerr := rows.Close()
	if rerr != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sales, nil
}
