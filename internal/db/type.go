package db

import "go.mongodb.org/mongo-driver/mongo"

// {
//  "hostname": [
//         "hostname:27017",
//         "hostname:27018",
//         "hostname:27019"
// 	],
// 	"replicaset": "namereplicaset"
//     "dbname": "dbname",
//     "user": "admin",
// 	"password": "123"
// }

type MongoConfig struct {
	Hostname   []string `json:"hostname"`
	DBname     string   `json:"dbname"`
	User       string   `json:"user"`
	Password   string   `json:"password"`
	ReplicaSet string   `json:"replicaset"`
	URI        string   `json:"-"`
}

type DB struct {
	Client *mongo.Client
}
