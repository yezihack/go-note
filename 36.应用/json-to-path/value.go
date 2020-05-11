package json_to_path

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/spf13/cast"
)

const DEFAULT_DELIM = "."

type Value interface {
	GetData() *map[string]interface{}

	String(key string) string
	Strings(key string) []string
	Int(key string) int

	StringE(key string) (string, error)
	StringsE(key string) ([]string, error)
	IntE(key string) (int, error)
	//Int64E(key string) (int64, error)
	//BoolE(key string) (bool, error)
	//FloatE(key string) (float64, error)
	Scan(val interface{}) error
	SetDelim(delim string)
	Load(val []byte) error
}

type kv struct {
	Data     map[string]interface{}
	mu       sync.RWMutex
	keyDelim string
}

func NewKv() Value {
	return &kv{
		Data:     make(map[string]interface{}),
		keyDelim: DEFAULT_DELIM,
	}
}

func (k *kv) SetDelim(delim string) {
	k.keyDelim = delim
}

func (k *kv) Load(val []byte) error {
	if err := json.Unmarshal(val, &k.Data); err != nil {
		return err
	}
	return nil
}

func (k *kv) Scan(val interface{}) error {
	buf, err := json.Marshal(k.GetData())
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, val)
	if err != nil {
		return err
	}
	return nil
}

func (k *kv) GetData() *map[string]interface{} {
	k.mu.RLock()
	data := &k.Data
	k.mu.RUnlock()
	return data
}

//func (k *kv) get(key string) interface{} {
//	k.mu.RLock()
//	value := k.Data[key]
//	k.mu.RUnlock()
//	return value
//}

func (k *kv) String(key string) string {
	return cast.ToString(k.get(key))
}

func (k *kv) Strings(key string) []string {
	return cast.ToStringSlice(k.get(key))
}

func (k *kv) Int(key string) int {
	return cast.ToInt(k.get(key))
}

func (k *kv) StringE(key string) (string, error) {
	return cast.ToStringE(k.get(key))
}

func (k *kv) StringsE(key string) ([]string, error) {
	return cast.ToStringSliceE(k.get(key))
}

func (k *kv) IntE(key string) (int, error) {
	return cast.ToIntE(k.get(key))
}

// 有远程配置中心， 下来配置到 --> config  保存到本地
// 没有远程配置中心 读取本地配置 --> config

// Get 将按照 flag  > config > defaults
func (v *kv) get(key string) interface{} {
	val := v.find(key)
	if val == nil {
		return nil
	}

	return val
}

func (v *kv) find(lcaseKey string) interface{} {
	var (
		val    interface{}
		path   = strings.Split(lcaseKey, v.keyDelim)
		nested = len(path) > 1
	)

	// Config file next
	val = v.searchMapWithPathPrefixes(*v.GetData(), path)
	if val != nil {
		return val
	}
	if nested && v.isPathShadowedInDeepMap(path, *v.GetData()) != "" {
		return nil
	}

	//// Default next
	//val = v.searchMap(v.defaults, path)
	//if val != nil {
	//	return val
	//}
	//if nested && v.isPathShadowedInDeepMap(path, v.defaults) != "" {
	//	return nil
	//}
	return nil

}

// searchMapWithPathPrefixes recursively searches for a value for path in source map.
//
// While searchMap() considers each path element as a single map key, this
// function searches for, and prioritizes, merged path elements.
// e.g., if in the source, "foo" is defined with a sub-key "bar", and "foo.bar"
// is also defined, this latter value is returned for path ["foo", "bar"].
//
// This should be useful only at config level (other maps may not contain dots
// in their keys).
//
// Note: This assumes that the path entries and map keys are lower cased.
func (v *kv) searchMapWithPathPrefixes(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}

	// search for path prefixes, starting from the longest one
	for i := len(path); i > 0; i-- {
		prefixKey := strings.Join(path[0:i], v.keyDelim)

		next, ok := source[prefixKey]
		if ok {
			// Fast path
			if i == len(path) {
				return next
			}

			// Nested case
			var val interface{}
			switch nextValue := next.(type) {
			case map[interface{}]interface{}:
				val = v.searchMapWithPathPrefixes(cast.ToStringMap(next), path[i:])
			case map[string]interface{}:
				// Type assertion is safe here since it is only reached
				// if the type of `next` is the same as the type being asserted
				val = v.searchMapWithPathPrefixes(nextValue, path[i:])
			default:
				// got a value but nested key expected, do nothing and look for next prefix
			}
			if val != nil {
				return val
			}
		}
	}

	// not found
	return nil
}

// isPathShadowedInDeepMap makes sure the given path is not shadowed somewhere
// on its path in the map.
// e.g., if "foo.bar" has a value in the given map, it “shadows”
//       "foo.bar.baz" in a lower-priority map
func (v *kv) isPathShadowedInDeepMap(path []string, m map[string]interface{}) string {
	var parentVal interface{}
	for i := 1; i < len(path); i++ {
		parentVal = v.searchMap(m, path[0:i])
		if parentVal == nil {
			// not found, no need to add more path elements
			return ""
		}
		switch parentVal.(type) {
		case map[interface{}]interface{}:
			continue
		case map[string]interface{}:
			continue
		default:
			// parentVal is a regular value which shadows "path"
			return strings.Join(path[0:i], v.keyDelim)
		}
	}
	return ""
}

// searchMap recursively searches for a value for path in source map.
// Returns nil if not found.
// Note: This assumes that the path entries and map keys are lower cased.
func (v *kv) searchMap(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}

	next, ok := source[path[0]]
	if ok {
		// Fast path
		if len(path) == 1 {
			return next
		}

		// Nested case
		switch nextValue := next.(type) {
		case map[interface{}]interface{}:
			return v.searchMap(cast.ToStringMap(next), path[1:])
		case map[string]interface{}:
			// Type assertion is safe here since it is only reached
			// if the type of `next` is the same as the type being asserted
			return v.searchMap(nextValue, path[1:])
		default:
			// got a value but nested key expected, return "nil" for not found
			return nil
		}
	}
	return nil
}
