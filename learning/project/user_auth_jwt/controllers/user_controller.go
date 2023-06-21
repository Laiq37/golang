package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"user-auth-jwt/database"
	helper "user-auth-jwt/helpers"
	"user-auth-jwt/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := helper.CheckUserType(ctx, "ADMIN"); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		fmt.Println("reteriving recordPerPage")
		recordPerPage, err := strconv.Atoi(ctx.Query("recordPerPage"))

		fmt.Println("after record page")

		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		fmt.Println("before page")
		page, err1 := strconv.Atoi(ctx.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}

		fmt.Println("after  page")

		startIndex := (page - 1) * recordPerPage
		// startIndex, err = strconv.Atoi(ctx.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		// groupStage := bson.D{{
		// 	"$group", bson.D{{"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}
		// projectStage := bson.D{{"$project", bson.D{
		// 	{"total_count", 1},
		// 	{
		// 		"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}}}}}
		// removalStage := bson.D{{"$unset", bson.D{{
		// 	"password", 0}}}}
		groupStage := bson.D{{"$group", bson.D{
			{"_id", bson.D{{"_id", "null"}}},
			{"total_count", bson.D{{"$sum", 1}}},
			{"data", bson.D{{"$push", "$$ROOT"}, {"$pull", "$$ROOT.password"}}}}}}

		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{
					// {"$$data.password", 0}
					// {"data.password", 0},
					{"$slice", []interface{}{"$data", startIndex, recordPerPage}},
				}},
			}},
		}

		//aggregating result
		result, err := userCollection.Aggregate(c, mongo.Pipeline{
			matchStage, groupStage, projectStage})
		defer cancel()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		}
		var allUsers []bson.M
		fmt.Println(result)
		if err = result.All(c, &allUsers); err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, allUsers[0])
	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//userid to find
		userId := ctx.Param("id")

		//checking permission if user allowed
		if err := helper.MatchUserTypeToUid(ctx, userId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		//User var
		var user models.User

		//finding desired user from database
		err := userCollection.FindOne(c, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()

		//handling error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//sending user data response
		ctx.JSON(http.StatusOK, user)
	}
}
