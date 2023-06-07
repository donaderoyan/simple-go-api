package createProduct

import "github.com/shopspring/decimal"

type InputCreateProduct struct {
	UserID           string          `json:"user_id" validate:"required"`
	ParentID         string          `json:"parent_id" validate:""`
	ProductImages    string          `json:"productImages" validate:""`
	Categories       string          `json:"categories" validate:""`
	Sku              string          `json:"sku" validate:"required"`
	Name             string          `json:"name" validate:"required"`
	Slug             string          `json:"slug" validate:"required"`
	Price            decimal.Decimal `json:"price" validate:"required"`
	Stock            int             `json:"stock" validate:""`
	Weight           decimal.Decimal `json:"weight" validate:""`
	ShortDescription string          `json:"shortDescription" validate:""`
	Description      string          `json:"description" validate:""`
	Status           int             `json:"status"`
}
