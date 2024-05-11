package member

import (
	"biomassx/models"

	"gorm.io/gorm"
)

// Service is the interface that provides methods for creating, reading, updating, deleting, listing, and searching members.
type Service interface {
	Create(member *models.Member) error
	Read(id uint) (*models.Member, error)
	Update(member *models.Member) error
	Delete(id uint) error
	List() ([]*models.Member, error)
	Search(name string) ([]*models.Member, error)
}

// ServiceImpl is a struct that implements the Service interface using a gorm.DB.
type ServiceImpl struct {
	db *gorm.DB
}

// NewService is a function that creates a new instance of the ServiceImpl struct.
func NewService(db *gorm.DB) *ServiceImpl {
	return &ServiceImpl{db: db}
}

// Create is a method that adds a new member to the database.
func (s *ServiceImpl) Create(member *models.Member) error {
	return s.db.Create(member).Error
}

// Read is a method that retrieves a member from the database by its ID.
func (s *ServiceImpl) Read(id uint) (*models.Member, error) {
	var member models.Member
	if err := s.db.First(&member, id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

// Update is a method that updates a member's details in the database.
func (s *ServiceImpl) Update(member *models.Member) error {
	return s.db.Save(member).Error
}

// Delete is a method that removes a member from the database by its ID.
func (s *ServiceImpl) Delete(id uint) error {
	return s.db.Delete(&models.Member{}, id).Error
}

// List is a method that retrieves all members from the database.
func (s *ServiceImpl) List() ([]*models.Member, error) {
	var members []*models.Member
	if err := s.db.Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

// Search is a method that finds members in the database by their name.
func (s *ServiceImpl) Search(name string) ([]*models.Member, error) {
	var members []*models.Member
	if err := s.db.Where("name LIKE ?", "%"+name+"%").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}
