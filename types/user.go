package types

type User struct {
	ID        string `bson:"_id"       json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"LastName"  json:"LastName"`
}