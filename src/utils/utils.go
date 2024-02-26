/*
General purpose utility functions.
*/

package utils

import (
	"fmt"
)

func HandleErr(e error, msg string) {
	if e != nil {
		fmt.Printf("ERROR: %s\n%s\n", msg, e)
	}
}