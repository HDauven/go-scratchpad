package hof

import "fmt"

func ExampleTimes() {
	double := Times(2)
	fmt.Println(double(5))
	fmt.Println(double(-5))
	negativeDouble := Times(-2)
	fmt.Println(negativeDouble(5))
	fmt.Println(negativeDouble(-5))
	triple := Times(3)
	fmt.Println(triple(5))
	fmt.Println(triple(-5))
	// Output:
	// 10
	// -10
	// -10
	// 10
	// 15
	// -15
}
