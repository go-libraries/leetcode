package _202kth_node_from_end_of_list_lcci

import (
	"fmt"
	"testing"
)

func TestRecover(t *testing.T)  {
	//1 2 4 6 9 13 18 24 31 39 49 60 73
	// +1 +2 +2 + 2 + 3 + 4 +5 + 6 + 7 + 8 + 8 + 10 + 11 + 13
	fmt.Println(waysToChange(10))
}