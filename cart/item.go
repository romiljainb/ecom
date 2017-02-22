package cart

//import "time"


type Item struct {
	Name         string
	ID           int
	UnitPrice    float32
	TotalPrice   float32
	Quantity     int
	Description  string
	AvgReview    int
	NumOfReviews int
	//RelatedItems map[Item]int //map
	Reviews      []Review
}

type Review struct {
	UserReview string
	ID         int
	Rating     float32
	User string
	DatePost string
	Location string
	Replies []Review
}

func (i* Item) InsertReview(r Review){
	i.Reviews = append(i.Reviews, r)
	//call updaterating to make a rolling average rating
}

func (i* Item) UpdateRatings(){
	//i.AvgReview =
}

func (i* Item) UpdateTotalPrice(){
	i.TotalPrice = float32(i.Quantity)*i.UnitPrice
}





//chat systems
//tracking items systems
//make a eCommerce platform with shoppingcart as a sub project
// search system

//func (i *Item)
