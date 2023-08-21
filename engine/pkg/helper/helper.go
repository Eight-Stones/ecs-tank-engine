package helper

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
)

var clearFunc = map[string]func(){
	"linux": func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
	"darwin": func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
	"windows": func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
}

func clear() {
	value, ok := clearFunc[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                              //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

type position struct {
	X         int
	Y         int
	Direction uint
}

type info struct {
	Id        string
	Kind      string
	Position  position
	HitPoints int
	Damage    int
}

func (i info) String() string {
	return fmt.Sprintf(
		"| \t %v \t | \t %v \t | \t %d \t | \t [x:%d y:%d] \t | \t %v \t |",
		i.Id,
		i.Kind,
		i.HitPoints,
		i.Position.X,
		i.Position.Y,
		i.Position.Direction,
	)
}

// DrawField отрисовывает в консоли игровые данные.
//
// TODO: переосмыслить вывод представления игрового поля и общий вывод в поток, а не принтами.
func DrawField(x, y int, in map[string]map[string]interface{}) {
	objects := make([]info, 0, len(in))
	for key, value := range in {
		coordinates := value[common.KeyPositionCoordinate].([]int)
		object := info{
			Id:   key,
			Kind: value[common.KeyObjectKind].(string),
			Position: position{
				X:         coordinates[0],
				Y:         coordinates[1],
				Direction: value[common.KeyMovementDirection].(uint),
			},
			HitPoints: value[common.KeyStatHitPoints].(int),
			Damage:    value[common.KeyStatDamage].(int),
		}
		objects = append(objects, object)
	}

	sort.SliceStable(objects, func(i, j int) bool {
		return objects[i].Id < objects[j].Id
	})
	clear()

	view := make([][]string, y)
	for idx := range view {
		view[idx] = make([]string, x)
	}

	for idx := range view {
		for idy := range view[idx] {
			view[idx][idy] = "."
		}
	}

	fmt.Printf("****************\n* PLAYERS INFO *\n****************\n")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("|\t ID \t\t\t\t\t | \t KIND \t | \t HP \t | \t COORDINATES \t | \t DIR \t |")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------")
	for i := range objects {
		fmt.Println(objects[i])
		view[objects[i].Position.Y][objects[i].Position.X] = "T"
		fmt.Println("--------------------------------------------------------------------------------------------------------------------------")
	}

	fmt.Printf("**************\n* GAME FIELD *\n**************\n")

	for idy := len(view) - 1; idy >= 0; idy-- {
		for idx := 0; idx < len(view[idy]); idx++ {
			fmt.Print(view[idy][idx] + "  ")
		}
		fmt.Printf("  %d\n", idy)
	}
	str := ""
	for i := 0; i < len(view); i++ {
		if len(strconv.Itoa(i)) == 2 {
			str += strconv.Itoa(i) + " "
			continue
		}
		str += strconv.Itoa(i) + "  "
	}
	fmt.Println(str)
}
