package item

import (
	"task-api/internal/constant"
	"task-api/internal/model"

	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
}

func NewService(db *gorm.DB) Service {
	return Service{
		Repository: NewRepository(db),
	}
}

func (service Service) Create(req model.RequestItem, userID uint) (model.Item, error) {
	item := model.Item{
		Title:    req.Title,
		Amount:   req.Amount,
		Quantity: req.Quantity,
		Status:   constant.ItemPendingStatus,
		OwnerID:  userID, // Set OwnerID from the passed userID
	}

	if err := service.Repository.Create(&item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (service Service) Find(query model.RequestFindItem) ([]model.Item, error) {
	return service.Repository.Find(query)
}
func (service Service) FindByID(id uint, query model.RequestFindItem) (model.Item, error) {
	return service.Repository.FindByID(id)
}
func (service Service) UpdateStatus(id uint, status constant.ItemStatus) (model.Item, error) {
	// Find item
	item, err := service.Repository.FindByID(id)
	if err != nil {
		return model.Item{}, err
	}

	// Fill data
	item.Status = status

	// Replace
	if err := service.Repository.Replace(item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}
func (service Service) UpdateItem(id uint, req model.RequestItem) (model.Item, error) {
	// Find item
	item, err := service.Repository.FindByID(id)
	if err != nil {
		return model.Item{}, err
	}
	// Fill data
	item.Title = req.Title
	item.Amount = req.Amount
	item.Quantity = req.Quantity
	item.Status = constant.ItemPendingStatus

	// Replace
	if err := service.Repository.Replace(item); err != nil {
		return model.Item{}, err
	}
	return item, nil

}
func (service Service) DeleteItem(id uint) (model.Item, error) {
	// Find item by ID
	_, err := service.Repository.FindByID(id)
	if err != nil {
		return model.Item{}, err // Return if not found or any other error occurs
	}

	// Use the Delete method from the repository to delete the item
	deletedItem, err := service.Repository.Delete(id)
	if err != nil {
		return model.Item{}, err
	}

	return deletedItem, nil
}
