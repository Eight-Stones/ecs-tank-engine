package main

import (
	"eight-stones/ecs-tank-engine/engine"
	"eight-stones/ecs-tank-engine/engine/common"
	"fmt"
	"strconv"
	"time"
)

func main() {
	gf := engine.Field{
		Params: engine.Params{
			MaxGamers: 4,
			FieldSize: 5,
		},
		Gamers: 0,
		Border: engine.Border{
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
			fmt.Println(strconv.FormatInt(int64(gf.Move(uuid, common.Right)), 2))
		}
	}()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(gf.Info())

	}

}
