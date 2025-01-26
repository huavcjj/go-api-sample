package repository

import (
	"fmt"
	"go-api-sample/domain/entity"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(user *entity.User) (*entity.User, error)
		FindByID(id string) (*entity.User, error)
		FindByEmail(email string) (*entity.User, error)
		FindAll() ([]entity.User, error)
		Save(user *entity.User) (*entity.User, error)
		Delete(id string) error
	}
)
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetOrCreateCategory(user *entity.User) error {
	var category entity.Category
	if err := r.db.FirstOrCreate(&category, entity.Category{Name: user.Category.Name}).Error; err != nil {
		return err
	}
	user.CategoryID = category.ID
	user.Category = category
	return nil
}
func (r *userRepository) Create(user *entity.User) (*entity.User, error) {
	if err := r.GetOrCreateCategory(user); err != nil {
		return nil, err
	}
	if _, err := r.FindByEmail(user.Email); err == nil {
		return nil, err
	}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByID(id string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Preload("Category").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Preload("Category").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Save(user *entity.User) (*entity.User, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	selectedUser, err := r.FindByID(user.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := r.GetOrCreateCategory(user); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := copier.CopyWithOption(selectedUser, user, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to copy album fields: %w", err)
	}

	if err := tx.Save(selectedUser).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return selectedUser, nil
}

func (r *userRepository) Delete(id string) error {
	user, err := r.FindByID(id)
	if err != nil {
		return err
	}
	if err := r.db.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
