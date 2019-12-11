package mlib

import "fmt"

type Player interface {
	Play(source string)
}
func Play(source, mtype string) {
	var p Player
	switch mtype {
	case "mp3":
		p = new(Mp3Play)
	case "wav":
		p = new(WavPlay)
	default:
		fmt.Println("Unsupported music type", source)
	}
	p.Play(source)
}

