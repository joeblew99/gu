# Limitations
 By design `gu` was made to be simple and least complex to allow it to only provide rendering capabilities which means that certain things
 are not inbuilt. I believe declaring this is very important for anyone who wishes to use it. The limitations are short and are as follows:

  1. That this library is not a react like library where structures are able to bind and listen for change, 
  the complexity of that approach is undesired.

  2 The library takes a proactive play where the developer uses either a function call or the notification system to update 
  changes to screen and not a auto-update approach.
  
  3. It has no synchronization primitive for both server and client but rather takes a simple approach of where the client 
  takes control of the page rendering by re-rendering at target spots. Hence developers should not expect a 
  magical hydration-dehydration mechanism.
  
Beyond these 3, gu does what it promises and lets you achieve that with minimal learning.
