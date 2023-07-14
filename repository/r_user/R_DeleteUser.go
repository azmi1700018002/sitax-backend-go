package r_user

import (
	"sitax/config/db"
	"sitax/model/m_user"
)

type DeleteUserRepository interface {
	DeleteUserByID(username string) error
}

type userDeleteRepository struct{}

func NewDeleteUserRepository() DeleteUserRepository {
	return &userDeleteRepository{}
}

func (*userDeleteRepository) DeleteUserByID(username string) (err error) {
	// Menghapus data user dari tabel User
	if err = db.Server().Unscoped().Where("username = ?", username).Delete(&m_user.User{}).Error; err != nil {
		return err
	}

	return nil
}
