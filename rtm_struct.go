package rtm_package

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Iduser    string             `bson:"iduser,omitempty" json:"iduser,omitempty"`
	Nama      string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Handphone string             `bson:"handphone,omitempty" json:"handphone,omitempty"`
}
