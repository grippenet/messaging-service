package messagedb

import (
	"context"
	"log"
	"time"

	"github.com/influenzanet/messaging-service/pkg/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageDBService struct {
	DBClient     *mongo.Client
	timeout      int
	DBNamePrefix string
}

func NewMessageDBService(configs types.DBConfig) *MessageDBService {
	var err error
	dbClient, err := mongo.NewClient(
		options.Client().ApplyURI(configs.URI),
		options.Client().SetMaxConnIdleTime(time.Duration(configs.IdleConnTimeout)*time.Second),
		options.Client().SetMaxPoolSize(configs.MaxPoolSize),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(configs.Timeout)*time.Second)
	defer cancel()

	err = dbClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctx, conCancel := context.WithTimeout(context.Background(), time.Duration(configs.Timeout)*time.Second)
	err = dbClient.Ping(ctx, nil)
	defer conCancel()
	if err != nil {
		log.Fatal("fail to connect to DB: " + err.Error())
	}

	return &MessageDBService{
		DBClient:     dbClient,
		timeout:      configs.Timeout,
		DBNamePrefix: configs.DBNamePrefix,
	}
}

// Collections
func (dbService *MessageDBService) collectionRefEmailTemplates(instanceID string) *mongo.Collection {
	return dbService.DBClient.Database(dbService.DBNamePrefix + instanceID + "_messageDB").Collection("email-templates")
}

func (dbService *MessageDBService) collectionRefOutgoingEmails(instanceID string) *mongo.Collection {
	return dbService.DBClient.Database(dbService.DBNamePrefix + instanceID + "_messageDB").Collection("outgoing-emails")
}

func (dbService *MessageDBService) collectionRefSentEmails(instanceID string) *mongo.Collection {
	return dbService.DBClient.Database(dbService.DBNamePrefix + instanceID + "_messageDB").Collection("sent-emails")
}

// DB utils
func (dbService *MessageDBService) getContext() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(dbService.timeout)*time.Second)
}
