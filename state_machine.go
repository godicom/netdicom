package netdicom

import (
	"fmt"
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
	currentAction int
}

type stateFn func(*assoc) stateFn

type actionFn func(*assoc) (actionFn, error)

func AE1Action(l *assoc) (actionFn, error) {

	return nil, nil
}

func AE2Action(l *assoc) (actionFn, error) {

	return nil, nil
}

func state1Idle(l *assoc) stateFn {
	fmt.Println("Idle")

	if l.currentAction == AE1 {
		return state4AwaitingTransportConnection(l)
	}

	return nil
}

//Need better names:
func state2TransportConnectionOpen(l *assoc) stateFn {
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

func dicomStateMachineExec() {
	a := assoc{AE1}

	for state := state1Idle(&a); state != nil; {
		state = state(&a)
	}
}
