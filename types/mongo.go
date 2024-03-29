package types

type User struct {
	User   string  `json:"user" bson:"user"`
	Bucket []int64 `json:"bucket" bson:"bucket"`
}

type Content struct {
	Name      string `json:"name" bson:"name"`
	Price     int64  `json:"price" bson:"price"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"`
}

type History struct {
	User        string   `json:"user" bson:"user"`
	ContentList []string `json:"contentList" bson:"contentList"`
	CreatedAt   int64    `json:"createdAt" bson:"createdAt"`
}

type BucketRequest struct {
	User string `form:"user" binding:"required"`
}

type BucketResponse struct {
	Response *User `json:"response" bson:"response"`
}
