package models

import "errors"

// OrderRequest represents the incoming JSON payload for an order
type OrderRequest struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
}

// Validate checks if the request data is valid
func (r *OrderRequest) Validate() error {
	if r.UserID <= 0 {
		return errors.New("user_id must be greater than 0")
	}
	if r.ProductID <= 0 {
		return errors.New("product_id must be greater than 0")
	}
	return nil
}
