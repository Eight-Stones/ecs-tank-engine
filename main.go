package main

import (
	"context"
	"eight-stones/ecs-tank-engine/engine"
	"eight-stones/ecs-tank-engine/engine/config"
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
	go gf.DrawConsole(ctx)

	gf.Shoot(uuid)

	for i := 0; i <= 20; i++ {
		time.Sleep(time.Second * 1)
		//gf.Move(uuid, common.Right)
	}

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
