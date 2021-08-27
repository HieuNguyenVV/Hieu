package requests

type CreateCustomerRequest struct {
	Id           int64  `json:"id"`
	CustomerName string `json:"name" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Address      string `json:"address" binding:"max=256,min=6"`
	License      string `json:"license" binding:"required,min=6,max=256"`
	Active       bool   `json:"active"`
	Phone        string `json:"phone" binding:"max=256,min=6"`
	Email        string `json:"email"binding:"max=256,min=6"`
}
type CreateCustomerRequestChangPassword struct {
	Id               int64  `json:"id"`
	Old_password     string `json:"old_password" binding:"required"`
	New_password     string `json:"new_password" binding:"required"`
	Confirm_password string `json:"confirm_password" binding:"required"`
}
