// This app run with mongoDB
// We don't want to have the configuration of the dsn set with global variable
// so we only defined it in a separate function DSN
// We don't want neither to have the context set with global variable
// so we only return it with passby function
// This package provides:
// - 3 core methods to interact with mongoDB :
// 		connect | ping | close
// - crud methods :
// 		insertOne | query | updateOne | deleteOne
// We don't want to use insertMany & updateMany,
// we loop insertOne & updateOne when needed
package mongoDB

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//  Returns the Data Source Name or error if any
func DSN(NameDB, dbuser, dbpass, dburl, dbport string) (dsn string, dsnErr error) {
	if dbuser == "" {
		// mongodb://dburl:dbport
		return fmt.Sprintf("mongodb://%s:%s", dburl, dbport), nil
	}
	if dbuser != "" {
		// mongodb://dbuser:dbpass@dburl:dbport/?authSource=NameDB&authMechanism=SCRAM-SHA-256&ssl=false
		// return fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s&authMechanism=SCRAM-SHA-256&ssl=false", dbuser, dbpass, dburl, dbport, NameDB), nil

		// mongodb+srv://<username>:<password>@bencluster.d1xfu.gcp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority
		return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", dbuser, dbpass, dburl, NameDB), nil
	}
	badDSN := errors.New("the DSN configuration is incomplete")
	return "", badDSN
}

// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.
func Connect(uri, NameDB, dbuser, dbpass string) (client *mongo.Client, ctx context.Context, cancel context.CancelFunc, err error) {
	// AUTHENTICATED METHOD
	if dbuser != "" {
		// To configure auth via URI with SCRAM-SHA-256 mechanism
		credential := options.Credential{
			AuthMechanism:           "SCRAM-SHA-1",       //SCRAM-SHA-1|SCRAM-SHA-256|MONGODB-X509|MONGODB-AWS|GSSAPI(Kerberos)|PLAIN (LDAP SASL)
			AuthMechanismProperties: map[string]string{}, //SERVICE_NAME|CANONICALIZE_HOST_NAME|SERVICE_REALM|AWS_SESSION_TOKEN
			AuthSource:              NameDB,
			Username:                dbuser,
			Password:                dbpass,
			PasswordSet:             false, //For GSSAPI, this must be true if a password is specified
		}

		// Configure a Client with SCRAM authentication
		// The default authentication database for SCRAM is "admin".
		// So we overwrite it with authSource query parameter in the URI
		// SCRAM is the default auth mechanism so specifying a mechanism is not required.
		opts := options.Client().ApplyURI(uri).SetAuth(credential)

		// ctx will be used to set deadline for process, here
		// deadline will of 10 seconds.
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		// mongo.Connect return mongo.Client method
		client, err = mongo.Connect(ctx, opts)

		return
	}

	// NON AUTHENTICATED METHOD

	// ctx will be used to set deadline for process, here
	// deadline will of 10 seconds.
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	// mongo.Connect return mongo.Client method
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method is used to ping the mongoDB, return error if any.
func PingDB(client *mongo.Client, ctx context.Context) (err error) {
	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	log.Println("mongoDB server is online")
	return nil
}

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			log.Println(err)
		}
	}()
}

// insertOne is a user defined method, used to insert
// documents into collection returns result of InsertOne
// and error if any.
func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	// select database and collection ith Client.Database method
	// and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

// query is user defined method used to query MongoDB,
// that accepts mongo.client,context, database name,
// collection name, a query and field.

//  database name and collection name is of type
// string. query is of type interface.
// field is of type interface, which limts
// the field being returned.

// query method returns a cursor and error.
func Query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
	return
}

// UpdateOne is a user defined method, that update
// a single document matching the filter.
// This methods accepts client, context, database,
// collection, filter and update filter and update
// is of type interface this method returns
// UpdateResult and an error if any.
func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {

	// select the database and the collection
	collection := client.Database(dataBase).Collection(col)

	// A single document that match with the
	// filter will get updated.
	// update contains the filed which should get updated.
	result, err = collection.UpdateOne(ctx, filter, update)
	return
}

// deleteOne is a user defined function that delete,
// a single document from the collection.
// Returns DeleteResult and an  error if any.
func DeleteOne(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {

	// select document and collection
	collection := client.Database(dataBase).Collection(col)

	// query is used to match a document  from the collection.
	result, err = collection.DeleteOne(ctx, query)
	return
}
