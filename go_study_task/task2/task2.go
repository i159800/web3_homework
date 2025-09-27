package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 指针相关
func Pointer1(numPtr *int) {
	*numPtr += 10
}

func Pointer2(slicePtr *[]int) {
	//使用解引用操作符 * 来访问切片指针指向的切片
	slice := *slicePtr
	for i := range slice {
		slice[i] *= 2
	}
}

// goroutine相关
func GoRoutine1() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Println("odd ", i)
			}
		}
	}()

	go func() {
		for j := 1; j <= 10; j++ {
			if j%2 == 0 {
				fmt.Println("even ", j)
			}
		}
	}()
	time.Sleep(time.Second)
}

var count uint32

func GoRoutine2() {
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {})
}

var trigger = func(i uint32, fn func()) {
	for {
		if n := atomic.LoadUint32(&count); n == i {
			fn()
			atomic.AddUint32(&count, 1)
			break
		}
		time.Sleep(time.Nanosecond)
	}
}

var num = 10
var chan1 = make(chan struct{}, num)

func GoRoutine3() {
	for i := 0; i < num; i++ {
		go func(i int) {
			fmt.Println(i)
			chan1 <- struct{}{}
		}(i)
	}
	for j := 0; j < num; j++ {
		<-chan1
	}
}

// 面向对象
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	width, length float64
}

func (rec Rectangle) Area() float64 {
	return rec.width * rec.length
}
func (rec Rectangle) Perimeter() float64 {
	return 2 * (rec.width + rec.length)
}

type Circle struct {
	radius float64
}

func (cir Circle) Area() float64 {
	return 3.14 * cir.radius * cir.radius
}

func (cir Circle) Perimeter() float64 {
	return 2 * 3.14 * cir.radius
}

func Area(s Shape) float64 {
	return s.Area()
}

func Perimeter(s Shape) float64 {
	return s.Perimeter()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Person Info name:%s age:%d", p.Name, p.Age)
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("Employee Info %s EmployeeID:%d", e.Person, e.EmployeeID)
}

// Channel
func Channel1() {
	var chan1 = make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			chan1 <- i
		}
	}()

	for i := 1; i <= 10; i++ {
		n := <-chan1
		fmt.Println("chan1 ", n)
	}
}

func Channel2() {
	num := 100
	var chan2 = make(chan int, 10)
	go func() {
		for i := 1; i <= num; i++ {
			chan2 <- i
			fmt.Println("send num: ", i)
		}
	}()

	for i := 1; i <= num; i++ {
		fmt.Println("receive num: ", <-chan2)
	}
}

// 锁机制
func MutexCounter() {
	var counter = uint32(0)
	var mu sync.Mutex
	num := 1000
	goNum := 10
	sign := make(chan struct{}, goNum)
	for i := 1; i <= goNum; i++ {
		go func() {
			defer func() {
				sign <- struct{}{}
			}()
			mu.Lock()
			for j := 0; j < num; j++ {
				counter++
			}
			mu.Unlock()
		}()
	}

	for i := 0; i < goNum; i++ {
		<-sign
	}
	fmt.Println("final counter ", counter)
}

func AtomicCounter() {
	var counter uint32
	var wg sync.WaitGroup
	goNum := 10
	numPerGoroutine := 1000
	wg.Add(goNum)
	for i := 0; i < goNum; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numPerGoroutine; j++ {
				atomic.AddUint32(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("final counter ", counter)

}

func main() {
	num1 := 10
	Pointer1(&num1)
	fmt.Println("num1 after ", num1)
	nums := []int{1, 2, 3}
	numsPtr := &nums
	Pointer2(numsPtr)
	fmt.Println("nums after ", nums)
	GoRoutine1()
	GoRoutine2()
	GoRoutine3()
	var rec1 = Rectangle{1.0, 2.0}
	var cir1 = Circle{3.0}
	var shape1, shape2 Shape
	shape1 = rec1
	shape2 = cir1
	fmt.Println("rec_area: ", shape1.Area(), " rec_perimeter: ", shape1.Perimeter())
	fmt.Println("cir_area: ", shape2.Area(), " cir_perimeter: ", shape2.Perimeter())
	var person = Person{Name: "zhangsan", Age: 20}
	var employee = Employee{Person: person, EmployeeID: 10001}
	employee.PrintInfo()
	Channel1()
	Channel2()
	MutexCounter()
	AtomicCounter()
}
