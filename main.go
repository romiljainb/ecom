package main

import (
	"fmt"
	"github.com/romiljainb/ecom/cart"
	//"github.com/romiljainb/ecom/user"
	//"github.com/romiljainb/ecom/fin"
)

func main() {


	r1 := cart.Review{UserReview:"Bruce Wayne", ID:1, Rating:3.5}
	r2 := cart.Review{UserReview:"Tony Stark", ID:2, Rating:4.5}
	r3 := cart.Review{UserReview:"Palpatine" , ID:3, Rating:1.5}
	r4 := cart.Review{UserReview:"Silas", ID:4, Rating:2.5}
	i1 := cart.Item{

		Name:"No to FB",
		ItemIDNum:1,
		UnitPrice:5,
		Quantity:5,
		Description:"He is Palpetine",
		Reviews: []cart.Review{r1,r2,r3,r4},
	}

	r5 := cart.Review{User:"DaVinci", UserReview:"stupid idiot", ID:4, Rating:2.5}
	r5.Replies = append(r5.Replies, r1)
	r5.Replies = append(r5.Replies, r2)
	r5.Replies = append(r5.Replies, r3)

	i1.InsertReview(r5)

	//fmt.Println(i1.Reviews[4].Replies[0].UserReview)
	i1.Reviews[4].DisplayReview()
	i1.UpdateTotalPrice()
	fmt.Println(i1.TotalPrice)
	i1.Reviews[4].DisplayAllReplies()

}
