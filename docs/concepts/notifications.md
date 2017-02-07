Notifications
==============
In Gu there exists the central notification backbone package `notifications`, which
exposes a system which allows registering specific functions for specific types of
structures to be called when such structures are dispatched into the system to allow
a decoupled form of communication.

*These allows components and pieces of Gu to communicate easily and also reduce tight
coupling of pieces.*

Using the `notifications` package is quite simple and need no elongated explanations,
by simply registering a function expecting a type, this sets this function to be called
once such a type is seen.

```go

import "github.com/gu-io/gu/notifications"

type event struct{
  EventName string
  EventType string
}

func main(){

  notifications.Subscribe(func(eventName string){
    fmt.Printf("EventName[%q] occured.\n", eventName)
  })


  notifications.Subscribe(func(e event){
    fmt.Printf("EventName[%q] and EventType[%q] occured.\n", e.EventName, e.EventType)
  })

  notifications.Dispatch("Click") => `EventName["Click"] occured.`
  notifications.Dispatch(event{EventName:"ScrollDown", EventType:"scroll"}) => `EventName["ScrollDown"] and EventType["scroll"] occured.`
}
```
