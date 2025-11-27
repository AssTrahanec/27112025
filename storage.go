package main

import (
	"encoding/json"
	"os"
	"sync"
)

const dataFile = "data.json"

type Storage struct {
	mu       sync.RWMutex
	data     map[int][]LinkStatus
	counter  int
}

func NewStorage() *Storage {
	s := &Storage{
		data: make(map[int][]LinkStatus),
	}
	s.Load()
	return s
}

func (s *Storage) Add(links []LinkStatus) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.counter++
	s.data[s.counter] = links
	s.Save()
	return s.counter
}

func (s *Storage) Get(nums []int) []LinkStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []LinkStatus
	for _, num := range nums {
		if links, ok := s.data[num]; ok {
			result = append(result, links...)
		}
	}
	return result
}

func (s *Storage) Save() error {
	type Data struct {
		Counter int                     `json:"counter"`
		Links   map[int][]LinkStatus    `json:"links"`
	}

	data := Data{
		Counter: s.counter,
		Links:   s.data,
	}

	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func (s *Storage) Load() error {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	type Data struct {
		Counter int                     `json:"counter"`
		Links   map[int][]LinkStatus    `json:"links"`
	}

	var data Data
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return err
	}

	s.counter = data.Counter
	s.data = data.Links
	return nil
}
