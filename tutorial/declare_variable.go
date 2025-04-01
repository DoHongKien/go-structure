package tutorial

import "fmt"

func Init() {

	var myName string = "Đỗ Hồng Kiên"

	const VERSION = "0.0.1"

	err3 := 0

	arr := []string{"A", "B", "C", "D", "E", "F"}

	map1 := map[int]string{
		1: "kiendh",
		2: "mybui",
	}

	switchValue := 1

	switch switchValue {
	case 1:
	case 2:
	case 3:
	}

	fmt.Printf("My name is: %v\n", myName)
	fmt.Printf("Err3 %v\n", err3)
	fmt.Printf("Arr %v\n", arr)

	for key, value := range map1 {
		fmt.Printf("Key: %v | Value: %v", key, value)
	}

}
