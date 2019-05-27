package hellootus

import "testing"

func TestHello(t *testing.T) {
	if greeting := Hello(); greeting != "Hello there!" {
		t.Error("doesnt work")
	}
}