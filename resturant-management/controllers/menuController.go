package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rohitbisht01/resturant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		menuId := c.Param("menu_id")
		if menuId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "menu_id is required"})
			return
		}

		var menu models.Menu
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		if err != mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, menu)
	}
}

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		opts := options.Find().SetSort(bson.D{
			{Key: "created_at", Value: -1},
		})
		cursor, err := menuCollection.Find(ctx, bson.M{}, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(ctx)

		var menus []models.Menu
		if err := cursor.All(ctx, &menus); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, menus)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validateErr := validate.Struct(menu)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}

		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()
		menu.Created_at = time.Now().UTC()
		menu.Updated_at = time.Now().UTC()

		result, insertErr := menuCollection.InsertOne(ctx, menu)
		if insertErr != nil {
			msg := fmt.Sprintf("Menu is not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusCreated, result)
	}
}
