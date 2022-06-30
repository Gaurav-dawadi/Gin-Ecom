package utils

import "fmt"

func IsSameUser(user1 int, user2 int) bool {
	return user1 == user2
}

func TypeAssertInt(arg interface{}) (int, bool) {
	fmt.Println("The arg is: ", arg)
	result, ok := arg.(int)
	fmt.Println(result, ok)
	return result, ok
}
