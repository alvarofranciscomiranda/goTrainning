package simple

import (
	"fmt"
	"time"
)

func DatesFunc() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at: %s\n", t)

	n := time.Now()
	fmt.Printf("The time currently is %s\n", n)
	fmt.Printf("This object type is is %T\n", n)

	fmt.Println(n.Format(time.ANSIC))

	tomorrow := n.AddDate(0, 0, 1)
	fmt.Println(tomorrow.Format(time.ANSIC))

	format := "Mon 2006-02-01"
	fmt.Println(tomorrow.Format(format))
}
