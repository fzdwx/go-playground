package semaphore

type (
	Element interface{}
	//
	// Semaphore is a simple implementation of a semaphore.
	//
	// based Go Channel
	Semaphore chan Element
)
