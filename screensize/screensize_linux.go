// full implementation see https://github.com/fd0/grobi/blob/master/randr.go

// +build linux

package screensize

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	width  int
	height int
	err    error
)

//XrandrReader does...
func XrandrReader() *bytes.Reader {
	cmd := exec.Command("xrandr") // | grep ' connected' | cut -d' ' -f4 | grep -Po '[0-9]{4,5}.[0-9]{3}'")
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error: ", err)
	}
	return bytes.NewReader(output)
}

//Get does
func Get(priority string) (int, int) {

	//
	// xrandr | grep ' connected' | cut -d' ' -f4 | grep -Po '[0-9]{4,5}.[0-9]{3}'
	width = 1024
	height = 800
	err = nil

	ls := bufio.NewScanner(XrandrReader())
	for ls.Scan() {
		line := ls.Text()
		if strings.Contains(line, " connected ") {
			s := strings.Split(line, " ")
			s2 := strings.Split(s[3], "+")
			s3 := strings.Split(s2[0], "x")
			//fmt.Println(s3[0], s3[1])
			width, _ = strconv.Atoi(s3[0])
			height, _ = strconv.Atoi(s3[1])
		}
	}

	//fmt.Println("screensize.get linux")
	return width, height
}
