package main

import (
	"log"
	"surveyGenerator/mongoDB"
)

// Check is mongoDB server is alive
func pingMongoDB() (err error) {
	// Specify db connection properties.
	// getEnv("Expected variable", "expected default value if not found in .env")
	var (
		dbname = getEnv("DBNAME", "")
		dbuser = getEnv("DBUSER", "")
		dbpass = getEnv("DBPASS", "")
		dburl  = getEnv("DBURL", "127.0.0.1")
		dbport = getEnv("DBPORT", "27017")
	)

	// Generate DSN
	dsn, err := mongoDB.DSN(dbname, dbuser, dbpass, dburl, dbport)
	if err != nil {
		log.Println(err)
		return
	}

	// Get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := mongoDB.Connect(dsn, dbname, dbuser, dbpass)
	if err != nil {
		log.Println(err)
		return
	}

	// Release resource when the main function is returned.
	defer mongoDB.Close(client, ctx, cancel)

	// Ping server and return error if any
	err = mongoDB.PingDB(client, ctx)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

// Create document
func createDocument(collection string, query interface{}) (rowID interface{}, err error) {
	// Specify db connection properties.
	// getEnv("Expected variable", "expected default value if not found in .env")
	var (
		dbname = getEnv("DBNAME", "")
		dbuser = getEnv("DBUSER", "")
		dbpass = getEnv("DBPASS", "")
		dburl  = getEnv("DBURL", "127.0.0.1")
		dbport = getEnv("DBPORT", "27017")
	)
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

	// insertOne accepts client , context, database
	// name collection name and an interface that
	// will be inserted into the  collection.
	// insertOne returns an error and a result of
	// insert in a single document into the collection.
	insertOneResult, err := mongoDB.InsertOne(client, ctx, dbname, collection, query)

	// handle the error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return insertOneResult.InsertedID, nil
}

// Update document
func updateDocument(collection string, filter, query interface{}) (int64, error) {
	// Specify db connection properties.
	// getEnv("Expected variable", "expected default value if not found in .env")
	var (
		dbname = getEnv("DBNAME", "")
		dbuser = getEnv("DBUSER", "")
		dbpass = getEnv("DBPASS", "")
		dburl  = getEnv("DBURL", "127.0.0.1")
		dbport = getEnv("DBPORT", "27017")
	)
	// Generate DSN
	dsn, err := mongoDB.DSN(dbname, dbuser, dbpass, dburl, dbport)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	log.Println("mongodb response:")
	// Get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := mongoDB.Connect(dsn, dbname, dbuser, dbpass)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// Release resource when the main function is returned.
	defer mongoDB.Close(client, ctx, cancel)

	// Returns result of updated document and a error.
	result, err := mongoDB.UpdateOne(client, ctx, dbname, collection, filter, query)

	// handle error
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// print count of documents that affected
	return result.ModifiedCount, err
}

// Delete document
func deleteDocument(collection string, query interface{}) (int64, error) {
	// Specify db connection properties.
	// getEnv("Expected variable", "expected default value if not found in .env")
	var (
		dbname = getEnv("DBNAME", "")
		dbuser = getEnv("DBUSER", "")
		dbpass = getEnv("DBPASS", "")
		dburl  = getEnv("DBURL", "127.0.0.1")
		dbport = getEnv("DBPORT", "27017")
	)
	// Generate DSN
	dsn, err := mongoDB.DSN(dbname, dbuser, dbpass, dburl, dbport)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	log.Println("mongodb response:")
	// Get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := mongoDB.Connect(dsn, dbname, dbuser, dbpass)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// Release resource when the main function is returned.
	defer mongoDB.Close(client, ctx, cancel)

	// Returns result of deletion and error
	result, err := mongoDB.DeleteOne(client, ctx, dbname, collection, query)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return result.DeletedCount, nil
}
