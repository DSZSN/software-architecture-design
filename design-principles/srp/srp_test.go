package srp

import (
	"testing"
)

func TestDiplomat(t *testing.T) {
	dip := NewDiplomat()

	// Speak chinese
	dip.SpeakChinese()

	// Speak english
	dip.SpeakEnglish()

	// Person running
	dip.Run()

	// Person jumping
	dip.Jump()
}
