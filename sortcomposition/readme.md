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
