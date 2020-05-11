package mlib

import (
	"fmt"
	"time"
)

type Mp3Play struct {
	progress int
}
func (p *Mp3Play) Play(source string) {
	fmt.Println("Playing MP3 music", source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(time.Millisecond * 100)
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}