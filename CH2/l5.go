package CH2

import "fmt"

func L5() {

	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64 = bulkMessageCost
	if isPremiumUser {
		finalCost = finalCost - finalCost*discountRate
	}
	if accountBalance-finalCost >= 0 {
		accountBalance = accountBalance - finalCost
		fmt.Println(purchaseSuccessMessage)
	} else {
		fmt.Println(insufficientFundMessage)
	}

	// don't edit above this line

	// ?

	// don't edit below this line

	fmt.Println("Account balance:", accountBalance)
}
