package types

type User struct {
	ID        string `bson:"_id,omitempty"       json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"LastName"  json:"LastName"`
}
