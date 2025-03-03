package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/Teaboyds/Todoist_By_Teaboyd/models"
	"github.com/Teaboyds/Todoist_By_Teaboyd/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		user := new(models.User)                   //สร้าง instance ของ User
		if err := c.BodyParser(user); err != nil { //รับ body จาก user
			return err
		}

		hashedPassword, err := utils.HashedPassword(user.Password) //hashed password

		if err != nil {
			response := map[string]string{"error": "hash problems"}
			fmt.Println(response)
		}

		user.Password = hashedPassword

		if err := db.Create(user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not create user",
			})
		} //build in user ใหม่

		return c.JSON(user)
	}
}

func Login(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {

		Inputuser := new(models.User)

		if err := c.BodyParser(Inputuser); err != nil {
			return c.SendStatus(fiber.ErrBadRequest.Code)
		}

		databaseUser := new(models.User)
		result := db.Where("username = ?", Inputuser.Username).First(databaseUser)
		if result.Error != nil {
			return result.Error
		}

		err := bcrypt.CompareHashAndPassword([]byte(databaseUser.Password), []byte(Inputuser.Password))

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "Password Invalid"})
		}

		// build to ken login
		jwtSecretKey := os.Getenv("SECRET")

		token := jwt.New(jwt.SigningMethodHS256)
		Claims := token.Claims.(jwt.MapClaims)
		Claims["user_id"] = uint(databaseUser.ID)
		Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(jwtSecretKey))

		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"message": "login successfully",
			"token":   t,
		})
	}
}
