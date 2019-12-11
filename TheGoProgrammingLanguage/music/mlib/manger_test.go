package mlib

import (
	"github.com/yezihack/go-study/TheGoProgrammingLanguage/src"
	"testing"
)

func TestNewMusicManager(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("music manger failed")
	}
	src.Asset(2, t, 0, mm.Len())
	mm.Add(&MusicEntry{
		Id:"1",
		Name:"yellow",
		Source:"163",
		Artist:"uas",
		Type:"jk",
	})
	src.Asset(3, t, 1, mm.Len())
	src.Asset(4, t, "163", mm.Find("yellow").Source)
	g, err := mm.Get(0)
	src.Asset(5, t, nil, err)
	src.Asset(6, t, "jk", g.Type)
	src.Asset(7, t, "uas", mm.Remove(0).Artist)
	src.Asset(8, t, 0, mm.Len())
}
