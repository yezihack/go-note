package mlib

import "code.qschou.com/golang/errors"

type MusicEntry struct {
	Id string
	Name string
	Artist string
	Source string
	Type string
}

type MusicManager struct {
	musics []MusicEntry
}
func NewMusicManager() *MusicManager {
	return &MusicManager{
		musics:make([]MusicEntry, 0),
	}
}
func (m *MusicManager) Len() int {
	return len(m.musics)
}
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index > m.Len() {
		return nil, errors.New("index out of range")
	}
	return &m.musics[index], nil
}
func (m *MusicManager) Find(name string) *MusicEntry {
	if m.Len() == 0 || name == "" {
		return nil
	}
	for _, row := range m.musics {
		if row.Name == name {
			return &row
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index > m.Len() {
		return nil
	}
	removeMusic := m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return &removeMusic
}

