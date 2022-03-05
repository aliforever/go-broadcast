package broadcast

import (
	"fmt"
	"sync"
	"testing"
)

type User struct {
	FirstName string
	LastName  string
}

func TestBroadcast(t *testing.T) {
	b := NewBroadcast[User]()

	ch, err := b.AddChannel("new_users")
	if err != nil {
		t.Fatal(err)
	}

	listener1 := ch.AddListener()
	listener2 := ch.AddListener()
	listener3 := ch.AddListener()

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		u := <-listener1
		fmt.Println("Received", u, "In Listener 1")
		wg.Done()
	}()

	go func() {
		u := <-listener2
		fmt.Println("Received", u, "In Listener 2")
		wg.Done()
	}()

	go func() {
		u := <-listener3
		fmt.Println("Received", u, "In Listener 3")
		wg.Done()
	}()

	ch.InformListeners(User{
		FirstName: "Ali",
		LastName:  "Error",
	})

	wg.Wait()
}
