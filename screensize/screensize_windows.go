// +build windows

package screensize

import (
	"github.com/lxn/win"
)

//Get does
func Get() (int, int, error) {
	hDC := win.GetDC(0)
	defer win.ReleaseDC(0, hDC)
	width := int(win.GetDeviceCaps(hDC, win.HORZRES))
	height := int(win.GetDeviceCaps(hDC, win.VERTRES))
	// fmt.Printf("%dx%d\n", width, height)
	return width, height, nil
}
