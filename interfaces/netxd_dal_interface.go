package interfaces

import "github.com/SELVAKANNAN-P/Netxd_Dal/models"

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.DBResponse, error)
}
