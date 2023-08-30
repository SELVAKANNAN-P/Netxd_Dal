package service

import (
	"Netxd_Dal/interfaces"
	"Netxd_Dal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}
func (c *CustomerService) CreateCustomer(Customer *models.Customer) (*models.DBResponse, error) {
	// Customer.CustomerID = 123
	// Customer.FirstName = "selva"
	// Customer.LastName = "kannan"
	// Customer.BankID = 1234
	// Customer.Balance = 89000
	// Customer.CreatedAt = "29.08.2023"
	// Customer.UpdateAt = "30.08.2023"
	// Customer.ISActive = true
	// res, err := c.CustomerCollection.InsertOne(c.ctx, &Customer)
	res, err := c.CustomerCollection.InsertOne(c.ctx, &Customer)

	if err != nil {
		return nil, err
	}
	Response := &models.DBResponse{
		CustomerID: res.InsertedID.(primitive.ObjectID),
		CreatedAt:  Customer.CreatedAt,
	}
	return Response, nil
}
