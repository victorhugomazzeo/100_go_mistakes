// Mistake #13: Creating utility packages
//
// Names like util, common or shared say nothing: name a package after what it
// provides (stringset), so calls read naturally, like stringset.New().
package main

import (
	"fmt"

	"github.com/victorhugomazzeo/100_go_mistakes/013-creating-utility-packages/stringset"
	"github.com/victorhugomazzeo/100_go_mistakes/013-creating-utility-packages/util"
)

func main() {
	mistake()
	fix()
}

func mistake() {
	values := []string{"banana", "apple", "orange"}
	set := util.NewStringSet(values...)
	fmt.Println("mistake:", util.SortStringSet(set))
}

func fix() {
	values := []string{"banana", "apple", "orange"}
	set := stringset.New(values...)
	fmt.Println("fix:", stringset.Sort(set))
}
