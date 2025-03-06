package repositories

import "github.com/viniciusith/expense-tracker/internal/domain/entities"

type UserRepository interface {
    Create(user *entities.User) error
    Update(user *entities.User) error
    Delete(user *entities.User) error
    FindById(id string) (*entities.User, error)
    FindByEmail(email string) (*entities.User, error)
}
