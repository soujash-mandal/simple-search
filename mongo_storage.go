package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func SaveDocumentMongo(doc Document) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	_, err := DocumentCollection().
		InsertOne(ctx, doc)

	return err
}

func LoadDocumentsMongo() ([]Document, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	cursor, err := DocumentCollection().
		Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var docs []Document

	if err := cursor.All(
		ctx,
		&docs,
	); err != nil {
		return nil, err
	}

	return docs, nil
}