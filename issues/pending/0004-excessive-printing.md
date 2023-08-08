# Excessive Printing

The code imported from APS for the callback server contains excessive print
statements -- and even a `log.Fatalf()` call. This is -- not good.

- [ ] Think about how to best clean this up!
- [ ] Does the Callback code even belong here?
  - It is definitely a part of **Auth0**, which probably should be supported once in this package rather than multiple times in every other package. But...
  - [ ] Is this code "generic" enough? Does some of it need to go back to APS?
    - Since APS is the first package I've written that supports Auth0, I don't know if the way I'm handling it is too specific to my one use-case!
    - [ ] Solution: write more code! 😊
