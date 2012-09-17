package main

import ("fmt"; "time")

func main() {
    c1 := make(chan string)
    c2 := make(chan string)
	d  := make(chan bool, 1)

    go func() {
        time.Sleep(time.Millisecond * 1500)
		c1 <- "from 1"
    }()
    go func() {
        time.Sleep(time.Millisecond * 2000)
		c2 <- "from 2"
    }()

    go func() {
        for i :=0; i < 2; i ++ {
            select {
            case msg1 := <- c1:
                fmt.Println(msg1)
            case msg2 := <- c2:
                fmt.Println(msg2)
            }
        }
		d <- true
    }()
	<- d
}
