# Notifications
The `Gu` library comes with a global notification system, found in [Gudispatch](./gudispatch). This allows one to listen/send for notification signals for views or any custom type you so wish. Infact views actually use this when listening to each other for changes or when set to listen for changes on browser history.


## Using the Notification System
Using the underline notification system is rather simple and needs no elongated explanations. The subpackage defines a subscription and publish function which allow functions to be called when the giving type of Notification is published. Using a bit of reflection, this system allows youu to Listen for any type.

- Listening And Publishing a ViewUpdate

```go

  gudispatch.Subscribe(func (update guviews.ViewUpdate){
    //....
  })

  gudispatch.Publish(guviews.ViewUpdate{
    ID: "some-view-id"
  })

```

- Listening And Publishing a Custom Type

```go

type InboxNotification struct{
  Read int
  UnRead int
}


  gudispatch.Subscribe(func (box InboxNotification){
    //....
  })

  gudispatch.Publish(InboxNotification{
    Read: 20,
    UnRead: 40,
  })

```
