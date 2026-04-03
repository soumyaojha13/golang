//creating a map

package main

import "fmt"

func main() {

	m := make(map[string]string)

	m["name"] = "go lang"
	m["area"] = "backend"

	fmt.Println(m["name"])
	fmt.Println(m["area"])

	//making map with values

	m1 := map[string]int{"price": 40, "phones": 10}

	v, ok := m["phones"]
	fmt.Println(v)

	if ok {
		fmt.Println("all ok")
	} else {

		fmt.Println(("not ok"))
	}

}
