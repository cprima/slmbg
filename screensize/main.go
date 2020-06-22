package screensize

import (
	"fmt"
)

//Monitor holds information about a single monitor
//type Monitor interface {
//	area() int
//}

func (r Monitor) area() int {
	return r.Width * r.Height
}

//Monitor holds information about a single monitor
type Monitor struct {
	Width, Height int
	priority      string
	Name          string
}

//NewMonitor returns...
func NewMonitor(priority string) Monitor {
	return Monitor{Width: 10, Height: 10, priority: priority}
}

//Monitors holds one or more monitor structs
type Monitors struct {
	monitor map[string]Monitor
}

func init() {
}

func main() {}

//GetMonitors returns all connected monitors
func GetMonitors() {
	fmt.Println("GetMonitors")
	mp := NewMonitor("primary")
	mp.Width, mp.Height, _ = Get("primary")
	fmt.Println("monitors:", mp)
	fmt.Println("monitor area:", mp.area())
}
