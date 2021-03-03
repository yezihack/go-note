package main

import (
	"bytes"
	"encoding/gob"

	"github.com/mohae/deepcopy"
)

// 深拷贝
// 第三方库："github.com/mohae/deepcopy"

func DeepCopy(s, t interface{}) error {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if err := enc.Encode(s); err != nil {
		return err
	}
	if err := dec.Decode(t); err != nil {
		return err
	}
	return nil
}
func DeepCopyReflect(s, t interface{}) {
	t = deepcopy.Copy(s)
}
