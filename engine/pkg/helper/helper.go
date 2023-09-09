package helper

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
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
	Icon      string
	Kind      string
	Position  position
	HitPoints int
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

func drawHeader(objects []info) {
	fmt.Printf("****************\n* PLAYERS INFO *\n****************\n")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("|\t ID \t\t\t\t\t | \t KIND \t | \t HP \t | \t COORDINATES \t | \t DIR \t |")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------")
	for i := range objects {
		fmt.Println(objects[i])
		fmt.Println("--------------------------------------------------------------------------------------------------------------------------")
	}
}

func drawMap(view [][]string) {
	fmt.Printf("**************\n* GAME FIELD *\n**************\n")
	for idy := len(view) - 1; idy >= 0; idy-- {
		for idx := 0; idx < len(view[idy]); idx++ {
			fmt.Print(view[idy][idx] + "  ")
		}
		fmt.Printf("  %d\n", idy)
	}

	for i := 0; i < len(view); i++ {
		if len(strconv.Itoa(i)) == 2 {
			fmt.Print(strconv.Itoa(i) + " ")
			continue
		}
		fmt.Print(strconv.Itoa(i) + "  ")
	}
	fmt.Println()
}

func prepareObjects(in map[string]map[string]interface{}) []info {
	objects := make([]info, 0, len(in))
	for key, value := range in {
		icon := ""
		switch value[common.KeyStatObjectKind].(string) {
		case common.KeyObjectKindTank:
			icon = "Ⓣ"
		case common.KeyObjectKindBullet:
			icon = "⍟"
		}

		coordinates := value[common.KeyStatPositionCoordinate].([]int)
		object := info{
			Id:   key,
			Icon: icon,
			Kind: value[common.KeyStatObjectKind].(string),
			Position: position{
				X:         coordinates[0],
				Y:         coordinates[1],
				Direction: value[common.KeyStatMovementDirection].(uint),
			},
			HitPoints: value[common.KeyStatHitPoints].(int),
		}
		objects = append(objects, object)
	}

	sort.SliceStable(objects, func(i, j int) bool {
		return objects[i].Id < objects[j].Id
	})

	return objects
}

func prepareView(x, y int, objects []info) [][]string {
	view := make([][]string, y)
	for idx := range view {
		view[idx] = make([]string, x)
	}

	for idx := range view {
		for idy := range view[idx] {
			view[idx][idy] = "."
		}
	}

	for idx := range objects {
		view[objects[idx].Position.Y][objects[idx].Position.X] = objects[idx].Icon
	}

	return view
}

// DrawField TODO: переосмыслить вывод представления игрового поля и общий вывод в поток, а не принтами.
func DrawField(x, y int, in map[string]map[string]interface{}) {
	objects := prepareObjects(in)
	view := prepareView(x, y, objects)
	clear()
	drawHeader(objects)
	drawMap(view)
}

func DrawResult(in map[string]map[string]interface{}) {
	clear()
	builder := strings.Builder{}
	for key, value := range in {
		builder.WriteString(fmt.Sprintf("id:%v \t", key))
		builder.WriteString(fmt.Sprintf("last position:%v \t", value[common.KeyStatPositionCoordinate].([]int)))
		builder.WriteString(fmt.Sprintf("last HP:%v", value[common.KeyStatHitPoints].(int)))
		builder.WriteString("actions:\n")
		for idx, action := range value[common.KeyStatActions].([]int) {
			builder.WriteString(fmt.Sprintf("%d:", idx))
			common.TranslateBuilder(action, &builder)
			builder.WriteString("\n")
		}
		builder.WriteString("\n")
	}
	fmt.Println(builder.String())
}
