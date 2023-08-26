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

	uuid1, err := gf.AddTank()
	if err != nil {
		panic(err)
	}
	uuid2, err := gf.AddTank()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	gf.Start(ctx)
	go gf.DrawConsole(ctx)

	time.Sleep(time.Second * 1)
	gf.Shoot(uuid1)
	gf.Shoot(uuid2)

	time.Sleep(time.Second * 2)
	gf.Shoot(uuid1)

	for i := 0; i <= 30; i++ {
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
