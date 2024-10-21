package repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"mycrudapp/internal/models"
)

type CustomerRepository struct {
	dbPool *pgxpool.Pool
}

func NewCustomerRepository(dbPool *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{
		dbPool: dbPool,
	}
}

func (c *CustomerRepository) GetById(ctx context.Context, id int64) (models.Customer, error) {
	var customer models.Customer

	err := c.dbPool.QueryRow(ctx, "SELECT * FROM customers WHERE id = $1", id).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Phone,
		&customer.Active,
		&customer.Created,
	)

	return customer, err
}

func (c *CustomerRepository) GetAll(ctx context.Context) ([]models.Customer, error) {
	var customers []models.Customer

	rows, err := c.dbPool.Query(ctx, "SELECT * FROM customers")

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var customer models.Customer

		err = rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Phone,
			&customer.Active,
			&customer.Created,
		)

		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (c *CustomerRepository) Create(ctx context.Context, customer models.Customer) (models.Customer, error) {
	err := c.dbPool.QueryRow(ctx, "INSERT INTO customers (name, phone, active) VALUES ($1, $2, $3) RETURNING id, created",
		customer.Name, customer.Phone, customer.Active).Scan(&customer.ID, &customer.Created)

	return customer, err
}

func (c *CustomerRepository) Update(ctx context.Context, customer models.Customer) (models.Customer, error) {
	_, err := c.dbPool.Exec(ctx, "UPDATE customers SET name = $1, phone = $2, active = $3 WHERE id = $4",
		customer.Name, customer.Phone, customer.Active, customer.ID)

	return customer, err
}

func (c *CustomerRepository) Delete(ctx context.Context, id int64) error {
	_, err := c.dbPool.Exec(ctx, "DELETE FROM customers WHERE id = $1", id)

	return err
}

func (c *CustomerRepository) GetAllActivated(ctx context.Context) ([]models.Customer, error) {
	var customers []models.Customer

	rows, err := c.dbPool.Query(ctx, "SELECT * FROM customers where active = true")

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var customer models.Customer

		err = rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Phone,
			&customer.Active,
			&customer.Created,
		)

		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (c *CustomerRepository) Activate(ctx context.Context, id int64) error {
	_, err := c.dbPool.Exec(ctx, "UPDATE customers SET active = true WHERE id = $1", id)

	return err
}

func (c *CustomerRepository) Deactivate(ctx context.Context, id int64) error {
	_, err := c.dbPool.Exec(ctx, "UPDATE customers SET active = false WHERE id = $1", id)

	return err
}
