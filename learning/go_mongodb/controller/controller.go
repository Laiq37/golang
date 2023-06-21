package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"go_mongo_backend/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb://localhost:27017"

const dbName string = "netflix"

const collectionName string = "watch_list"

var collection *mongo.Collection

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connecting to db
	client, err := mongo.Connect(context.TODO(), clientOption)

	//handle error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to database")

	collection = client.Database(dbName).Collection(collectionName)
}

//db helpers

// insert a movie
func inserMovie(movie *model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), &movie)
	fmt.Println(&movie)
	//unsuccess
	if err != nil {
		log.Fatal(err)
	}

	//success
	fmt.Printf("Movie has been succesfully added in the db, movie_id : %v", inserted.InsertedID)
}

// updateMovie
func updateMovie(movieId string) {
	//required data for updating in db
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"iswatched": true}}

	//updating
	result, err := collection.UpdateOne(context.Background(), filter, update)

	//unsuccess
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", result)
	//onSuccess
	fmt.Printf("Movie has been updated successfullty, modified_count: %v", result.ModifiedCount)
}

// deleteMovie
func deleteMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)

	//onUnSuccess
	if err != nil {
		log.Fatal(err)
	}

	//onSuccess
	fmt.Printf("Movie has been deleted successfully, movies_deleted : %v", result.DeletedCount)
}

// deleteAll
func deleteMovies() {
	filter := bson.D{{}} //{} means all and we can also diectly put it in the delete query
	result, err := collection.DeleteMany(context.Background(), filter, nil)

	//onUnSuccess
	if err != nil {
		log.Fatal(err)
	}

	//onSuccess
	fmt.Printf("All movies has been deleted, movies_count: %v", result.DeletedCount)
}

// getMovies
func getMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	//onGetUnSuccess
	if err != nil {
		log.Fatal(err)
	}

	//onGetSuccess
	var movies []primitive.M //type is primitive.M because our data in the form of bson

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		fmt.Println("cursor value")
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// Controllers
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	fmt.Println(movie)
	inserMovie(&movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkMovieWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	fmt.Println(params["id"])
	updateMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteMovies()
	json.NewEncoder(w).Encode("All movies has been deleted")
}
