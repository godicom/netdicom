package netdicom

import (
	"errors"
)

type AE struct {
	aet     string
	addr    string
	timeout uint
}

func (a *AE) AET() string {
	return a.aet
}

func (a *AE) Address() string {
	return a.addr
}

func (a *AE) Timeout() uint {
	return a.timeout
}

func Address(addr string) func(*AE) error {
	return func(a *AE) error {
		a.addr = addr
		// TODO: setup listener here
		return nil
	}
}

func AET(aet string) func(*AE) error {
	return func(a *AE) error {
		if len(aet) > 16 {
			return errors.New("AE Title should be 16 or less characters")
		}
		a.aet = aet
		return nil
	}
}

func Timeout(timeout uint) func(*AE) error {
	return func(a *AE) error {
		a.timeout = timeout
		return nil
	}
}

func NewAE(options ...func(*AE) error) (*AE, error) {
	ae := &AE{"aet", "", 0}
	for _, option := range options {
		err := option(ae)
		if err != nil {
			return nil, err
		}
	}

	return ae, nil
}
