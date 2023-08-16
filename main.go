package main

import (
	"eight-stones/ecs-tank-engine/components"
	"fmt"
	"strconv"
	"time"
)

func main() {
	gf := components.Field{
		Params: components.Params{
			MaxGamers: 4,
			FieldSize: 5,
		},
		Gamers: 0,
		Border: components.Border{
			X: 5,
			Y: 5,
		},
		Objects: nil,
	}

	uuid, err := gf.AddTank()
	if err != nil {
		panic(err)
	}

	gf.Start()

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println(strconv.FormatInt(int64(gf.Move(uuid, components.Right)), 2))
		}
	}()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(gf.Info())

	}

}
