package controllers

import (
	"context"
	"net/http"
	"yol-arkadasim/database"
	"yol-arkadasim/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUserProfileHandler(c *gin.Context) {
	// Kullanıcı bilgilerini al
	var updateUser models.UpdateableUser
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Kullanıcıya erişimi doğrula
	userID := c.GetString("userID") // Örnek: Kullanıcı kimliğini JWT token'dan al

	// Kullanıcıyı veritabanından bul
	existingUser, err := findUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Güncellenebilir alanları güncelle
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

	// Kullanıcıyı kaydet
	err = existingUser.SaveToMongoDB(database.GetMongoClient(), "mydatabase", "users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Başarı durumunda, kullanıcıya yanıt gönderin veya başka bir işlem yapın
	c.JSON(http.StatusOK, gin.H{"message": "Profil bilgileri başarıyla güncellendi"})
}

func findUserByID(userID string) (*models.User, error) {
	// MongoDB bağlantısını al
	client := database.GetMongoClient()

	// Kullanıcıyı bulmak için filtre oluştur
	filter := bson.M{"_id": userID}

	// Kullanıcıyı veritabanından bul
	var user models.User
	err := client.Database("mydatabase").Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Kullanıcı bulunamadı
			return nil, nil
		}
		// Diğer hatalar için
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
