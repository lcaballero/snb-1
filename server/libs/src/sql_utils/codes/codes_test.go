package codes

import (
	"testing"
)

func Test_Checking_Equality(t *testing.T) {

	if Success != Success {
		t.Error("Comparing same codes should produce boolean 'true'.")
	}
}

func Test_Comparing_Different_Codes(t *testing.T) {

	if Success == Db_Error {
		t.Error("Comparing different codes should produce boolean 'false'.")
	}
}

// This function basically makes sure that someone cannot make a new
// code outside of this package and compare those codes with the
// codes found in this package.  The idea being that these structs
// are specified and remain unique.
func Test_Comparing_Codes_To_New_Instances(t *testing.T) {

	c := StatusCode{200, "Success", "", ""}

	if Success == c {
		t.Error("New codes should not be equal.")
	}
}
