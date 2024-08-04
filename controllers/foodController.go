package controller

import (
	"context"
	"go-restaurant-management/database"
	"go-restaurant-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


var foodCollection *mongo.Collection = database.OpenCollection(database.Client,"food")

func GetFoods() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func GetFood() gin.HandlerFunc{
	return func(c *gin.Context){
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		foodId := c.Param("food_id")

		var food models.Food

		err:=foodCollection.FindOne(ctx, bson.M{"food_id":foodId}).Decode(&food)

		defer cancel()

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while fetching the food items"})
		}
		c.JSON(http.StatusOK, food)

	}
}

func CreateFood() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}

func UpdateFood() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}



func round(num float64)int{}

func toFixed(num float64, precision int) float64{}
