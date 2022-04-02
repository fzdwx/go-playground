package semaphore

// Acquire amount resource
func (this Semaphore) Acquire(amount int) {
	e := new(Element)
	for i := 0; i < amount; i++ {
		this <- e
	}
}

// Release amount resource
func (this Semaphore) Release(amount int) {
	for i := 0; i < amount; i++ {
		<-this
	}
}

//
// mutex
//

// Lock current block
func (this Semaphore) Lock() {
	this.Acquire(1)
}

// Unlock current block
func (this Semaphore) Unlock() {
	this.Release(1)
}

//
// signal
//

func (this Semaphore) Wait(amount int) {
	this.Acquire(amount)
}

func (this Semaphore) Signal() {
	this.Release(1)
}
