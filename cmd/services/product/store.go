package product

import (
	"Ecom/types"
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts()([]*types.Product, error){
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	// initialize an extendable array
	products := make([]*types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (s *Store) CreateProduct(product types.CreateProductPayload)error{
	_, err := s.db.Exec(
		"INSERT INTO products (name, description, image, price, quantity) VALUES(?,?,?,?,?)",
		product.Name, product.Description, product.Image, product.Price, product.Quantity,
	)

	if err != nil{
		return err
	}
	return nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}