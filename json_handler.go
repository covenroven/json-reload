package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
)

const filepath = "storage/status.json"

type Wind uint64

// String returns string representation of the wind
func (w Wind) String() string {
	return fmt.Sprintf("%d meter per detik", w)
}

// StatusText returns status text representation of the wind
func (w Wind) StatusText() string {
	var status string
	switch {
	case w < 6:
		status = "aman"
	case w <= 15:
		status = "siaga"
	default:
		status = "bahaya"
	}

	return status
}

type Water uint64

// String returns string representation of the water
func (w Water) String() string {
	return fmt.Sprintf("%d meter", w)
}

// StatusText returns status text representation of the water
func (w Water) StatusText() string {
	var status string
	switch {
	case w < 5:
		status = "aman"
	case w <= 8:
		status = "siaga"
	default:
		status = "bahaya"
	}

	return status
}

type StatusJSON struct {
	Wind  Wind
	Water Water
}

// Randomize generates new value in secure number range 1-100
func (sj *StatusJSON) Randomize() {
	max := big.NewInt(99)

	wind, _ := rand.Int(rand.Reader, max)
	sj.Wind = Wind(wind.Uint64())
	sj.Wind++

	water, _ := rand.Int(rand.Reader, max)
	sj.Water = Water(water.Uint64())
	sj.Water++
}

// ReadJSON reads file from JSON and return it
func ReadJSON() (StatusJSON, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return StatusJSON{}, err
	}

	var status struct {
		Status StatusJSON
	}
	err = json.Unmarshal(content, &status)
	if err != nil {
		return StatusJSON{}, err
	}

	return status.Status, nil
}

// WriteJSON writes new values to the JSON file
func WriteJSON(status StatusJSON) error {
	result, err := json.Marshal(struct {
		Status StatusJSON
	}{
		Status: status,
	})
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, result, 0644); err != nil {
		return err
	}

	return nil
}
