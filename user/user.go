package user

import (
	//"fmt"
	"github.com/romiljainb/ecom/cart"
)

type Account struct {
	ID int
	Name string
	AccountType string  //admin vs seller vs buyer ... use const values
	GiftCardsVal	float32
	Addresses []Address
	Payments []PayMethod	//list of all payment options
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

/* pay stuff */
type PayMethod struct {
	ID int
	Name string
	Num int
	SecCode int
	AccountType string
	ExpDate string
}

type PaymentGateway interface {
	Authorize(p PayMethod) int
	Charge(authCode float32, auth int)
}

/*
type Card struct{	//credit,debit
	PayMethod
	ExpDate string
}


type PayAcct struct {	//venmo,paypal,venmo
	PayMethod
}
*/
/* order+track methods */

/* payment methods */

