package controllers

import (
	"context"
	"net/http"
	"yol-arkadasim/database"
	"yol-arkadasim/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUserProfileHandler(c *gin.Context) {
	var updateUser models.UpdateableUser
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userID := c.GetString("userID")

	existingUser, err := findUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if existingUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if updateUser.Name != nil {
		existingUser.Name = updateUser.Name
	}
	if updateUser.Surname != nil {
		existingUser.Surname = updateUser.Surname
	}
	if updateUser.Username != nil {
		existingUser.Username = updateUser.Username
	}
	if updateUser.Password != nil {
		existingUser.Password = updateUser.Password
	}
	if updateUser.DateOfBirth != nil {
		existingUser.DateOfBirth = updateUser.DateOfBirth
	}
	if updateUser.Phone != nil {
		existingUser.Phone = updateUser.Phone
		//konum bilgisi ekle -- biyografi eklenecek--ilgi alanları ekle--
	}

	err = existingUser.SaveToMongoDB(database.GetMongoClient(), "mydatabase", "users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func findUserByID(userID string) (*models.User, error) {
	client := database.GetMongoClient()
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	var user models.User
	err = client.Database("mydatabase").Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func GetAllUsersHandler(c *gin.Context) {
	// Veritabanı bağlantısını al
	client := database.GetMongoClient()

	// Collection belirle
	collection := client.Database("mydatabase").Collection("users")

	// Tüm kullanıcıları bul
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer cursor.Close(context.Background())

	// Kullanıcıları bir dilimde depolamak için boş bir dilim oluştur
	var users []models.User

	// Tüm kullanıcıları döngü ile al
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		users = append(users, user)
	}

	// Kullanıcıları başarıyla aldıktan sonra, JSON olarak yanıt ver
	c.JSON(http.StatusOK, gin.H{"users": users})
}
