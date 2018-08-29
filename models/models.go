package models

import (
  "time"

  "gopkg.in/mgo.v2/bson"
)

type (
  User struct {
    Id  bson.ObjectId `bson:"_id,omitempty" json:"id"`
    FirstName string `json:"firstname"`
    LastName string `json:"lastname"`
    Email string `json:"email"`
    Password string `json:"password,omitempty"`
    HashPassword []byte `json:"hashpassword,omitempty "`
  }
  Student struct {
    Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
    FirstName string `json:"firstname"`
    LastName string `json:"lastname"`
    CreatedBy string `json:"createdBy"`
  }
  Lunch struct {
    Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
    StudentId string `json:"studentId"`
    Name string `json:"name"`
    Description string `json:"description"`
    Schedule time.Time `json:"
  }
)





