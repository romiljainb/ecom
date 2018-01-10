package fin

import (
	//"fmt"
	"github.com/romiljainb/ecom/user"
	"time"
)


type Bank struct {
	ID int
	Name string
	Accounts []BankAccount

}

type BankAccount struct {
	ID                 int
	Name               string
	AcctType           string
	Balance            float32
	Interest           float32
	PendingWithdrawals float32
	PendingDeposit     float32
	Transactions	[]Transaction

}

type Transaction struct {
	ID          int
	Name        string
	PostedDate  time.Time //date/time posted on
	Description string
	TransactionType string
	Amount float32

}

func (b* Bank) Authorize(p user.PayMethod) int{
	AuthCode := 0

	/* Implementaion coming */

	return AuthCode
}

/* methods TODO:
transfer
withdraw
deposit
*/
