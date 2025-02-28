package main

import "fmt"

type Blacklist func(string) bool

func regiterUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Printf("You are blocked, %c***%c\n", name[0], name[len(name)-1])
	} else {
		fmt.Println("Welcome, ", name)
	}
}

func user() string {
	var user string
	fmt.Println("Input User: ")
	fmt.Scanln(&user)
	return user
}

func main() {
	// blacklist := func(name string) bool {
	// 	return name == "anjing"
	// }
	// regiterUser(user(), blacklist)

	// Or
	regiterUser(user(), func(name string) bool {
		blockedNames := []string{"anjing", "fuck", "bitch"}
		for _, blockedName := range blockedNames {
			if name == blockedName {
				return true
			}
		}
		return false
	})
}
