package repo

import (
	"context"
	"mycrudapp/internal/db"
	"mycrudapp/internal/models"
)

type CustomerRepository struct{}

func (c *CustomerRepository) GetById(ctx context.Context, id int64) (models.Customer, error) {
	var customer models.Customer

	err := db.DB.QueryRowContext(ctx, "SELECT * FROM customers WHERE id = $1", id).Scan(
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

	rows, err := db.DB.QueryContext(ctx, "SELECT * FROM customers")

	defer func() {
		if cerr := rows.Close(); cerr != nil {
			err = cerr
		}
	}()

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
	err := db.DB.QueryRowContext(ctx, "INSERT INTO customers (name, phone, active) VALUES ($1, $2, $3) RETURNING id, created",
		customer.Name, customer.Phone, customer.Active).Scan(&customer.ID, &customer.Created)

	return customer, err
}

func (c *CustomerRepository) Update(ctx context.Context, customer models.Customer) (models.Customer, error) {
	_, err := db.DB.ExecContext(ctx, "UPDATE customers SET name = $1, phone = $2, active = $3 WHERE id = $4",
		customer.Name, customer.Phone, customer.Active, customer.ID)

	return customer, err
}

func (c *CustomerRepository) Delete(ctx context.Context, id int64) error {
	_, err := db.DB.ExecContext(ctx, "DELETE FROM customers WHERE id = $1", id)

	return err
}

func (c *CustomerRepository) GetAllActivated(ctx context.Context) ([]models.Customer, error) {
	var customers []models.Customer

	rows, err := db.DB.QueryContext(ctx, "SELECT * FROM customers where active = true")

	defer func() {
		if cerr := rows.Close(); cerr != nil {
			err = cerr
		}
	}()

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
	_, err := db.DB.ExecContext(ctx, "UPDATE customers SET active = true WHERE id = $1", id)

	return err
}

func (c *CustomerRepository) Deactivate(ctx context.Context, id int64) error {
	_, err := db.DB.ExecContext(ctx, "UPDATE customers SET active = false WHERE id = $1", id)

	return err
}
