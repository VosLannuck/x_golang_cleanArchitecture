package db

import (
	"VosLannuck/x_golang_cleanArchitecture/src/domain"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBHandler struct {
  MongoClient mongo.Client
  database *mongo.Database
}


func NewDBHandler(connectionString string, dbName string) (DBHandler, error) {
  dbHandler := DBHandler{}
  
  clientOptions := options.Client().ApplyURI(connectionString)
  client, err := mongo.Connect(context.TODO(), clientOptions)

  if err != nil {
    log.Fatal(err)
    return dbHandler, err
  }

  err = client.Ping(context.TODO(), nil)

  if err != nil {
    log.Fatal(err)
    return dbHandler, err
  }

  dbHandler.MongoClient = *client
  dbHandler.database = client.Database(dbName)


  return dbHandler, nil
}

func (dbHandler DBHandler) FindAllBooks() ([]*domain.Book, error){
  
  var books []*domain.Book

  collection := dbHandler.database.Collection("books")
  cursor, err := collection.Find(context.TODO(), bson.D{})
  
  if err != nil {
    return nil, err
  }

  for cursor.Next(context.TODO()) {
    var book domain.Book
    var err_decode error = cursor.Decode(&book)
    
    if err_decode != nil {
      log.Fatal(err_decode)
    }

    books = append(books, &book)
    
  }
  
  return books, nil
}

func (dbHandler DBHandler) SaveBook(book domain.Book) (error){
  
  collection := dbHandler.database.Collection("books")
  _, err := collection.InsertOne(context.TODO(), book)
  
  if err != nil {
      return err
  }

  return nil
}


func (dbHandler DBHandler) SaveAuthor(author domain.Author) (error) {

  collection := dbHandler.database.Collection("authors")
  _, err := collection.InsertOne(context.TODO(), author)

  if err != nil {
    return err
  }

  return nil
}
