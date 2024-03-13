package lfsr

import (
	"errors"
	"strings"
)

var ErrLFSRCantBeEmpty = errors.New("lfsr can not be empty")

func NewLFSRFromBitString(state string, feedback string, output string) (*LFSR, error) {
	pState, err := bitStringToUint8Slice(state)
	if err != nil {
		return nil, err
	}

	pFeedback, err := bitStringToUint8Slice(feedback)
	if err != nil {
		return nil, err
	}

	pOutputs := [][]uint8{}

	for _, s := range strings.Split(output, "+") {
		res, err := bitStringToUint8Slice(s)
		if err != nil {
			return nil, err
		}

		pOutputs = append(pOutputs, res)
	}

	return NewLFSR(pState, pFeedback, pOutputs)
}

func NewLFSR(state []uint8, feedback []uint8, output [][]uint8) (*LFSR, error) {
	if len(state) == 0 && len(feedback) == 0 && len(output) == 0 {
		return nil, ErrLFSRCantBeEmpty
	}

	scaled := make([][]uint8, 0, len(output)+2)
	scaled = append(scaled, state, feedback)
	scaled = append(scaled, output...)
	scaled = scaleFromLeft(scaled...)

	state, feedback, output = scaled[0], scaled[1], scaled[2:]

	return &LFSR{
		state:    state,
		feedback: feedback,
		output:   output,
	}, nil
}

type LFSR struct {
	// We don't care about memory consumption
	// x_n, x_{n-1} ... x_0 >>
	state    []uint8
	feedback []uint8
	output   [][]uint8
}

func (r *LFSR) Move() uint8 {
	var newItem uint8
	var out uint8

	for i := len(r.state) - 1; i >= 0; i-- {
		newItem ^= r.state[i] & r.feedback[i]

		if i != 0 {
			r.state[i] = r.state[i-1]
		}
	}

	r.state[0] = newItem

	for i := 0; i < len(r.output); i++ {

		tmp := uint8(1)

		for j := len(r.state) - 1; j >= 0; j-- {
			if r.output[i][j] == 1 {
				tmp *= r.state[j]
			}
		}

		out ^= tmp
	}

	return out
}
