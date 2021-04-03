package main

func main() {

	c := make(chan int)
	close(c)
	c <- 10
	//go func() {
	//	c <- 10
	//	close(c)
	//}()
	//v, ok := <-c
	//fmt.Printf("数据：%v %v\n", v, ok)
}
