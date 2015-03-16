package netdicom

import (
	"fmt"
	"errors"
)

const (
	Ev_A_ASSOCIATE_Request             = iota
	Ev_Transport_Connection_Confirm    
	Ev_A_ASSOCIATE_AC_PDU              
	Ev_A_ASSOCIATE_RJ_PDU              
	Ev_Transport_Connection_Indication 
	Ev_A_ASSOCIATE_RQ_PDU              
	Ev_A_ASSOCIATE_Response_Accept     
	Ev_A_ASSOCIATE_Response_Reject     
	Ev_P_DATA_Request                  
	Ev_P_DATA_TF_PDU                   
	Ev_A_RELEASE_Request               
	Ev_A_RELEASE_RQ_PDU                
	Ev_A_RELEASE_RP_PDU                
	Ev_A_RELEASE_Response              
	Ev_A_ABORT_Request                 
	Ev_A_ABORT_PDU                     
	Ev_Transport_Connection_Closed     
	Ev_ATRIM_Timer_Expired             
	Ev_Unrecognized_Or_Invalid_Pdu     
)

type stateFn func(*assoc) stateFn

type assoc struct {
	nextState stateFn
	eventChanel chan int
	currentEvent int
}

func newAssoc() * assoc {
	var a assoc
	a.eventChanel = make(chan int)
	a.nextState = state1Idle(&a)
	return &a
}


func (a * assoc) putEvent(event int) error {
	a.eventChanel <- event

	return nil
}

// Need better names for actions
func AE1Action(l *assoc) error {
	fmt.Println("AE1Action")
	return nil
}

func AE2Action(l *assoc) error {
	fmt.Println("AE2Action")
	return nil
}

func AE3Action(l *assoc) error {
	fmt.Println("AE3Action")
	return nil
}
func AE4Action(l *assoc) error {
	return nil
} 
func AE5Action(l *assoc) error {
	return nil
} 
func AE6Action(l *assoc) error {
	return nil
} 
func AE7Action(l *assoc) error {
	return nil
} 
func AE8Action(l *assoc) error {
	return nil
} 

func DT1Action(l *assoc) error {
	return nil
} 
func DT2Action(l *assoc) error {
	return nil
} 

func AR1Action(l *assoc) error {
	return nil
} 
func AR2Action(l *assoc) error {
	return nil
} 
func AR3Action(l *assoc) error {
	return nil
} 
func AR4Action(l *assoc) error {
	return nil
} 
func AR5Action(l *assoc) error {
	return nil
} 
func AR6Action(l *assoc) error {
	return nil
} 
func AR7Action(l *assoc) error {
	return nil
} 
func AR8Action(l *assoc) error {
	return nil
} 
func AR9Action(l *assoc) error {
	return nil
} 
func AR10Action(l *assoc) error {
	return nil
}

func AA1Action(l *assoc) error {
	return nil
}
func AA2Action(l *assoc) error {
	return nil
}
func AA3Action(l *assoc) error {
	return nil
}
func AA4Action(l *assoc) error {
	return nil
}
func AA5Action(l *assoc) error {
	return nil
}
func AA6Action(l *assoc) error {
	return nil
}
func AA7Action(l *assoc) error {
	return nil
}
func AA8Action(l *assoc) error {
	return nil
}


func (a * assoc) proc() error {
	a.currentEvent = <- a.eventChanel
	if a.nextState != nil {
		a.nextState = a.nextState(a)
	} else {
		return errors.New("undefined state")
	}
	return nil
}

func state1Idle(l *assoc) stateFn {
	fmt.Println("state1 Idle")

	if l.currentEvent == Ev_A_ASSOCIATE_Request {
		if AE1Action(l) != nil {
			return nil
		}
		return state4AwaitingTransportConnection(l)
	}

	return nil
}

//Need better names:
func state2TransportConnectionOpen(l *assoc) stateFn {
	fmt.Println("state2 TransportConnectionOpen")
	if l.currentEvent == Ev_Transport_Connection_Confirm {
		if AE2Action(l) != nil {
			return nil
		}
		return state5AwaitingAAssocACOrAAssocRJ(l)
	}

	return nil
}

func state3AwaitingLocalAAssocRsp(l *assoc) stateFn {
	return nil
}

func state4AwaitingTransportConnection(l *assoc) stateFn {
	fmt.Println("Waiting transport connection")
	return nil
}

func state5AwaitingAAssocACOrAAssocRJ(l *assoc) stateFn {
	return nil
}

func state6ReadyFoDataTransfer(l *assoc) stateFn {
	return nil
}

func state7AwaitingAReleaseRsp(l *assoc) stateFn {
	return nil
}

func state8AwaitingLocalAREleaseRsp(l *assoc) stateFn {
	return nil
}

func state9CollisionAwaitingLocalAReleaseRsp(l *assoc) stateFn {
	return nil
}

func state10CollisionAwaitingAReleaseRsp(l *assoc) stateFn {
	return nil
}

func state11CollisionAwaitingAReleaseRsp(l *assoc) stateFn {
	return nil
}

func state12CollisionAwaitingLocalAReleaseRsp(l *assoc) stateFn {
	return nil
}
