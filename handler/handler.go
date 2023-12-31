package handler

import (
	"sitax/config/helper"
	"sitax/model/m_user"
	"sitax/repository/r_auth"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var identityKey = "id"

func LoginHandler(c *gin.Context) {
	var loginVals m_user.User
	if err := c.ShouldBind(&loginVals); err != nil {
		c.JSON(400, gin.H{"error": "missing login values"})
		return
	}
	email := loginVals.Email
	password := loginVals.Password

	// Check if user exist
	user, err := r_auth.QAuthUser(email, password)

	// Check if username and password match using CheckPasswordHash
	if !helper.CheckPasswordHash(password, user.Password) {
		c.JSON(401, gin.H{"error": "failed authentication"})
		return
	}

	if err != nil {
		c.JSON(401, gin.H{"error": "failed authentication"})
		return
	}

	// Update last login time if authentication succeeds
	if err := r_auth.UpdateLastLogin(&user); err != nil {
		c.JSON(401, gin.H{"error": "failed authentication"})
		return
	}

	expireTime := time.Now().Add(1 * time.Hour)

	// konversi Unix ke Time
	expireTimeUTC := time.Unix(expireTime.Unix(), 0).UTC()

	// format waktu sebagai string
	expireTimeString := expireTimeUTC.Format(time.RFC3339)

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		identityKey:      user.Username,
		"Username":       user.Username,
		"ProfilePicture": user.ProfilePicture,
		"exp":            expireTime.Unix(),
	}).SignedString([]byte("secret key"))

	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token":          token,
		"Username":       user.Username,
		"StsUser":        user.StsUser,
		"GroupID":        user.GroupID,
		"ProfilePicture": user.ProfilePicture,
		"expired":        expireTimeString,
		"status":         200,
	})
}
