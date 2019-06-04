package controller

import (
	"net/http"

	"github.com/music/model"
	"github.com/music/store"
	"github.com/music/utils"
)

// Track
type Track struct {
	Routes model.Routes
	Prefix string
	Store  store.TrackStore
}

// NewTrack oui
func NewTrack(s store.TrackStore) *Track {
	t := &Track{
		Routes: model.Routes{},
		Prefix: "/tracks",
		Store:  s,
	}

	routes := model.Routes{
		model.Route{
			Name:    "Create",
			Method:  "POST",
			Path:    t.Prefix,
			Handler: t.Create,
		},
	}

	t.Routes = append(t.Routes, routes...)

	return t
}

func (t Track) Create(w http.ResponseWriter, r *http.Request) {
	track := &model.Track{}

	utils.ParseJSON(r, track)
}
