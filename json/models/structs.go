package models

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Address struct {
	City        string      `json:"city"`
	Street      string      `json:"street"`
	Building    string      `json:"building"`
	Coordinates Coordinates `json:"coordinates"`
}

type Restaurant struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
	Phone   string  `json:"phone"`
	Rating  float64 `json:"rating"`
	Is_open bool    `json:"is_open"`
}

type MenuItem struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Price        int      `json:"price"`
	Ingredients  []string `json:"ingredients"`
	Weight       int      `json:"weight"`
	IsVegetarian bool     `json:"is_vegetarian"`
	Spiciness    int      `json:"spiciness"`
}

type Category struct {
	Name  string     `json:"name"`
	Items []MenuItem `json:"items"`
}

type Menu struct {
	Categories []Category `json:"categories"`
}

type Review struct {
	User    string `json:"user"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
	Date    string `json:"date"`
	Likes   int    `json:"likes"`
}

type Statistics struct {
	TotalReviews  int      `json:"total_reviews"`
	AverageRating float64  `json:"average_rating"`
	PopularDishes []string `json:"popular_dishes"`
}

type Data struct {
	Restaurant Restaurant `json:"restaurant"`
	Menu       Menu       `json:"menu"`
	Reviews    []Review   `json:"reviews"`
	Statistics Statistics `json:"statistics"`
}
