package interfaces

import "Netxd_Dal/models"

type ICustomer interface {
	CreateCustomer(*models.Customer) (*models.DBResponse, error)
}
