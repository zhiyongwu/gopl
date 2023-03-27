package main

import (
	"fmt"
	"github.com/sony/sonyflake"
)

func main() {
	s := "hello 你"
	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(i)

	}
	sn := sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return 1, nil
		},
	})
	for i := 0; i < 100; i++ {
		fmt.Println(sn.NextID())
		fmt.Println(sn.NextID())
		fmt.Println(sn.NextID())
	}
}
