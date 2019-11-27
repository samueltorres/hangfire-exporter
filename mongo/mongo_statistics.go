package mongo

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoStatistics is a MongoDB implementation to get hangfire statistics
type MongoStatistics struct {
	client       *mongo.Client
	databaseName string
	logger       *log.Logger
}

// NewMongoStatistics creates a new MongoStatistics
func NewMongoStatistics(conn string, databaseName string, logger *log.Logger) (*MongoStatistics, error) {
	clientOpts := options.Client().ApplyURI(conn)

	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &MongoStatistics{
		client:       client,
		logger:       logger,
		databaseName: databaseName,
	}, nil
}

// Available checks if the hangfire database is reachable
func (m *MongoStatistics) Available() bool {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := m.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return false
	}

	return true
}

// DeletedJobs gets the total number of deleted jobs
func (m *MongoStatistics) DeletedJobs() float64 {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	r := m.client.Database(m.databaseName).
		Collection("hangfire.jobGraph").
		FindOne(
			ctx,
			bson.M{
				"_t":  "CounterDto",
				"Key": "stats:deleted",
			})

	if err := r.Err(); err != nil {
		return 0
	}

	var counter struct {
		Value int64
	}

	err := r.Decode(&counter)
	if err != nil {
		return 0
	}

	return float64(counter.Value)
}

// EnqueuedJobs gets the current number of enqueued jobs
func (m *MongoStatistics) EnqueuedJobs() float64 {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c, err := m.client.Database(m.databaseName).
		Collection("hangfire.jobGraph").
		CountDocuments(
			ctx,
			bson.M{
				"_t":        "JobQueueDto",
				"FetchedAt": nil,
			})

	if err != nil {
		return 0
	}

	return float64(c)
}

// FailedJobs gets the total number of failed jobs
func (m *MongoStatistics) FailedJobs() float64 {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c, err := m.client.Database(m.databaseName).
		Collection("hangfire.jobGraph").
		CountDocuments(
			ctx,
			bson.M{
				"_t":        "JobDto",
				"StateName": "Failed",
			})

	if err != nil {
		return 0
	}

	return float64(c)
}

// FetchedJobs gets the current number of fetched jobs
func (m *MongoStatistics) FetchedJobs() float64 {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c, err := m.client.Database(m.databaseName).
		Collection("hangfire.jobGraph").
		CountDocuments(
			ctx,
			bson.M{
				"_t": "JobQueueDto",
				"FetchedAt": bson.M{
					"$ne": nil,
				},
			})

	if err != nil {
		return 0
	}

	return float64(c)
}

// ProcessingJobs gets the current number of processing jobs
func (m *MongoStatistics) ProcessingJobs() float64 {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c, err := m.client.Database(m.databaseName).
		Collection("hangfire.jobGraph").
		CountDocuments(
			ctx,
			bson.M{
				"_t":        "JobDto",
				"StateName": "Processing",
			})

	if err != nil {
		return 0
	}

	return float64(c)
}

// Queues gets the current number of queues
func (m *MongoStatistics) Queues() float64 {
	return 0
}

// RecurringJobs gets the current number of recurring jobs
func (m *MongoStatistics) RecurringJobs() float64 {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c, err := m.client.Database(m.databaseName).
		Collection("hangfire.jobGraph").
		CountDocuments(
			ctx,
			bson.M{
				"_t": "SetDto",
				"Key": bson.M{
					"$regex": "^recurring-jobs",
				},
			})

	if err != nil {
		return 0
	}

	return float64(c)
}

// ScheduledJobs gets the current nubmer of scheduled jobs
func (m *MongoStatistics) ScheduledJobs() float64 {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c, err := m.client.Database(m.databaseName).
		Collection("hangfire.jobGraph").
		CountDocuments(
			ctx,
			bson.M{
				"_t":        "JobDto",
				"StateName": "Scheduled",
			})

	if err != nil {
		return 0
	}

	return float64(c)
}

// Servers gets the number of registered servers on hangfire database
func (m *MongoStatistics) Servers() float64 {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	c, err := m.client.Database(m.databaseName).
		Collection("hangfire.server").
		CountDocuments(ctx, bson.M{})

	if err != nil {
		return 0
	}

	return float64(c)
}

// SucceededJobs gets the total number of succeeded jobs
func (m *MongoStatistics) SucceededJobs() float64 {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	r := m.client.Database(m.databaseName).
		Collection("hangfire.jobGraph").
		FindOne(
			ctx,
			bson.M{
				"_t":  "CounterDto",
				"Key": "stats:succeeded",
			})

	if err := r.Err(); err != nil {
		return 0
	}

	var counter struct {
		Value int64
	}

	err := r.Decode(&counter)
	if err != nil {
		return 0
	}

	return float64(counter.Value)
}
