package netdicom

import (
	"testing"
)

func TestStateMachineSimple(t *testing.T) {
	a := newAssoc()
		
 	go a.proc()
 	a.putEvent(Ev_A_ASSOCIATE_Request)
 
}
