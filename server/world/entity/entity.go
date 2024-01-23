package entity

type Entity interface {
	Interactor
}

type Interactor interface {

	// Interact target is the entity we left-clicked or are interacting with
	Interact(target Entity)
}
