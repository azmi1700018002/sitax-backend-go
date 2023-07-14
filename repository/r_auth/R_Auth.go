package r_auth

import (
	"sitax/config/db"
	"sitax/model/m_user"
	"time"
)

func QAuthUser(email, password string) (user m_user.User, err error) {

	err = db.Server().Where("email = ?", email).First(&user).Error

	return user, err

}

func UpdateLastLogin(user *m_user.User) error {
	now := time.Now()

	user.LastLogin = &now
	return db.Server().Save(user).Error
}
