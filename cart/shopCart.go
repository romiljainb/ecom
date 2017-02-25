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

func (s* shopCart) ShowItems(){ 	//current
	for _, i := range s.Items{
		i.DisplayItem()
	}
}

func (s* shopCart) VerifyPayments(){   //call within checkout

}

func (s* shopCart) Checkout(){

	s.VerifyPayments()

}




