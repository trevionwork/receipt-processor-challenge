package repositories

import (
	"errors"
	"sync"

	"github.com/fetch-rewards/receipt-processor-challenge/pkg/models"
	"github.com/google/uuid"
)

var ErrNotFound = errors.New("not found")

type ReceiptPointsRepository interface {
	// SaveReceiptPoints save points for a receipt and return an ID for this receipt and a possible error if operation failed
	SaveReceiptPoints(receipt *models.Receipt, points int) (id string, err error)
	// GetReceiptPointsById get points that computed for a receipt with specified ID
	GetReceiptPointsById(id string) (int, error)
}

type inMemoryReceiptPointsRepository struct {
	data sync.Map
}

func NewInMemoryReceiptPointsRepository() ReceiptPointsRepository {
	return &inMemoryReceiptPointsRepository{}
}

func (repo *inMemoryReceiptPointsRepository) SaveReceiptPoints(receipt *models.Receipt, points int) (id string, err error) {
	id = uuid.NewString()
	repo.data.Store(id, points)
	// err is nil
	return
}

func (repo *inMemoryReceiptPointsRepository) GetReceiptPointsById(id string) (int, error) {
	value, exists := repo.data.Load(id)
	if !exists {
		return 0, ErrNotFound
	}

	return value.(int), nil
}
