package lfsr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_LFSR_NewFromCorrectString(t *testing.T) {
	_, err := NewLFSRFromBitString("101010100101", "100010101010101010", "010101010100")
	require.Nil(t, err)
}

func Test_LFSR_NewFromIncorrectString(t *testing.T) {
	_, err := NewLFSRFromBitString("121011140101", "100010101010101010", "010101010100")
	require.NotNil(t, err)
}

func Test_LFSR_Move(t *testing.T) {
	// from practice
	lfsr, err := NewLFSRFromBitString("100", "11", "100+11")
	require.Nil(t, err)

	expect := []uint8{0, 1, 1}

	for _, item := range expect {
		res := lfsr.Move()

		require.Equal(t, item, res)
	}
}
