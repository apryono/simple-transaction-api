package models

import "simple-transaction-api/usecase/viewmodel"

// Transaction ...
type Transaction struct {
	ID                 int                  `json:"id"`
	TransactionNumber  string               `json:"transaction_number"`
	CustomerID         int                  `json:"customer_id"`
	Customer           viewmodel.CustomerVM `json:"customer"`
	PicName            string               `json:"pic_name"`
	PicPhone           string               `json:"pic_phone"`
	PicEmail           string               `json:"pic_email"`
	TransactionDetails []TransactionDetail  `json:"transaction_details"`
	TotalPrice         float64              `json:"total_price"`
	TypeOfPayment      string               `json:"type_of_payment"`
	Note               string               `json:"note"`
	Status             string               `json:"status"`
	CreatedAt          string               `json:"created_at"`
	UpdatedAt          string               `json:"updated_at"`
	DeletedAt          *string              `json:"deleted_at"`
}

// TransactionParameter ...
type TransactionParameter struct {
	ID                 int                 `json:"id"`
	TransactionNumber  string              `json:"transaction_number"`
	CustomerID         int                 `json:"customer_id"`
	Customer           Customer            `json:"customer"`
	PicName            string              `json:"pic_name"`
	PicPhone           string              `json:"pic_phone"`
	PicEmail           string              `json:"pic_email"`
	TransactionDetails []TransactionDetail `json:"transaction_details"`
	TotalPrice         float64             `json:"total_price"`
	TypeOfPayment      string              `json:"type_of_payment"`
	Note               string              `json:"note"`
	Status             string              `json:"status"`
	Search             string              `json:"search"`
}

var (
	StatusPending  = "pending"
	StatusPaid     = "paid"
	StatusCanceled = "canceled"

	TransactionSelectStatement = `SELECT def.id, def.transaction_number, def.customer_id, jsonb_build_object(
			'id', cs.id,
			'name', cs.name
		), 
		def.pic_name, def.pic_phone, def.pic_email, 
		def.total_price, def.type_of_payment, jsonb_agg( distinct 
			jsonb_build_object(
				'id', td.id,
				'transaction_id', td.transaction_id,
				'product_id', td.product_id,
				'product_name', td.product_name,
				'product_price', td.product_price,
				'product_quantity', td.product_quantity
			) 
		) "transaction_details", 
		def.note, def.status, def.created_at, def.updated_at  
	FROM transactions def 
	LEFT JOIN customers cs ON cs.id = def.customer_id 
	LEFT JOIN transaction_details td ON def.id = td.transaction_id
	`
	TransactionWhereStatement = `WHERE def.deleted_at IS NULL `

	TransactionGroupByStatement = `GROUP BY def.id, cs.id`
)
