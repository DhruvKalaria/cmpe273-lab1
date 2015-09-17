package fibonacci
import "testing"

func Test(t *testing.T)	{
	
	num := Fibo(5)
	
	if num!=3	{
		t.Error("Expected 3, but got",num)
	}
}