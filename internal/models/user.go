package models

type User struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	Contact     int               `json:"contact"`
	PaymentCard string            `json:"paymentCard"`
	Role        map[string]string `json:"role"`
}
