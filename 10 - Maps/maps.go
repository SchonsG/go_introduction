package main

import "fmt"

func main() {
	user := map[string]string {
		"name": "Guilherme",
		"last_name": "Schons",
	}
	fmt.Println(user["name"])

	userTwo := map[string]map[string]string {
		"personal": {
			"name": "Guilherme",
			"last_name": "Schons",
		},
		"course": {
			"name": "Sistemas de Informação",
			"period": "Seventh",
		},
	}

	fmt.Println(userTwo["personal"]["last_name"])
	fmt.Println(userTwo["course"]["name"])

	// Delete key
	fmt.Println(user)
	delete(user, "last_name")
	fmt.Println(user)
	user["last_name"] = "Schons"
	fmt.Println(user)
}