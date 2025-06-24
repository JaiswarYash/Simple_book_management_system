package models

import (
	"context"
	"log"
	"time"

	"github.com/mr-yash-dev/Book-management-system/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

type Book struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Author      string             `bson:"author" json:"author"`
	Publication string             `bson:"publication" json:"publication"`
}

func init() {
	db := config.Connect()
	collection = db.Collection("books")
}

func (b *Book) CreateBook() *Book {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, b)
	if err != nil {
		log.Printf("Error creating book: %v", err)
		return nil
	}

	b.ID = result.InsertedID.(primitive.ObjectID)
	return b
}

func GetAllBooks() []Book {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var books []Book
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error finding books: %v", err)
		return nil
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &books); err != nil {
		log.Printf("Error decoding books: %v", err)
		return nil
	}

	return books
}

func GetBookById(id string) *Book {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting id: %v", err)
		return nil
	}

	var book Book
	err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&book)
	if err != nil {
		log.Printf("Error finding book: %v", err)
		return nil
	}

	return &book
}

func DeleteBook(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting id: %v", err)
		return false
	}

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		log.Printf("Error deleting book: %v", err)
		return false
	}

	return result.DeletedCount > 0
}

func UpdateBook(id string, book *Book) *Book {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting id: %v", err)
		return nil
	}

	update := bson.M{
		"$set": bson.M{
			"name":        book.Name,
			"author":      book.Author,
			"publication": book.Publication,
		},
	}

	result := collection.FindOneAndUpdate(ctx, bson.M{"_id": objectId}, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if err := result.Decode(book); err != nil {
		log.Printf("Error updating book: %v", err)
		return nil
	}

	return book
}
