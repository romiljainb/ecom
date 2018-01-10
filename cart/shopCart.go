package cart

import (
	//"github.com/romiljainb/ecom/user"
	//"github.com/romiljainb/ecom/fin"
)

type shopCart struct {
	ID int
	Name string
	TotalItems int
	Items	[]Item
	DiscountItems map[ItemID]float32
}

func (s* shopCart) InsertItem(i* Item){
	s.InsertItem(i)
}

func (s* shopCart) RemoveItem(i* Item){
	s.RemoveItem(i)
}

func (s* shopCart) ShowItems(){ 	//shows all current items in cart
	for _, i := range s.Items{
		i.DisplayItem()
	}
}

func (s* shopCart) VerifyPayments(){   //call within checkout
	/* TODO soon */
}

func (s* shopCart) Checkout(){

	s.VerifyPayments()

}

/* Add catalog struct and functions for sellers to populate and for buyers to select from */




