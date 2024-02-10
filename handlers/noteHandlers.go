package handlers

import (
	"context"
	"time"

	"example.com/jotbytes-server/config"
	"example.com/jotbytes-server/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNote creates a new note in the database
func CreateNote(c *fiber.Ctx) error {
	db := config.GetCollection("notes")
	note := new(models.Note)

	// Parse the body into the note model
	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Set creation and update times
	note.CreatedAt = time.Now()
	note.UpdatedAt = note.CreatedAt

	// Insert the note into the database
	result, err := db.InsertOne(context.Background(), note)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Fetch the inserted ID
	note.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return c.Status(fiber.StatusCreated).JSON(note)
}

// GetNotes fetches all notes from the database
func GetNotes(c *fiber.Ctx) error {
	db := config.GetCollection("notes")
	var notes []models.Note

	// Find all notes
	cursor, err := db.Find(context.Background(), bson.D{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and decode each document
	for cursor.Next(context.Background()) {
		var note models.Note
		if err := cursor.Decode(&note); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		notes = append(notes, note)
	}

	return c.JSON(notes)
}

// CreateSecureNote creates a new secure note in the database, requiring authentication
func CreateSecureNote(c *fiber.Ctx) error {
	db := config.GetCollection("secure_notes")
	note := new(models.Note)

	// Parse the body into the note model
	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Set creation and update times
	note.CreatedAt = time.Now()
	note.UpdatedAt = note.CreatedAt

	// Insert the note into the database
	result, err := db.InsertOne(context.Background(), note)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Fetch the inserted ID
	note.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return c.Status(fiber.StatusCreated).JSON(note)
}
