package dogs

func Bark() string {
	return "Woof!!"
}

func BarkAt(name string) string {
	return "I am barking at" + name + Bark()
}
