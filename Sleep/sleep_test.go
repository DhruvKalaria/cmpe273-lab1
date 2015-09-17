package Sleep
import (
	"time"
	"testing"
)
func Test(t *testing.T)	{

	t1 := time.Now().Second()
	sleep(10)
	t2 := time.Now().Second()
	
	if (t2-t1)!=10	{
		t.Error("Expected 10, but got",t2-t1)
	}
}