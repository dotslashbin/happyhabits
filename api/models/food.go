package models

type Food struct {
	Name     string `bson:"name"`
	Note     string `bson:"note"`
	Carbs    string `bson:"carbs"`
	Proteins string `bson:"proteins"`
	Fat      string `bson:"fat"`
	Fiber    string `bson:"fiber"`
}
