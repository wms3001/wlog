package wlog

import "testing"

var wlog = Wlog{}

func TestWlog(t *testing.T) {
	wlog.Show("This is Test!!")
}
