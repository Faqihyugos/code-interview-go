package vokal

import (
	"testing"
)

func TestCountVokal(t *testing.T) {
	result := CountVokal("Faqih")
	if result != "kalimat Faqih mengandung vokal sebanyak: 2" {
		panic("Result is not vokal")
	}
}

// func TestVokalSwitch(t *testing.T) {
// 	result := VokalSwitch('a')
// 	if result != "'a' adalah vokal" {
// 		panic("Result is not vokal")
// 	}
// }

// func TestIsVokal(t *testing.T) {
// 	result := IsVokal('b')
// 	if result != "'a' adalah vokal" {
// 		panic("Result is not vokal")
// 	}
// }


