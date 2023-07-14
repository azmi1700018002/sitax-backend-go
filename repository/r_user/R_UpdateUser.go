package r_user

import (
	"sitax/config/db"
	"sitax/config/helper"
	"sitax/model/m_user"
	"time"
)

type UpdateUserRepository interface {
	UpdateUser(user *m_user.User) (*m_user.User, error)
}

type updateUserRepository struct{}

func NewUpdateUserRepository() UpdateUserRepository {
	return &updateUserRepository{}
}

func (r *updateUserRepository) UpdateUser(user *m_user.User) (*m_user.User, error) {
	// Hash password before update
	if user.Password != "" {
		hashedPassword, err := helper.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	// Check if the user exists
	var existingUser m_user.User
	if err := db.Server().Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	user.CreatedAt = existingUser.CreatedAt // keep the existing created_at value
	user.UpdatedAt = time.Now()             // update the updated_at value
	if err := db.Server().Model(&m_user.User{}).
		Where("username = ?", user.Username).
		Updates(map[string]interface{}{
			"username":        user.Username,
			"password":        user.Password,
			"nama_lengkap":    user.NamaLengkap,
			"email":           user.Email,
			"host_ip":         user.HostIP,
			"KdKantor":        user.KdKantor,
			"profile_picture": user.ProfilePicture,
			"created_at":      user.CreatedAt,
			"updated_at":      user.UpdatedAt, // update the updated_at value
		}).Error; err != nil {
		return nil, err
	}

	return user, nil
}
