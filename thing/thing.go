package thing

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Thing struct {
	Id    int
	Title string
}

var listOfThings = []Thing{{1, "My thing 1"}, {2, "My thing 2"}, {3, "My thing 3"}}

func GetAll(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"success": true,
		"things":  listOfThings,
		"message": "All thing listed successfully",
	})
}

func Create(c *fiber.Ctx) error {
	thing := new(Thing)

	if err := c.BodyParser(thing); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err,
		})
	}

	listOfThings = append(listOfThings, *thing)

	return c.JSON(&fiber.Map{
		"success": true,
		"things":  listOfThings,
		"message": "Thing was created successfully",
	})
}

func GetById(c *fiber.Ctx) error {
	thingId := c.Params("id")
	var thing = []Thing{}

	for _, v := range listOfThings {
		if strconv.Itoa(v.Id) == thingId {
			thing = append(thing, v)
		}
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"thing":   thing[0],
		"message": "Thing retrived successfully",
	})
}

func Update(c *fiber.Ctx) error {
	thingId := c.Params("id")
	thing := new(Thing)

	c.BodyParser(thing)

	for i, v := range listOfThings {
		if strconv.Itoa(v.Id) == thingId {
			newData := &listOfThings[i]
			newData.Title = thing.Title
		}
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"thing":   listOfThings,
		"message": "Thing with ID: " + thingId + " has been update successfully",
	})
}

func Delete(c *fiber.Ctx) error {
	thingId := c.Params("id")

	for i, v := range listOfThings {
		if strconv.Itoa(v.Id) == thingId {
			listOfThings = append(listOfThings[0:i], listOfThings[i+1:]...)
		}
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"things":  listOfThings,
		"message": "Thing was deleted successfully",
	})
}
