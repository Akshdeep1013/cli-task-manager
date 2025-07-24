package main

import (
	"flag"
)

func main() {
	//first define cl flags
	add := flag.String("add", "", "The task you want to add in your to-do list")
	list := flag.Bool("list", false, "List of all to-do list")
	delete := flag.Int("delete", -1, "Delete the given task id")
	complete := flag.Int("complete", -1, "Marked true for given task id")

	//parse flags
	flag.Parse()

	///

}
