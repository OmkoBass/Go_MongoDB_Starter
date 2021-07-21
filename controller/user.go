package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	Database "Go_Fiber_Starter/database"
	Models "Go_Fiber_Starter/models"

	"github.com/go-playground/validator"
)

func Validate(user *Models.User) []*Models.ErrorResponse {
	var errors []*Models.ErrorResponse
	validate := validator.New()

	err := validate.Struct(user)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
            var element Models.ErrorResponse
            element.FailedField = err.StructNamespace()
            element.Tag = err.Tag()
            element.Value = err.Param()
            errors = append(errors, &element)
        }
	}

	return errors
}

func GetAllUsers(c *fiber.Ctx) error {
	var mg Models.MongoInstance = Database.GetDB()

	query := bson.D{{}}
	cursor, err := mg.Db.Collection("users").Find(c.Context(), query)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var users []Models.User = make([]Models.User, 0)

	// iterate the cursor and decode each item into an Employee
	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())

	}
	// return employees list in JSON format
	return c.Status(fiber.StatusOK).JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	var mg Models.MongoInstance = Database.GetDB()

	objectId, _ := primitive.ObjectIDFromHex(c.Params("id"));

	filter := bson.M{"_id": objectId}

	var user Models.User

	err := mg.Db.Collection("users").FindOne(c.Context(), filter).Decode(&user)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Could not find a user with that id"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func GetUserByUsername(c *fiber.Ctx, username string) Models.User {
	var mg Models.MongoInstance = Database.GetDB()

	filter := bson.M{"username": username}

	var user Models.User

	mg.Db.Collection("users").FindOne(c.Context(), filter).Decode(&user)

	return user
}

func CreateUser(c *fiber.Ctx) error {
	var mg Models.MongoInstance = Database.GetDB()

	collection := mg.Db.Collection("users")

	user := new(Models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	// Validation errors
	errors := Validate(user)
    if errors != nil {
       return c.Status(fiber.StatusBadRequest).JSON(errors);
    }

	// force MongoDB to always set its own generated ObjectIDs
	user.Id = ""

	insertionResult, err := collection.InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// get the just inserted record in order to return it as response
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	// decode the Mongo record into User
	createdUser := &Models.User{}
	createdRecord.Decode(createdUser)

	// return the created User in JSON format
	return c.Status(fiber.StatusCreated).JSON(createdUser)
}
