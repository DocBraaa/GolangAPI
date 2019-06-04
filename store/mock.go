package store

import (
	"errors"
	"strconv"
	"time"

	"github.com/music/model"
)

// MockStore Défini un type MockStore
type MockStore map[string]model.Track

// All fonction de la "classe" MockStore, qui respectent l'interface TrackStore
func (s MockStore) All() ([]model.Track, error) {
	var data []model.Track

	for _, track := range s {
		data = append(data, track)
	}

	return data, nil
}

// Create save a new track
func (s MockStore) Create(t *model.Track) error {
	timestamp := time.Now().Unix()
	id := strconv.FormatInt(timestamp, 10)

	t.ID = id
	s[id] = *t
	return nil
}

// Find search a track by its id and return it
func (s MockStore) Find(ID string) (*model.Track, error) {
	if track, ok := s[ID]; ok {
		return &track, nil
	}

	return nil, errors.New("track not found")
}

// Update a track
func (s MockStore) Update(t *model.Track) error {
	return nil
}

// Delete a track
func (s MockStore) Delete(ID string) error {
	delete(s, ID)
	return nil
}
