package mlib

import (
	"fmt"
	"time"
)

type WavPlay struct {
	progress int
}
func (p *WavPlay) Play(source string) {
	fmt.Println("Wav music playing", source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(time.Millisecond * 100)
		p.progress += 10
	}
	fmt.Println("\nFinish Playing", source)
}
