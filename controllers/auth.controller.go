package controllers

import (
	"context"
	"fiber/configs"
	"fiber/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var userCollection = configs.GetCollection(configs.DB, "users")

func hashPassword(password string) string {
	rounds, _ := strconv.Atoi(os.Getenv("ROUNDS"))
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), rounds)
	return string(hash)
}

func comparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
func Register(c *fiber.Ctx) error {
	newUser := new(models.User)
	newUser.Id = primitive.NewObjectID()
	if err := c.BodyParser(newUser); err != nil {
		log.Fatal(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	hashedPassword := hashPassword(newUser.Password)
	newUser.Password = hashedPassword
	res, err := userCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID
	fmt.Println(res)
	return c.JSON((&fiber.Map{
		"status": "success",
		"id":     id,
	}))
}

func Login(c *fiber.Ctx) error {
	user := new(models.Login)
	if err := c.BodyParser(user); err != nil {
		log.Fatal(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	var u models.User
	userCollection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&u)
	if u.Username == "" {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "Username not found",
		})
	} else {
		e := comparePassword(u.Password, user.Password)
		if e != nil {
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": "Password is not valid!",
			})
		} else {
			userId := string(u.Id.Hex())
			claims := &jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				ID:        userId,
			}
			mysecret := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, _ := token.SignedString(mysecret)
			var result models.LoginSuccess
			result.Token = tokenString
			result.Id = u.Id
			result.Username = u.Username
			return c.Status(200).JSON(&fiber.Map{
				"status": true,
				"data":   result,
			})
		}
	}

}
