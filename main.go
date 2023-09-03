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
	_, err = gf.AddTank()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	gf.Start(ctx)
	go gf.DrawConsole(ctx)

	time.Sleep(time.Second * 1)
	gf.Shoot(uuid1)

	time.Sleep(time.Second * 1)
	gf.Shoot(uuid1)

	time.Sleep(time.Second * 1)
	gf.Shoot(uuid1)

	for i := 0; i <= 15; i++ {
		time.Sleep(time.Second * 1)
	}

	cancel()

	time.Sleep(time.Second)

	gf.DrawResult()
}
