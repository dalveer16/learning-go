package CH1

import "fmt"

func L2() {
	smsSendingLimit := 0
	costPerSMS := 0.0
	hasPermission := false
	username := ""

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
