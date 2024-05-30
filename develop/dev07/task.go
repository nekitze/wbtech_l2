package main

func or(channels ...<-chan interface{}) <-chan interface{} {
	orChan := make(chan interface{})
	for _, channel := range channels {
		go func() {
			<-channel
			close(orChan)
		}()
	}
	return orChan
}
