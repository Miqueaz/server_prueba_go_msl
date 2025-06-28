package user_model

type Struct struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
