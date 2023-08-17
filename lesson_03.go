package main

import (
	"fmt";
);

func main() {
	// 条件分岐処理

	age := 22;

	if age >= 18 {
		fmt.Println("成年");
	} else {
		fmt.Println("未成年");
	}

	switch {
	case age >= 18:
		fmt.Println("成年")
	default:
		fmt.Println("未成年")
	}
}