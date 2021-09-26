package entity

type MovieDetail struct {
	ID       string
	Title    string
	Year     string
	Rated    string
	Runtime  string
	Genre    string
	Director string
	Writer   string
	Actors   string
	Plot     string
	Ratings  []struct {
		Source string
		Value  string
	}
}
