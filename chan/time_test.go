package _chan

import (
	"testing"
	"time"
)

func TestTimeAfter(t *testing.T) {
	t.Log("test")
	<-time.After(3 * time.Second)
	t.Log("test2")
}
