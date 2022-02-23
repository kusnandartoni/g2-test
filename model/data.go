package model

type User struct {
	UserId  int
	Email   string
	Friends []int
}

type Rating struct {
	UserId int
	Rating int
}

type Movie struct {
	Title     string
	Duration  string
	Actors    []string
	Ratings   []Rating
	Favorites []int
	Watchlist []int
}

type Data struct {
	Movies []Movie
	Users  []User
}
