package main

import "fmt"

func func1() {
	i := 0
	defer fmt.Println(i) //输出0
	i++
	defer fmt.Println(i) //输出1
	return
}

func func2() (val int) {
	val = 10
	defer func() {
		fmt.Printf("defer里面 val:%d\n", val)
		val += 2
	}()
	val++
	return 100
}
func printFunc2() {
	fmt.Println(func2())
}
func func3() {
	i := 0
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
}
func func4() (val int) {
	val = 0
	defer func() {
		fmt.Printf("func3 before val += 5:%d\n", val)
	}()
	val += 5
	defer func() {
		fmt.Printf("func3 after val += 5:%d\n", val)
	}()
	return val
}
func printFunc4() {
	fmt.Println(func4())
}
func func5() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
func deferFuncReturn() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
}
func foo() int {
	var i int
	defer func() {
		i++
		fmt.Println(i)
	}()
	return i
}

func main() {
	//func1()
	//printFunc2()
	//func3()
	//printFunc4()
	//func5()
	fmt.Println(foo())
}

/*//1.释放资源
m.mutex.Lock()
defer m.mutex.Unlock()*/

/*//2.流程控制
var wg wait.Group
defer wg.Wait()*/

//3.异常处理
//defer func() { recover() } ()
