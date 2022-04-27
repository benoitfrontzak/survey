package main

import (
	"log"
	"surveyGenerator/mongoDB"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get all documents from questionnaires collection
func getQuestionnaires() (questionnaires []Questionnaire, err error) {
	// Specify db connection properties.
	// getEnv("Expected variable", "expected default value if not found in .env")
	var (
		dbname     = getEnv("DBNAME", "")
		dbuser     = getEnv("DBUSER", "")
		dbpass     = getEnv("DBPASS", "")
		dburl      = getEnv("DBURL", "127.0.0.1")
		dbport     = getEnv("DBPORT", "27017")
		collection = "questionnaires"
	)
	// filter (all not soft deleted)
	filter := bson.D{
		primitive.E{Key: "deletedAt", Value: nil},
	}
	//option
	option := bson.D{}

	// Generate DSN
	dsn, err := mongoDB.DSN(dbname, dbuser, dbpass, dburl, dbport)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("mongodb response:")
	// Get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := mongoDB.Connect(dsn, dbname, dbuser, dbpass)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Release resource when the main function is returned.
	defer mongoDB.Close(client, ctx, cancel)

	// call the query method with client, context,
	// database name, collection  name, filter and option
	// This method returns mongo.cursor and error if any.
	cursor, err := mongoDB.Query(client, ctx, dbname, collection, filter, option)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// to get bson object  from cursor,
	// returns error if any.
	if err := cursor.All(ctx, &questionnaires); err != nil {
		log.Println(err)
		return nil, err
	}

	return questionnaires, nil
}

// Get all documents from questionnaires collection
func getQuestionnaire(hexid string) (questionnaire Questionnaire, err error) {
	// Specify db connection properties.
	// getEnv("Expected variable", "expected default value if not found in .env")
	var (
		dbname     = getEnv("DBNAME", "")
		dbuser     = getEnv("DBUSER", "")
		dbpass     = getEnv("DBPASS", "")
		dburl      = getEnv("DBURL", "127.0.0.1")
		dbport     = getEnv("DBPORT", "27017")
		collection = "questionnaires"
	)

	var questionnaires = []Questionnaire{}

	id, err := primitive.ObjectIDFromHex(hexid)
	if err != nil {
		log.Println(err)
	}

	// filter
	filter := bson.D{
		primitive.E{Key: "_id", Value: id},
	}

	//option
	option := bson.D{}

	// Generate DSN
	dsn, err := mongoDB.DSN(dbname, dbuser, dbpass, dburl, dbport)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("mongodb response:")
	// Get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := mongoDB.Connect(dsn, dbname, dbuser, dbpass)
	if err != nil {
		log.Println(err)
		return
	}

	// Release resource when the main function is returned.
	defer mongoDB.Close(client, ctx, cancel)

	// call the query method with client, context,
	// database name, collection  name, filter and option
	// This method returns mongo.cursor and error if any.
	cursor, err := mongoDB.Query(client, ctx, dbname, collection, filter, option)
	if err != nil {
		log.Println(err)
		return
	}

	// to get bson object  from cursor,
	// returns error if any.
	if errr := cursor.All(ctx, &questionnaires); err != nil {
		log.Println(errr)
		return
	}

	return questionnaires[0], err
}

// Get all documents from evaluations collection
func getEvaluations() (evaluations []Evaluation, err error) {
	// Specify db connection properties.
	// getEnv("Expected variable", "expected default value if not found in .env")
	var (
		dbname     = getEnv("DBNAME", "")
		dbuser     = getEnv("DBUSER", "")
		dbpass     = getEnv("DBPASS", "")
		dburl      = getEnv("DBURL", "127.0.0.1")
		dbport     = getEnv("DBPORT", "27017")
		collection = "evaluations"
	)
	// filter (all not soft deleted)
	filter := bson.D{
		primitive.E{Key: "deletedAt", Value: nil},
	}
	//option
	option := bson.D{}

	// Generate DSN
	dsn, err := mongoDB.DSN(dbname, dbuser, dbpass, dburl, dbport)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("mongodb response:")
	// Get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := mongoDB.Connect(dsn, dbname, dbuser, dbpass)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Release resource when the main function is returned.
	defer mongoDB.Close(client, ctx, cancel)

	// call the query method with client, context,
	// database name, collection  name, filter and option
	// This method returns mongo.cursor and error if any.
	cursor, err := mongoDB.Query(client, ctx, dbname, collection, filter, option)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// to get bson object  from cursor,
	// returns error if any.
	if err := cursor.All(ctx, &evaluations); err != nil {
		log.Println(err)
		return nil, err
	}

	return evaluations, nil
}

// Get evaluation id from evaluations collection
func getEvaluation(hexid string) (evaluation Evaluation, err error) {
	// Specify db connection properties.
	// getEnv("Expected variable", "expected default value if not found in .env")
	var (
		dbname     = getEnv("DBNAME", "")
		dbuser     = getEnv("DBUSER", "")
		dbpass     = getEnv("DBPASS", "")
		dburl      = getEnv("DBURL", "127.0.0.1")
		dbport     = getEnv("DBPORT", "27017")
		collection = "evaluations"
	)

	var evaluations = []Evaluation{}

	id, err := primitive.ObjectIDFromHex(hexid)
	if err != nil {
		log.Println(err)
	}

	// filter
	filter := bson.D{
		primitive.E{Key: "_id", Value: id},
	}

	//option
	option := bson.D{}

	// Generate DSN
	dsn, err := mongoDB.DSN(dbname, dbuser, dbpass, dburl, dbport)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("mongodb response:")
	// Get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := mongoDB.Connect(dsn, dbname, dbuser, dbpass)
	if err != nil {
		log.Println(err)
		return
	}

	// Release resource when the main function is returned.
	defer mongoDB.Close(client, ctx, cancel)

	// call the query method with client, context,
	// database name, collection  name, filter and option
	// This method returns mongo.cursor and error if any.
	cursor, err := mongoDB.Query(client, ctx, dbname, collection, filter, option)
	if err != nil {
		log.Println(err)
		return
	}

	// to get bson object  from cursor,
	// returns error if any.
	if errr := cursor.All(ctx, &evaluations); err != nil {
		log.Println(errr)
		return
	}

	return evaluations[0], err
}

// Get all documents of questionnaire id from evaluations collection
func getAllEvaluations(hexid string) (evaluations []Evaluation, err error) {
	// Specify db connection properties.
	// getEnv("Expected variable", "expected default value if not found in .env")
	var (
		dbname     = getEnv("DBNAME", "")
		dbuser     = getEnv("DBUSER", "")
		dbpass     = getEnv("DBPASS", "")
		dburl      = getEnv("DBURL", "127.0.0.1")
		dbport     = getEnv("DBPORT", "27017")
		collection = "evaluations"
	)

	// filter
	filter := bson.D{
		primitive.E{Key: "qid", Value: hexid},
	}

	//option
	option := bson.D{}

	// Generate DSN
	dsn, err := mongoDB.DSN(dbname, dbuser, dbpass, dburl, dbport)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("mongodb response:")
	// Get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := mongoDB.Connect(dsn, dbname, dbuser, dbpass)
	if err != nil {
		log.Println(err)
		return
	}

	// Release resource when the main function is returned.
	defer mongoDB.Close(client, ctx, cancel)

	// call the query method with client, context,
	// database name, collection  name, filter and option
	// This method returns mongo.cursor and error if any.
	cursor, err := mongoDB.Query(client, ctx, dbname, collection, filter, option)
	if err != nil {
		log.Println(err)
		return
	}

	// to get bson object  from cursor,
	// returns error if any.
	if errr := cursor.All(ctx, &evaluations); err != nil {
		log.Println(errr)
		return
	}

	return evaluations, err
}
