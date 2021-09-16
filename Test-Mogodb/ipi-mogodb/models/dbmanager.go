package models

import (
	"Test-Mogodb/ipi-mogodb/responest"
	"Test-Mogodb/ipi-mogodb/resquest"
	"errors"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	// CONN_HOST    = "localhost"
	// CONN_PORT    = "8080"
	MONGO_DB_URL = "127.0.0.1"
)

var session *mgo.Session
var connectionError error

func init() {
	session, connectionError = mgo.Dial(MONGO_DB_URL)
	if connectionError != nil {
		log.Fatal("error connecting to database :: ", connectionError)
	}
	session.SetMode(mgo.Monotonic, true)
}

func CreateUser(user *resquest.User) (*responest.User, error) {
	user1 := resquest.User{}
	collection := session.DB("HieuDB").C("User")
	err := collection.Find(bson.M{"id": user.Id}).One(&user1)
	if err == nil {
		fmt.Println("Object already exists in db")
		return nil, errors.New("Object already exists in db")
	}
	user.Created = time.Now()
	err = collection.Insert(&user)
	if err != nil {
		log.Print("error occurred while inserting document in database :: ", err)
		return nil, err
	}
	result := responest.User{}
	err = collection.Find(bson.M{"id": user.Id}).One(&result)
	if err != nil {
		fmt.Println("Object does not exist in db")
		return nil, errors.New("Object does not exist in db")
	}
	return &result, nil
}
func UpdateUser(user *resquest.Userudate) (*responest.User, error) {
	user1 := resquest.User{}
	collection := session.DB("HieuDB").C("User")
	err := collection.Find(bson.M{"id": user.Id}).One(&user1)
	if err != nil {
		fmt.Println("Object does not exist in db")
		return nil, errors.New("Object does not exist in db")
	}
	// _, err = collection.Upsert(bson.M{"id": user.Id}, &user)
	// if err != nil {
	// 	log.Print("error occurred while updating document in database :: ", err)
	// 	return nil, err
	// }
	if user.Name != "" {
		err = collection.Update(bson.M{"id": user.Id}, bson.M{"$set": bson.M{"name": user.Name}})
		if err != nil {
			log.Print("error occurred while updating Name in database :: ", err)
			return nil, err
		}
	}
	if user.Password != "" {
		err = collection.Update((bson.M{"id": user.Id}), bson.M{"$set": bson.M{"password": user.Password}})
		if err != nil {
			log.Print("error occurred while updating Password in database :: ", err)
			return nil, err
		}
	}
	if user.Phone != "" {
		err = collection.Update(bson.M{"id": user.Id}, bson.M{"$set": bson.M{"phone": user.Phone}})
		if err != nil {
			log.Print("error occurred while updating Phone in database :: ", err)
			return nil, err
		}
	}
	if user.Imail != "" {
		err = collection.Update(bson.M{"id": user.Id}, bson.M{"$set": bson.M{"imail": user.Imail}})
		if err != nil {
			log.Print("error occurred while updating imail in database :: ", err)
			return nil, err
		}
	}
	user.Update = time.Now()
	err = collection.Update(bson.M{"id": user.Id}, bson.M{"$set": bson.M{"update": user.Update}})
	if err != nil {
		log.Print("error occurred while updating imail in database :: ", err)
		return nil, err
	}
	result := responest.User{}
	err = collection.Find(bson.M{"id": user.Id}).One(&result)
	if err != nil {
		fmt.Println("Object does not exist in db")
		return nil, errors.New("Object does not exist in db")
	}
	return &result, nil
}

func FindUser(id int64) (*responest.User, error) {
	user1 := responest.User{}
	collection := session.DB("HieuDB").C("User")
	err := collection.Find(bson.M{"id": id}).One(&user1)
	if err != nil {
		fmt.Println("Object does not exist in db")
		return nil, errors.New("Object does not exist in db")
	}
	return &user1, nil
}
func LisAllUser() ([]*responest.User, error) {
	list := []*responest.User{}
	collection := session.DB("HieuDB").C("User")
	err := collection.Find(bson.M{}).All(&list)
	if err != nil {
		log.Print("error occurred while reding document in database :: ", err)
		return nil, err
	}
	return list, nil
}
func LisName(name string) ([]*responest.User, error) {
	list := []*responest.User{}
	collection := session.DB("HieuDB").C("User")
	err := collection.Find(bson.M{"name": name}).All(&list)
	if err != nil {
		log.Print("error occurred while reding document in database :: ", err)
		return nil, err
	}
	return list, nil
}
func DeleteUser(id int64) (string, error) {
	user := responest.User{}
	collection := session.DB("HieuDB").C("User")
	err := collection.Find(bson.M{"id": id}).One(&user)
	if err != nil {
		fmt.Println("Object does not exist in db")
		return "", errors.New("Object does not exist in db")
	}
	err = collection.Remove(bson.M{"id": id})
	if err != nil {
		log.Print("error removing document from database :: ", err)
		return "", err
	}
	return "Deleted successfully", nil
}
