package user

import (
	//"fmt"
	"eComm/mylib"
)

type Person struct {
	ID int
	Name string
	AccountType string
	Addresses []Address
	Payments []Pay
	Ordered		//embedded type to inherient

}

type Address struct{
	ID int
	Name string
	ZipCode int
}

type Ordered struct {
	Items map[cart.Item]Track	//map of items and their tracking
	Previous []cart.Item
}

type Track struct {
	Num int
	ETA int
	Current string
}

type Pay struct {
	ID int
	Num int


}


//gift card, paypal, card, other types of account,paytm, venmo

//order tracking

//previous orders