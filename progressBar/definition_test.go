package progressBar

import (
	"fmt"
	"testing"
	"time"
)

func TestProgressBar(t *testing.T) {
	fmt.Println("this is a progress bar")
	var title = "progress bar"
	var current = 12
	var total = 100
	var unit = "Mib"
	var newBar = NewProgressBar(title, current, total)
	newBar.SetUnit(unit)
	newBar.SetGraph(">")
	for i := current; i <= total; i++ {
		newBar.Run(i)
		time.Sleep(time.Second / 100)
	}
}
