package Patterm

import "testing"

func TestNewSDCard(t *testing.T) {
	var computer ComputerI
	computer = NewComputer()
	sd := NewSDCard()
	sd.WriteSD("智零小姐姐")
	computer.ReadSD(sd)

	tf := NewTFCard()
	tf.WriteTF("我是一张旧照片")
	ad := NewSDAdapterTF(tf)
	ad.ReadSD()
}
