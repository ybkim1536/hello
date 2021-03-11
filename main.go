package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ybkim1536/greeting"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	message := greeting.Hello("Gladys")
	fmt.Println(message, rand.Intn(10))
	message1 := greeting.Goodby("Gladys")
	fmt.Println(message1, rand.Intn(10))

}
