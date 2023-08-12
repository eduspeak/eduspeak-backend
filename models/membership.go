package models

import ( 
	"time"
	"gorm.io/gorm"
)

type Membership struct{
	gorm.Model
	Id int `json:"id,omitempty,autoIncrement"`
	UserId int `json:"user_id,omitempty" validate:"required"`
	PaymentDate time.Time `json:"payment_date" validate:"required"`
	Status int `json:"status" validate:"required"`
	ProofOfPayment string `json:"proof_of_payment" validate:"required"`
}