# Go-Broadcast
You can use this package to broadcast any message (by the power of generics) to multiple channels like a piece of cake.

## Requirements
- go v1.18

## Download
`go get -u github.com/aliforever/go-broadcast`

## Example
Custom Type:
```go
type User struct {
	FirstName string
	LastName  string
}
```
Initialize:
```go
b := NewBroadcast[User]()
```
Add New Channel:
```go
ch, err := b.AddChannel("new_users")
if err != nil {
    fmt.Println(err)
    return
}
```
Add Listeners:
```go
listener1 := ch.AddListener()
listener2 := ch.AddListener()
listener3 := ch.AddListener()
```
Listen:
```go
go func() {
    fmt.Println("Received", <-listener1, "In Listener 1")
}()

go func() {
    fmt.Println("Received", <-listener2, "In Listener 2")
}()

go func() {
    fmt.Println("Received", <-listener3, "In Listener 3")
}()
```
Inform Listeners:
```go
ch.InformListeners(User{
    FirstName: "Ali",
    LastName:  "Error",
})
```