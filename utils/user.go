package utils

func IsSameUser(user1 int, user2 int) bool {
	return user1 == user2
}

func TypeAssertInt(arg interface{}) (int, bool) {
	result, ok := arg.(int)
	return result, ok
}
