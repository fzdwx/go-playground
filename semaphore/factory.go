package semaphore

// Create new Semaphore
func Create() Semaphore {
	return make(Semaphore)
}

// Of Create a Semaphore with amount buffers
func Of(amount int) Semaphore {
	return make(Semaphore, amount)
}
