package main

import "fmt"
import "reflect"
import "sync"
import "time"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int)(int,int) {
	return num1/num2, num1 % num2
}

func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("some error happened")
			ret = -1
		}
	}()
	arr := [3]int{2,3,4}
	return arr[index]
}

type Student struct {
	name string
	age int
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, i am %s",person, stu.name)
}

var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download",url)
	time.Sleep(time.Second)
	wg.Done()
}

var ch = make(chan string,10)
func download2(url string) {
	fmt.Println("start to download2",url)
	time.Sleep(time.Second)
	ch <- url
}

func main() {
	fmt.Println("Hello World!")
	str1 := "go语言"
	fmt.Printf("%d %c\n",str1[2],str1[2]);
	fmt.Println(reflect.TypeOf(str1[2]).Kind())

	arr := [5]int{1,2,3,4,5}
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	fmt.Println(arr)
	
	type Gender int8
	const (
		MALE Gender = 1
		FEMALE Gender = 2
	)
	gender := MALE

	switch gender {
	case MALE:
		fmt.Println("male")
	case FEMALE:
		fmt.Println("female")
	default:
		fmt.Println("unknown")
	}

	quo, rem := div(100,17)
	fmt.Println(quo,rem)
	fmt.Println(add(100,17))

	fmt.Println(get(5))
	fmt.Println("finished")

	stu := &Student{
		name : "edward",
	}
	msg := stu.hello("Jack")
	fmt.Println(msg)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go download("a.com/" + string(i+'0'))
	}
	wg.Wait()
	fmt.Println("Done")

	for i := 0; i < 4; i++ {
		go download2("b.com/" + string(i+'0'))
	}
	for i := 0; i < 4; i++ {
		msg := <-ch
		fmt.Println("finish",msg)
	}
	fmt.Println("Done!")
}
