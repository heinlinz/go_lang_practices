package main

import (
	"fmt"
)

func NewRateLimiter(limit int) (func(string) bool, func(string)) {
	counter := make(map[string]int)

	limiter := func(username string) bool {
		counter[username]++
		return counter[username] <= limit
	}

	reset := func(username string) {
		fmt.Printf("Resetting the user request limit for %s...\n", username)
		delete(counter, username)
	}

	return limiter, reset

}

func main() {
	check, clear := NewRateLimiter(2)

	users := []string{"user1", "user1", "user1", "user2"}

	for _, user := range users {
		result := check(user)
		fmt.Println(result)
		if !result {
			clear(user)
		}
	}
}
