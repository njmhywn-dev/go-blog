package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/njmhywn-dev/go-blog/database"
	"github.com/njmhywn-dev/go-blog/helper"
	"github.com/njmhywn-dev/go-blog/model"
	"golang.org/x/crypto/bcrypt"
)

type formData struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "",
		"msg":        "Something went wrong.",
	}

	var formData formData

	if err := c.BodyParser(&formData); err != nil {
		log.Println("Form binding error.")

		c.Status(400)
		return c.JSON(context)
	}

	var user model.User

	database.DBConn.First(&user, "user_id=?", formData.UserId)

	if user.ID == 0 {
		context["msg"] = "User not found."

		c.Status(400)
		return c.JSON(context)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))

	if err != nil {
		log.Println("Invalid password.")

		context["msg"] = "Password doesnt match"
		c.Status(401)
		return c.JSON(context)
	}

	token, err := helper.GenerateToken(user)

	if err != nil {
		context["msg"] = "Token creation error."
		c.Status(401)
		return c.JSON(context)
	}

	context["token"] = token
	context["user"] = user
	context["statusText"] = "Ok"
	context["msg"] = "User authenticated."

	c.Status(200)
	return c.JSON(context)
}

func Register(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Register route.",
	}

	var formData formData

	if err := c.BodyParser(&formData); err != nil {
		log.Println("Error in json binding.")
		context["statusText"] = ""
		context["msg"] = "Error occurs."

		c.Status(400)
		return c.JSON(context)
	}

	var user model.User

	user.UserId = formData.UserId
	user.Password = helper.HashPassword(formData.Password)

	result := database.DBConn.Create(&user)

	if result.Error != nil {
		log.Println(result.Error)
		context["msg"] = "User already exists."
		c.Status(400)
		return c.JSON(context)
	}

	context["msg"] = "User added successfully."
	c.Status(201)
	return c.JSON(context)
}

func Loguot() {}

func RefreshToken(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Refresh Token  route",
	}

	user_id := c.Locals("user_id")

	if user_id == nil {
		context["msg"] = "User not found."

		c.Status(400)
		return c.JSON(context)
	}

	var user model.User

	database.DBConn.First(&user, "user_id=?", user_id)

	if user.ID == 0 {
		context["msg"] = "User not found."

		c.Status(400)
		return c.JSON(context)
	}

	token, err := helper.GenerateToken(user)

	if err != nil {
		context["msg"] = "Token creation error."

		c.Status(401)
		return c.JSON(context)
	}

	context["token"] = token
	context["user"] = user

	c.Status(200)
	return c.JSON(context)
}

func Profile(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "User profile route. This is protected route accessible to authenticated user only.",
	}

	c.Status(200)
	return c.JSON(context)
}
