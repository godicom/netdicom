package netdicom

// import (
// 	"fmt"
// )

const (
	Ev_A_ASSOCIATE_Request             = 0
	Ev_Transport_Connection_Confirm    = 1
	Ev_A_ASSOCIATE_AC_PDU              = 2
	Ev_A_ASSOCIATE_RJ_PDU              = 3
	Ev_Transport_Connection_Indication = 4
	Ev_A_ASSOCIATE_RQ_PDU              = 5
	Ev_A_ASSOCIATE_Response_Accept     = 6
	Ev_A_ASSOCIATE_Response_Reject     = 7
	Ev_P_DATA_Request                  = 8
	Ev_P_DATA_TF_PDU                   = 9
	Ev_A_RELEASE_Request               = 10
	Ev_A_RELEASE_RQ_PDU                = 11
	Ev_A_RELEASE_RP_PDU                = 12
	Ev_A_RELEASE_Response              = 13
	Ev_A_ABORT_Request                 = 14
	Ev_A_ABORT_PDU                     = 15
	Ev_Transport_Connection_Closed     = 16
	Ev_ATRIM_Timer_Expired             = 17
	Ev_Unrecognized_Or_Invalid_Pdu     = 18
)

const (
	state1Idle                               = 0
	state2TransportConnectionOpen            = 1
	state3AwaitingLocalAAssocRsp             = 2
	state4AwaitingTransportConnection        = 3
	state5AwaitingAAssocACOrAAssocRJ         = 4
	state6ReadyFoDataTransfer                = 5
	state7AwaitingAReleaseRsp                = 6
	state8AwaitingLocalAREleaseRsp           = 7
	state9CollisionAwaitingLocalAReleaseRsp  = 8
	state10CollisionAwaitingAReleaseRsp      = 9
	state11CollisionAwaitingAReleaseRsp      = 10
	state12CollisionAwaitingLocalAReleaseRsp = 11
)

const (
	AE1 = 0
	AE2 = 1
	AE3 = 2
	AE4 = 3
	AE5 = 4
	AE6 = 5
	AE7 = 6
	AE8 = 7

	DT1 = 8
	DT2 = 9

	AR1  = 10
	AR2  = 11
	AR3  = 12
	AR4  = 13
	AR5  = 14
	AR6  = 15
	AR7  = 16
	AR8  = 17
	AR9  = 18
	AR10 = 19

	AA1 = 20
	AA2 = 21
	AA3 = 22
	AA4 = 23
	AA5 = 24
	AA6 = 25
	AA7 = 26
	AA8 = 27
)

type assoc struct {
	currentState int
	// currentEvent int
}

// type stateFn func(*assoc) stateFn

// // Need better names for actions
// func AE1Action(l *assoc) error {

// 	return nil, nil
// }

// func AE2Action(l *assoc) error {
// 	return nil, nil
// }

// func callAction(int action, l *assoc) error {
// 	// switch(l.currentAction)
// 	//What to do when it fail???
// }

func doNextState(l *assoc) {
	switch l.currentState {
	case state1Idle:
	case state2TransportConnectionOpen:
	case state3AwaitingLocalAAssocRsp:
	case state4AwaitingTransportConnection:
	case state5AwaitingAAssocACOrAAssocRJ:
	case state6ReadyFoDataTransfer:
	case state7AwaitingAReleaseRsp:
	case state8AwaitingLocalAREleaseRsp:
	case state9CollisionAwaitingLocalAReleaseRsp:
	case state10CollisionAwaitingAReleaseRsp:
	case state11CollisionAwaitingAReleaseRsp:
	case state12CollisionAwaitingLocalAReleaseRsp:
	}

}

// func state1Idle(l *assoc) stateFn {
// 	fmt.Println("st1 Idle")

// 	if l.currentEvent == Ev_A_ASSOCIATE_Request {
// 		callAction(AE1, l)
// 		return state4AwaitingTransportConnection(l)
// 	}

// 	return nil
// }

// //Need better names:
// func state2TransportConnectionOpen(l *assoc) stateFn {
// 	fmt.Println("st2 TransportConnectionOpen")
// 	if l.currentEvent == Ev_Transport_Connection_Confirm {
// 		callAction(AE2, l)
// 		return state5AwaitingAAssocACOrAAssocRJ(l)
// 	}

// 	return nil
// }

// func state3AwaitingLocalAAssocRsp(l *assoc) stateFn {
// 	return nil
// }

// func state4AwaitingTransportConnection(l *assoc) stateFn {
// 	fmt.Println("Waiting transport connection")
// 	return nil
// }

// func state5AwaitingAAssocACOrAAssocRJ(l *assoc) stateFn {
// 	return nil
// }

// func state6ReadyFoDataTransfer(l *assoc) stateFn {
// 	return nil
// }

// func state7AwaitingAReleaseRsp(l *assoc) stateFn {
// 	return nil
// }

// func state8AwaitingLocalAREleaseRsp(l *assoc) stateFn {
// 	return nil
// }

// func state9CollisionAwaitingLocalAReleaseRsp(l *assoc) stateFn {
// 	return nil
// }

// func state10CollisionAwaitingAReleaseRsp(l *assoc) stateFn {
// 	return nil
// }

// func state11CollisionAwaitingAReleaseRsp(l *assoc) stateFn {
// 	return nil
// }

// func state12CollisionAwaitingLocalAReleaseRsp(l *assoc) stateFn {
// 	return nil
// }

// func dicomStateMachineExec() {
// 	a := assoc{state1Idle}

// 	for state := state1Idle(&a); state != nil; {
// 		state = state(&a)
// 	}
// }
