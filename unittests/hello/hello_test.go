package hello
import ("testing")
func TestSayHello(t* testing.T){
	expected := "Hello:TestUser\n"
	result := SayHello("TestUser")
	if result != expected{
		t.Errorf("Expected %v but got %v", expected, result)
	}
}