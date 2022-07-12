package repository

import (
	"golang-mongodb/model"
	"golang-mongodb/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository interface {
	Add(newProduct *model.Product) error
	Retrieve(page int64, limit int64) ([]model.Product, error)
	FindById(id primitive.ObjectID) (model.Product, error)
	FindByCategory(category string) ([]model.Product, error)
	Update(id primitive.ObjectID, newData bson.D) (int64, error)
	Delete(id primitive.ObjectID) (int64, error)
}

type productRepository struct {
	db *mongo.Database
}

func (p *productRepository) Add(newProduct *model.Product) error {
	ctx, cancel := utils.InitContext()
	defer cancel()

	newProduct.Id = primitive.NewObjectID()
	_, err := p.db.Collection("products").InsertOne(ctx, newProduct)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) Retrieve(page int64, limit int64) ([]model.Product, error) {
	var products []model.Product
	ctx, cancel := utils.InitContext()
	defer cancel()

	opts := options.FindOptions{
		Skip:  &page,
		Limit: &limit,
	}

	cursor, err := p.db.Collection("products").Find(ctx, bson.D{}, &opts)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var product model.Product
		err = cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *productRepository) FindById(id primitive.ObjectID) (model.Product, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	filterID := bson.M{"_id": id}

	result := p.db.Collection("products").FindOne(ctx, filterID)

	var product model.Product
	err := result.Decode(&product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (p *productRepository) FindByCategory(category string) ([]model.Product, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	filterID := bson.M{"category": category}

	result, err := p.db.Collection("products").Find(ctx, filterID)
	if err != nil {
		return nil, err
	}

	var products []model.Product
	for result.Next(ctx) {
		var product model.Product
		err = result.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *productRepository) Update(id primitive.ObjectID, newData bson.D) (int64, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	filterID := bson.M{"_id": id}
	update := bson.D{{"$set", newData}}
	result, err := p.db.Collection("products").UpdateOne(ctx, filterID, update)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (p *productRepository) Delete(id primitive.ObjectID) (int64, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	filter := bson.D{{"_id", bson.D{{"$eq", id}}}}
	deleteDoc, err := p.db.Collection("products").DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}

	return deleteDoc.DeletedCount, nil

}

func NewProductRepository(db *mongo.Database) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
