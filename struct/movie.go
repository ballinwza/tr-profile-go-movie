package struct_movie

type Movie struct {
	TmdbId      int     `json:"tmbd_id" bson:"tmbd_id"`
	Title       string  `json:"title" bson:"title"`
	Description string  `json:"description" bson:"description"`
	Adult       bool    `json:"adult" bson:"adult"`
	Genres      []int   `json:"genres" bson:"genres"`
	Rating      float32 `json:"rating" bson:"rating"`
}
