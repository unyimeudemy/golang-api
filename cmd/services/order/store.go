package order

import (
	"Ecom/types"
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{db: db}
}

// CreateOrder(Order) (int, error)
// CreateOrderItem(OrderItem) error

func (s *Store) CreateOrder(order types.Order)(int, error){
	res, err := s.db.Exec(
		"INSERT INTO orders (userId, total, status, address) VALUES (?,?,?,?)",
		order.UserID, order.Total, order.Status, order.Address,
	)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *Store) CreateOrderItem(orderItems types.OrderItem) error {
	_, err := s.db.Exec(
		"INSERT INTO order_items (orderId, productId, quantity, price) VALUES(?,?,?,?)",
		orderItems.OrderID, orderItems.ProductID, orderItems.Quantity, orderItems.Price,
	)
	return err
}