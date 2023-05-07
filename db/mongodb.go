package db

import (
	"context"
	"dipay-test/models"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbCon struct {
	mongodbUri string
	db         *mongo.Client
}

func NewMongodb(mongoaddr string) *MongodbCon {
	return &MongodbCon{
		mongodbUri: mongoaddr,
		// db: ,
	}
}

func (m *MongodbCon) InitDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.mongodbUri))
	if err != nil {
		panic(err)
	}
	m.db = client
}

func (m *MongodbCon) GetMongoInstance() *mongo.Client {
	return m.db
}

func (m *MongodbCon) IsCompanyExist(company_name string) bool {
	coll := m.db.Database("dipayDB").Collection("companies")
	filter := bson.D{{"company_name", company_name}}
	var result models.Company
	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return false
		}
	}

	return true
}

func (m *MongodbCon) AddCompany(c models.Company) (interface{}, error) {
	coll := m.db.Database("dipayDB").Collection("companies")
	res, err := coll.InsertOne(context.TODO(), c)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func (m *MongodbCon) GetComapnies() ([]models.Company, error) {
	coll := m.db.Database("dipayDB").Collection("companies")
	filter := bson.D{}
	var results []models.Company

	curr, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return []models.Company{}, err
	}

	if err := curr.All(context.TODO(), &results); err != nil {
		return []models.Company{}, err
	}

	fmt.Println(results)

	return results, nil
}

func (m *MongodbCon) SetCompanyToActive(id string) error {
	coll := m.db.Database("dipayDB").Collection("companies")

	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", obj_id}}
	update := bson.D{{"$set", bson.D{{"is_active", true}}}}

	res, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	//kalo matchednya 0 berarti ga ada (not found)
	if res.MatchedCount == 0 {
		return errors.New("not_found")
	}

	//kalo modifiednya 0 artinya sudah aktif
	if res.ModifiedCount == 0 {
		return errors.New("already_active")
	}

	return nil
}

func (m *MongodbCon) IsEmployeeEmailExist(email string) bool {
	coll := m.db.Database("dipayDB").Collection("employees")
	filter := bson.D{{"email", email}}
	var result models.Employee
	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return false
		}
	}

	return true
}

func (m *MongodbCon) AddEmployee(e models.Employee) (interface{}, error) {
	coll := m.db.Database("dipayDB").Collection("employees")
	res, err := coll.InsertOne(context.TODO(), e)
	if err != nil {
		return models.Employee{}, err
	}

	return res.InsertedID, nil
}

func (m *MongodbCon) GetEmployee(id string) (models.Employee, error) {
	coll := m.db.Database("dipayDB").Collection("employees")
	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Employee{}, err
	}

	filter := bson.D{{"_id", obj_id}}
	var result models.Employee
	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return models.Employee{}, err
		}
	}

	return result, nil
}
func (m *MongodbCon) GetEmployeeByCompanyId(id string) ([]models.Employee, error) {
	coll := m.db.Database("dipayDB").Collection("employees")
	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return []models.Employee{}, err
	}

	filter := bson.D{{"company_id", obj_id}}
	var results []models.Employee
	curr, err := coll.Find(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return []models.Employee{}, err
		}
	}

	if err := curr.All(context.TODO(), &results); err != nil {
		return []models.Employee{}, err
	}

	return results, nil
}

func (m *MongodbCon) DeleteEmployee(id string) error {
	coll := m.db.Database("dipayDB").Collection("employees")
	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", obj_id}}
	res, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("nothing_to_delete")
	}

	return nil
}

func (m *MongodbCon) UpdateEmployee(d models.Employee, employee_id, company_id string) error {
	coll := m.db.Database("dipayDB").Collection("employees")

	employee_obj_id, err := primitive.ObjectIDFromHex(employee_id)
	if err != nil {
		return err
	}
	company_obj_id, err := primitive.ObjectIDFromHex(company_id)
	if err != nil {
		return err
	}

	filter := bson.D{{"company_id", company_obj_id}, {"_id", employee_obj_id}}
	update := bson.D{{"$set", bson.D{{"name", d.Name}, {"email", d.Email}, {"phone_number", d.Phone_Number}, {"jobtitle", d.Job_Title}}}}

	_, err = coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
