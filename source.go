package source

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Source map[string]string
type UnmarshalFunc func(data []byte, v interface{}) error

var source = Source{}
var delimiter = "."

func NewSource() *Source {
	return &Source{}
}

// SetDelimeter - set custom delimeter
// default delimeter "."
func SetDelimeter(d string) {
	delimiter = d
}

// FromFile - fill source values ​​from file on disk
// path - path to file on disk
// unmarshalFunc - func to unmarshal (JSON, TOML, YAML...)
// default: json.Unmarshal
func (s *Source) FromFile(path string, unmarshalFunc UnmarshalFunc) error {
	data, err := readFile(path)
	if err != nil {
		return err
	}
	return unmarshalData(s, data, unmarshalFunc)
}

// FromData - fill source values ​​from data
// data - []byte
// unmarshalFunc - func to unmarshal (default: json.Unmarshal)
func (s *Source) FromData(data []byte, unmarshalFunc UnmarshalFunc) error {
	return unmarshalData(s, data, unmarshalFunc)
}

// FromFile - fill source values ​​from file on disk
// path - path to file on disk
// unmarshalFunc - func to unmarshal (JSON, TOML, YAML...)
// default: json.Unmarshal
func FromFile(path string, unmarshalFunc UnmarshalFunc) error {
	data, err := readFile(path)
	if err != nil {
		return err
	}
	return unmarshalData(&source, data, unmarshalFunc)
}

// FromData - fill source values ​​from data
// data - []byte
// unmarshalFunc - func to unmarshal (default: json.Unmarshal)
func FromData(data []byte, unmarshalFunc UnmarshalFunc) error {
	return unmarshalData(&source, data, unmarshalFunc)
}

func unmarshalData(s *Source, d []byte, unmarshalFunc UnmarshalFunc) error {
	if unmarshalFunc == nil {
		unmarshalFunc = json.Unmarshal
	}

	var raw interface{}
	err := unmarshalFunc(d, &raw)
	if err != nil {
		return err
	}
	prepareData(s, raw, nil, "")
	return nil
}

func prepareData(src *Source, v interface{}, path []string, key string) {
	//is root
	if path == nil {
		path = []string{}
	}
	if key != "" {
		path = append(path, key)
	}

	switch v.(type) {
	case string:
		key := strings.Join(path, delimiter)
		put(src, key, v.(string))
	case map[string]interface{}:
		m := v.(map[string]interface{})
		for k, v := range m {
			prepareData(src, v, path, k)
		}
	}
}

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func put(s *Source, key, value string) {
	src := *s
	src[key] = value
}

func get(src *Source, key string) string {
	s := *src
	if res, ok := s[key]; ok {
		return res
	}
	return key
}

func getList(s *Source, keys ...string) []string {
	lst := []string{}
	for _, k := range keys {
		lst = append(lst, get(s, k))
	}
	return lst
}

// Get - returns dict value by key
// If not exists, returns transmitted key
func (s *Source) Get(key string) string {
	return get(s, key)
}

// Get - returns dict value by key
// If not exists, returns transmitted key
func Get(key string) string {
	return get(&source, key)
}

// GetList - returns dict values by keys
// If not exists, returns transmitted key
// for each no found dictionary key
func (s *Source) GetList(keys ...string) []string {
	return getList(s, keys...)
}

// GetList - returns dict values by keys
// If not exists, returns transmitted key
// for each no found dictionary key
func GetList(keys ...string) []string {
	return getList(&source, keys...)
}
