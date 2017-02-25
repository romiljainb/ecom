package user

import (
	//"fmt"
	"github.com/romiljainb/ecom/cart"
)

type Account struct {
	ID int
	Name string
	AccountType string  //admin vs seller vs buyer ... use const values
	Addresses []Address
	Payments []Pay	//list of all payment options
	Ordered		//embedded type to inherient

}

type Address struct{
	ID int
	Name string
	ZipCode int
}

type Ordered struct {
	Items map[cart.ItemID]Track	//map of current items and their tracking
	Previous []cart.Item
}

type Track struct {
	ID int
	ETA int
	Current string  //current location
}

////
type Pay struct {
	ID int
	Num int


}

type Payment interface {
	Verify()
}

/*type Account struct {
	ID int
	Name string
	Balance float32
	AccountType string
}
*/
type Card struct{
	ID int
	Name string

}
////

//gift card, paypal, card, other types of account,paytm, venmo

//order tracking

//previous orders