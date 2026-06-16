package searchadapters

import (
	"context"
	models "simple-search/search_service/model"
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
	doc models.CreateDocumentRequest,
) (models.Document, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()
	result, err := r.collection.InsertOne(
		ctx,
		doc,
	)
	new_doc := models.Document{
		ID:      result.InsertedID.(primitive.ObjectID).Hex(),
		Title:   doc.Title,
		Content: doc.Content,
	}
	return new_doc, err
}

func (r *MongoSearchRepository) GetAll() (
	[]models.Document,
	error,
) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	cursor, err := r.collection.Find(
		ctx,
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var mongoDocs []MongoDocument

	if err := cursor.All(
		ctx,
		&mongoDocs,
	); err != nil {
		return nil, err
	}

	var domain_docs []models.Document
	for _, doc := range mongoDocs {
		domain_docs = append(
			domain_docs,
			models.Document{ID: doc.ID.Hex(), Title: doc.Title, Content: doc.Content},
		)
	}

	return domain_docs, nil
}
