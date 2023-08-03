package s_user

import (
	"mime/multipart"

	// "strings"
	"sitax/model/m_user"
	"sitax/repository/r_user"

	"github.com/gin-gonic/gin"
)

type UpdateUserService struct {
	updateUserRepo r_user.UpdateUserRepository
}

func NewUpdateUserService(updateUserRepo r_user.UpdateUserRepository) *UpdateUserService {
	return &UpdateUserService{updateUserRepo}
}

func (s *UpdateUserService) UpdateUser(ctx *gin.Context, user m_user.User, gambarFile *multipart.FileHeader) (*m_user.User, error) {
	// Mendapatkan ID grup dari parameter rute
	username := ctx.Param("username")

	// Set ID grup ke dalam struct grup
	user.Username = username

	// Memanggil repository untuk update user
	updatedUser, err := s.updateUserRepo.UpdateUser(&user, gambarFile)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
