package helpers

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	passInBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(passInBytes)
}

func VerifyPassword(enteredPass string, userPass string) (bool, string) {
	fmt.Println(enteredPass)
	fmt.Println(userPass)
	err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(enteredPass))
	if err != nil {
		msg := fmt.Sprintf("email or password")
		return false, msg
	}
	return true, ""
}

func CheckUserType(ctx *gin.Context, role string) (err error) {
	userType := ctx.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(ctx *gin.Context, userId string) (err error) {
	userType := ctx.GetString("user_type")
	uid := ctx.GetString("id")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(ctx, userType)
	return err
}
