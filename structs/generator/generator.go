package generator

func IDGenerator() func() int {
	id := 1
	return func() int {
		current := id
		id++
		return current
	}
} //замыкание

var NewID = IDGenerator()
var NextOrderID = IDGenerator()
