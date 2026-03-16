package models

type UserCreds struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type UserProfile struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Email    string `json:"email"    bson:"email"`
	Name     string `json:"name"     bson:"name"`
	Age      int    `json:"age"      bson:"age"`
	Gender   string `json:"gender"   bson:"gender"`
}
