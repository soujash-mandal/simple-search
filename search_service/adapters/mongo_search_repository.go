package searchadapters

import (
	"context"
	internal_model "simple-search/internal/model"
	"simple-search/search_service/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDocument struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
}

type MongoSearchRepository struct {
	collection *mongo.Collection
}

func NewMongoSearchRepository(
	collection *mongo.Collection,
) *MongoSearchRepository {
	return &MongoSearchRepository{
		collection: collection,
	}
}

func (r *MongoSearchRepository) Save(
	doc model.CreateDocumentRequest,
) (internal_model.Document, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()
	result, err := r.collection.InsertOne(
		ctx,
		doc,
	)
	new_doc := internal_model.Document{
		ID:      result.InsertedID.(primitive.ObjectID).Hex(),
		Title:   doc.Title,
		Content: doc.Content,
	}
	return new_doc, err
}

func (r *MongoSearchRepository) GetAll() (
	[]internal_model.Document,
	error,
) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var docs []internal_model.Document

	for cursor.Next(ctx) {
		var mongoDoc MongoDocument

		if err := cursor.Decode(&mongoDoc); err != nil {
			return nil, err
		}

		docs = append(docs, internal_model.Document{
			ID:      mongoDoc.ID.Hex(),
			Title:   mongoDoc.Title,
			Content: mongoDoc.Content,
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return docs, nil
}