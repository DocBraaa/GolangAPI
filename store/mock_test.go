package store

import (
	"testing"

	"github.com/music/model"
)

func TestCreate(t *testing.T) {

	track := model.Track{
		Title: "Titre of my track",
	}

	store := MockStore{}
	store.Create(&track)

	t.Log(track.ID)

	if track.ID == "" {
		t.Error("ID of track is empty")
	}
}
