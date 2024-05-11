package member

import (
	"biomassx/models"

	"gorm.io/gorm"
)

// gormRepository struct provides the implementation of the Repository interface
// using gorm as the underlying database ORM.
type gormRepository struct {
	db *gorm.DB
}

// Repository interface provides access to a member store.
type Repository interface {
	Create(m *models.Member) error
	Read(id uint) (*models.Member, error)
	Update(m *models.Member) error
	Delete(id uint) error
	List() ([]*models.Member, error)
	Search(name string) ([]*models.Member, error)
	ReadByEmail(email string) (*models.Member, error)
}

// NewgormRepository is a factory function that returns a new instance of gormRepository.
func NewgormRepository(db *gorm.DB) *gormRepository {
	return &gormRepository{db: db}
}

// Create inserts a new member into the database.
func (r *gormRepository) Create(m *models.Member) error {
	return r.db.Create(m).Error
}

// Read retrieves a member from the database by its ID.
func (r *gormRepository) Read(id uint) (*models.Member, error) {
	var m models.Member
	if err := r.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

// Update modifies an existing member in the database.
func (r *gormRepository) Update(m *models.Member) error {
	return r.db.Save(m).Error
}

// Delete removes a member from the database by its ID.
func (r *gormRepository) Delete(id uint) error {
	return r.db.Delete(&models.Member{}, id).Error
}

// List retrieves all members from the database.
func (r *gormRepository) List() ([]*models.Member, error) {
	var members []*models.Member
	if err := r.db.Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

// Search finds members by name in the database.
func (r *gormRepository) Search(name string) ([]*models.Member, error) {
	var members []*models.Member
	if err := r.db.Where("name LIKE ?", "%"+name+"%").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

// ReadByEmail retrieves a member from the database by its email.
func (r *gormRepository) ReadByEmail(email string) (*models.Member, error) {
	var m models.Member
	if err := r.db.Where("email = ?", email).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
