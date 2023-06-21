package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	helper "user-auth-jwt/helpers"
	"user-auth-jwt/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		//binding request body with user interface/struct and checking error
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			cancel()
			return
		}
		fmt.Println(*user.User_type)
		//validating user data acc to struct model which we have defined in struct model
		validationErr := validate.Struct(user)
		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			cancel()
			return
		}
		fmt.Println("getting email cound")
		//getting email count which matches with request body email
		count, err := userCollection.CountDocuments(c, bson.M{"email": user.Email})
		fmt.Println(err)
		defer cancel()
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking email"})
			return
		}

		fmt.Println("getting phone cound")
		*user.Password = helper.HashPassword(*user.Password)
		//getting email count which matches with request body phone
		count, err = userCollection.CountDocuments(c, bson.M{"phone": user.Phone})
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking phone number"})
			cancel()
			return
		}

		//if count > 0 it means phone or email has already exist
		if count > 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email or phone number already exist!"})
			cancel()
			return
		}
		fmt.Println("putting data in model")
		//putting data in remaining field
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Id = primitive.NewObjectID()
		user.User_id = user.Id.Hex()

		//getting token and refresh token
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id, *user.User_type)
		user.Token = &token
		user.Refresh_token = &refreshToken

		//inserting user data in db
		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("user not created")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		ctx.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			cancel()
			return
		}
		err := userCollection.FindOne(c, bson.M{"email": *user.Email}).Decode(&foundUser)
		fmt.Println(foundUser)
		defer cancel()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}
		if foundUser.Email == nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			cancel()
			return
		}
		isPasswordValid, msg := helper.VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if isPasswordValid != true {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			cancel()
			return
		}
		token, refreshToken, err := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, foundUser.User_id, *foundUser.User_type)
		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		err = userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			cancel()
			return
		}
		ctx.JSON(http.StatusOK, foundUser)
	}
}
