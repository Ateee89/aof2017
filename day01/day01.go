package day01

import (
	"bytes"
	"fmt"
)

func main()  {

}

func Solve01(input []byte) {
	bs := bytes.TrimSpace(input)

	sum := 0
	next := bs[0]
	for i := len(bs) - 1; i >= 0; i-- {
		if bs[i] == next {
			sum += int(next - '0')
		}
		next = bs[i]
	}

	sum2 := 0
	for i := 0; len(bs)/2-1 >= i; i++ {
		if bs[i] == bs[i+len(bs)/2] {
			sum2 += int(bs[i]-'0') + int(bs[i+len(bs)/2]-'0')
		}
	}
	fmt.Printf("Solution for the first part is %v\n", sum)
	fmt.Printf("Solution for the second part is %v\n", sum2)
}


