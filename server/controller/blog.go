package controller

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/njmhywn-dev/go-blog/database"
	"github.com/njmhywn-dev/go-blog/model"
)

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog List",
	}

	time.Sleep(time.Millisecond * 1000)

	db := database.DBConn

	var records []model.Blog

	db.Find(&records)

	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
}

func BlogDetail(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."

		c.Status(404)
		return c.JSON(context)
	}

	context["record"] = record
	context["statusText"] = "Ok"
	context["msg"] = "Record detail."

	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a Blog",
	}

	record := new(model.Blog)

	if err := c.BodyParser(record); err != nil {
		log.Println("Error in parsing request:", err)
		context["statusText"] = ""
		context["msg"] = "Invalid input data."
		c.Status(400)
		return c.JSON(context)
	}

	// File upload
	file, err := c.FormFile("file")

	if err != nil && err != fiber.ErrUnprocessableEntity {
		log.Println("Error in file upload.", err)
		context["msg"] = "There was an error with file upload."
		c.Status(400)
		return c.JSON(context)
	}

	if file != nil && file.Size > 0 {
		filename := "./static/uploads/" + file.Filename

		if err := c.SaveFile(file, filename); err != nil {
			log.Println("Error in file uploading...", err)
			context["msg"] = "Error in file uploading."
			c.Status(500)
			return c.JSON(context)
		}

		// Set image path to the struct
		record.Image = filename
	}

	result := database.DBConn.Create(&record)

	if result.Error != nil {
		log.Println("Error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
		c.Status(500)
		return c.JSON(context)
	}

	context["msg"] = "Record is saved successfully."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update a Blog",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found.")

		context["statusText"] = ""
		context["msg"] = "Record not Found."

		c.Status(400)
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")

		context["msg"] = "Something went wrong."
		c.Status(400)
		return c.JSON(context)
	}

	// File upload
	file, err := c.FormFile("file")

	if err != nil && err != fiber.ErrUnprocessableEntity {
		log.Println("Error in file upload.", err)

	}

	if file != nil && file.Size > 0 {
		filename := "static/uploads/" + file.Filename

		if err := c.SaveFile(file, filename); err != nil {
			log.Println("Error in file uploading...", err)
			context["msg"] = "Error in file uploading."
			c.Status(500)
			return c.JSON(context)
		}

		// Set image path to the struct
		record.Image = filename
	}

	result := database.DBConn.Save(&record)

	if result.Error != nil {
		log.Println("Error in saving data.")
		context["msg"] = "Error in saving data."
		c.Status(500)
		return c.JSON(context)
	}

	context["msg"] = "Record update successfully."
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {

	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."

		return c.JSON(context)
	}

	result := database.DBConn.Delete(record)

	if result.Error != nil {
		//log.Println("Something went wrong.")
		context["msg"] = "Something went wrong."
		return c.JSON(context)
	}

	context["statusText"] = "Ok"
	context["msg"] = "Record deleted successfully."

	c.Status(200)
	return c.JSON(context)
}
