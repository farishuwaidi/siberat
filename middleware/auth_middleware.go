package middleware

import (
	"Siberat/config"
	"Siberat/entity"
	errorhandler "Siberat/errorHandler"
	"Siberat/helper"

	"github.com/gin-gonic/gin"
)

func JWTMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			errorhandler.HandlerError(c, &errorhandler.UnAuthorizedError{Message: "Unauthorized"})
			c.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenString)
		if err != nil {
			errorhandler.HandlerError(c,&errorhandler.UnAuthorizedError{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", userId)
		c.Next()
	}
}

func AuthorizeRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			errorhandler.HandlerError(c, &errorhandler.UnAuthorizedError{Message: "Unathorized"})
			c.Abort()
			return
		}

		// ambil user lengkap dari DB termasuk rolenya
		var user entity.User
		if err := config.DB.Preload("Role").First(&user, userID).Error; err != nil {
			errorhandler.HandlerError(c, &errorhandler.UnAuthorizedError{Message: "User Not Found"})
			c.Abort()
			return
		}

		// cek apakah role user ada di daftar allowedroles
		isAllowed := false
		for _, role := range allowedRoles{
			if user.Role.Role == role {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			errorhandler.HandlerError(c, &errorhandler.UnAuthorizedError{Message: "Access denied"})
			c.Abort()
			return
		}

		// set user ke context agar bisa di pakai
		c.Set("user", user)
		c.Next()
	}
}
