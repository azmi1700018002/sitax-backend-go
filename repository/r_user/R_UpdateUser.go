package r_user

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sitax/config/db"
	"sitax/config/helper"
	"sitax/model/m_user"
	"time"

	"github.com/google/uuid"
)

type UpdateUserRepository interface {
	UpdateUser(user *m_user.User, gambarFile *multipart.FileHeader) (*m_user.User, error)
}

type updateUserRepository struct{}

func NewUpdateUserRepository() UpdateUserRepository {
	return &updateUserRepository{}
}

func (r *updateUserRepository) UpdateUser(user *m_user.User, gambarFile *multipart.FileHeader) (*m_user.User, error) {
	// Check if the user exists
	var existingUser m_user.User
	if err := db.Server().Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		return nil, err
	}

	// If password is empty, use the existing password value from the database
	if user.Password == "" {
		user.Password = existingUser.Password
	} else {
		// If password is not empty, hash the new password before update
		hashedPassword, err := helper.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	//Path for update profile
	basePath := "../../../../../../../laragon/www/Sitax/assets/img/profile/"

	fmt.Println("Deleting old image file:", filepath.Join(basePath, existingUser.ProfilePicture))

	// Jika ada file gambar baru yang dimasukkan, hapus file lama dari local storage
	if gambarFile != nil {
		if existingUser.ProfilePicture != "" {
			err := os.Remove(filepath.Join(basePath, existingUser.ProfilePicture))
			if err != nil {
				return nil, err
			}
		}
	} else {
		// Jika tidak ada file gambar baru yang dimasukkan, tetapkan kembali nama gambar sebelumnya
		user.ProfilePicture = existingUser.ProfilePicture
	}

	// Jika ada file gambar baru yang dimasukkan, proses gambar baru
	if gambarFile != nil {
		fileExt := filepath.Ext(gambarFile.Filename)
		newFileName := uuid.New().String() + fileExt

		// Open the source file
		src, err := gambarFile.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		// Create the destination file
		savePath := filepath.Join(basePath, newFileName)

		dst, err := os.Create(savePath)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		// Copy the contents from source to destination
		_, err = io.Copy(dst, src)
		if err != nil {
			return nil, err
		}

		// Update the product's image path
		user.ProfilePicture = newFileName
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
			"sts_user":        user.StsUser,
			"KdKantor":        user.KdKantor,
			"profile_picture": user.ProfilePicture,
			"created_at":      user.CreatedAt,
			"updated_at":      user.UpdatedAt, // update the updated_at value
		}).Error; err != nil {
		return nil, err
	}

	return user, nil
}
