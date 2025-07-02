package product

import (
	"context"

	"github.com/verofreitt/golang_microservice_stack/internal/models"

	"github.com/verofreitt/golang_microservice_stack/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoRepository Product
type MongoRepository interface {
	Create(ctx context.Context, product *models.Product) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) (*models.Product, error)
	GetByID(ctx context.Context, productID primitive.ObjectID) (*models.Product, error)
	Search(ctx context.Context, search string, pagination *utils.Pagination) (*models.ProductsList, error)
}

// RedisRepository Product
type RedisRepository interface {
	SetProduct(ctx context.Context, product *models.Product) error
	GetProductByID(ctx context.Context, productID primitive.ObjectID) (*models.Product, error)
	DeleteProduct(ctx context.Context, productID primitive.ObjectID) error
}
