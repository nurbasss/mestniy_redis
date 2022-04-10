package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var (
	ErrNotFound = errors.New("not Found")
)

type Store struct {
	Database   map[string]string
	lock sync.Mutex
}

func NewStore() *Store {
	db := new(Store)
	return db
}

func (s *Store) Set(key string, value string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.Database == nil {
		s.Database = make(map[string]string)
	}
	s.Database[key] = value
}

func (s *Store) Get(key string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	v, ok := s.Database[key]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

func (s *Store) Save() error {
	_, err := json.Marshal(s.Database)
	if err != nil {
		return err
	}
	file, _ := json.MarshalIndent(s.Database, "", " ")
	_ = ioutil.WriteFile("mestniy_database.json", file, 0644)
	return nil
}

func (s *Store) Recover() error {
	c := make(map[string]string)

	absPath, _ := filepath.Abs("mestniy_database.json")
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return err
	}

	e := json.Unmarshal(data, &c)
	if e != nil {
		panic(e)
	}

	i := 0
	for k, v := range c {
		s.Set(k, v)
		i++
	}
	return nil
}