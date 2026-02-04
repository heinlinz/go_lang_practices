package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["Answers"] = "Incorrect"
	fmt.Println("The value is ", m["Answers"])

	m["Dogs"] = "Leo"
	fmt.Println("The name of the dog is ", m["Dogs"])

	delete(m, "Answers")
	fmt.Println(m["Answers"])

	v, ok := m["Answers"]
	if !ok {
		fmt.Println("doesn't exist")
	} else {
		fmt.Println(v)
	}
}

