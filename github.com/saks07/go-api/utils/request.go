package utils

func CheckGetMethod(requestMethod string) bool {
	return requestMethod == "GET"
}

func CheckPostMethod(requestMethod string) bool {
	return requestMethod == "POST"
}