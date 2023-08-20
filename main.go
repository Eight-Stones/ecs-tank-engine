package main

import (
	"eight-stones/ecs-tank-engine/engine"
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/config"
	"fmt"
)

func main() {
	cfg := config.Default()
	gf := engine.New(cfg)

	uuid, err := gf.AddTank()
	if err != nil {
		panic(err)
	}
	_, err = gf.AddTank()
	if err != nil {
		panic(err)
	}

	gf.Start()
	fmt.Println(gf.Info())
	gf.Move(uuid, common.Right)
	fmt.Println(gf.Info())

	/*go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println(strconv.FormatInt(int64(gf.OkStep(uuid, common.Right)), 2))
		}
	}()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(gf.Info())

	}*/

}
