common gotchas with map and struct
        """
        type Employee struct{
            age int
            name String
            number int
        }
        mymap[String]Employee
        mymap["sashank"].number++;// This is not possible because mymap["sashank"] returns a copy of struct
        go doesnt allow to modify the copy

        type User struct{
            age int `json:"foo"`
        }

        type User1 struct{
            age int `json:"foo"`
        }
        func main(){
            v1 := User{1}
            v2 := User1{2}
            v1 = User(v2)
            fmt.Println(v1) //2
        }
        Structs get copied when passed , modifying the struct doesn't modify the original struct unless passed with a reference
        Any struct with a mutex passed by reference
        type Emplyee Struct {
            mu sync.Mutex
            Name string
        }
        func do(emp *Employee){
            emp.mu.lock()
            defer mu.unlock()
        }