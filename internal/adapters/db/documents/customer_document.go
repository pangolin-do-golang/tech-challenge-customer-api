package documents

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CustomerDocument struct {
	ID    uuid.UUID `json:"id" bson:"_id"`
	Name  string    `json:"name,omitempty" bson:"name"`
	Cpf   string    `json:"cpf" bson:"cpf"`
	Email string    `json:"email,omitempty" bson:"email"`
	Age   int       `json:"age,omitempty" bson:"age"`
}

func (d *CustomerDocument) BSON() bson.M {
	return bson.M{
		"name":  d.Name,
		"cpf":   d.Cpf,
		"age":   d.Age,
		"email": d.Email,
	}
}

func (d *CustomerDocument) BSONID() bson.M {
	return bson.M{"_id": d.ID}
}
