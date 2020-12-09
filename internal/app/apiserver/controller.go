package apiserver

import "github.com/marktsoy/flashcards_api/internal/app/store"

type Controller struct {
	store store.Store

	// TODO add validator Instances
}
