package main

// var x = "aa"

func main() {
	var i int
	i = 0
LOOP:
	for {
		println(i)
		switch {
		case i%2 == 1:
			break LOOP
		}
		i++
	}

}
