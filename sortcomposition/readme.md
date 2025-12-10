 Go's sort package can sort any custom data type if you implement this interface:
type Interface interface {
    Len() int
   Less(i, j int) bool
   Swap(i, j int)
  }
 If your type supports these three methods → Go can sort it.

⭐ The official Go concept name
Implicit Interface Implementation

 Go interfaces are satisfied implicitly, not explicitly.

This is one of the most powerful and unique design philosophies of Go.

🧠 Why Go does this?

It avoids tight coupling.
Your type doesn’t even need to know the interface exists.
As long as it has the right method set → BOOM, it works.
| Concept                    | Meaning                                                                             |
| -------------------------- | ----------------------------------------------------------------------------------- |
| **Implicit interfaces**    | You don't “implement” an interface explicitly. Go checks the methods automatically. |
| **Duck typing**            | If it behaves like the interface, Go treats it as that interface.                   |
| **Type satisfy interface** | Your struct “satisfies” the interface because it has the required methods.          |


A method in Go can be called even when the receiver is nil,
👉 as long as the method itself handles nil safely.

In other words:

Go doesn’t automatically stop you from calling a method on a nil pointer.

This is different from languages where calling a method on null crashes immediately.

In Go, the method gets called, and inside it, the receiver value (r) might be nil.

Up to YOU to handle it.


Example:

    package main

import "fmt"

type User struct {
    Name string
}

func (u *User) PrintName() {
    if u == nil {
        fmt.Println("User is nil!")
        return
    }
    fmt.Println("Name:", u.Name)
}

func main() {
    var u *User = nil
    u.PrintName()  // METHOD STILL GETS CALLED!
}
Go does not automatically block the call because the receiver is nil

The compiler only ensures the receiver's type matches, not its value.