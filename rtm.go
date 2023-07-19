package rtm_package

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataUser(db *mongo.Database, id, nama, email string, hp string) (InsertedID interface{}) {
	var datacomp User
	datacomp.Iduser = id
	datacomp.Nama = nama
	datacomp.Email = email
	datacomp.Handphone = hp
	return InsertOneDoc(db, "data_user", datacomp)
}

func GetDataUserFromPhone(phone string, db *mongo.Database, col string) (data User) {
	user := db.Collection(col)
	filter := bson.M{"handphone": phone}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdatauserbyhp: %v\n", err)
	}
	return data
}

func GetDatauser(id string, db *mongo.Database, col string) (data []User) {
	user := db.Collection(col)
	filter := bson.M{"iduser": id}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("getDataUser: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func DeleteData(phone string, db *mongo.Database, col string) (data User) {
	user := db.Collection(col)
	filter := bson.M{"handphone": phone}
	err, _ := user.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataHelper : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}
