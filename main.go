package main

import (
	"context"
	"eight-stones/ecs-tank-engine/engine"
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/config"
	"fmt"
	"time"
)

func main() {
	cfg := config.Default()
	cfg.Game.Jobs.Recharger = time.Second
	gf := engine.New(cfg)

	uuid, err := gf.AddTank()
	if err != nil {
		panic(err)
	}
	_, err = gf.AddTank()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	gf.Start(ctx)
	fmt.Println(gf.Info())

	code := gf.Move(uuid, common.Right)
	fmt.Println(common.Translate(code))
	fmt.Println(gf.Info())
	time.Sleep(time.Millisecond * 500)

	code = gf.Move(uuid, common.Right)
	fmt.Println(common.Translate(code))
	fmt.Println(gf.Info())
	time.Sleep(time.Second)

	code = gf.Move(uuid, common.Right)
	fmt.Println(common.Translate(code))
	fmt.Println(gf.Info())

	cancel()

	time.Sleep(time.Second)

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
