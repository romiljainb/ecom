package cart

//import "time"
import "fmt"

type ItemID int

type Item struct {
	Name         string
	ItemIDNum   ItemID
	UnitPrice    float32
	TotalPrice   float32
	Quantity     int
	Description  string
	AvgReview    float32
	NumOfReviews int
	RelatedItems map[ItemID]int //map
	Reviews      []Review
}

type Review struct {
	UserReview string
	ID         int
	Rating     float32
	User 	   string
	DatePost   string
	Location   string
	Likes	   int
	Dislikes   int
	Points     int
	Replies    []Review
}

/* item functions */
// init()??

func (i* Item) InsertReview(r Review){
	i.Reviews = append(i.Reviews, r)
	i.NumOfReviews++
	//call update rating to make a rolling average rating later
}

func (i* Item) UpdateRatings(){
	for _, revs := range i.Reviews {
		i.AvgReview += revs.Rating
	}
	i.AvgReview = i.AvgReview/float32(i.NumOfReviews)
}

func (i* Item) UpdateTotalPrice(){
	i.TotalPrice = float32(i.Quantity)*i.UnitPrice
}

func (i* Item) DisplayItem(){
	fmt.Println()
	fmt.Printf("Item Number: %d. \n", i.ItemIDNum)
	fmt.Printf("Item Name: %s. \n", i.Name)
	fmt.Printf("Item Quantitiy: %d. \n", i.Quantity)
	fmt.Printf("Item Total Price: %f. \n", i.TotalPrice)
	fmt.Printf("Item Average Rating: %f. \n", i.AvgReview)
	fmt.Printf("Number of replies: %d. \n", len(i.Reviews))   //user recursion to get full length
	fmt.Println()
}

/* rewview functions */

func (r* Review) IncreaseLike(){
	r.Likes++
}

func (r* Review) IncreaseDilLike(){
	r.Dislikes++
}

func (r* Review) IncreasePoints(){
	r.Points++
}

func (r* Review) DisplayReview(){
	fmt.Println()
	fmt.Printf("Review number: %d. \n", r.ID)
	fmt.Printf("User: %s. \n", r.User)
	fmt.Printf("User Review: %s. \n", r.UserReview)
	fmt.Printf("User Rating: %f. \n", r.Rating)
	fmt.Printf("Number of replies: %d. \n", len(r.Replies))
	fmt.Println()
}

func (r* Review) DisplayAllReplies(){
	for _, replies:= range r.Replies {
		replies.DisplayReview()
	}
}

/* Other functions:

item spec compare

chat systems
tracking items systems
search system

*/

