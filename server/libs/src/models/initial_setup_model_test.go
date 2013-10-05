package models

import (
	"fmt"
	"testing"
)

func Test_Nothing(t *testing.T) {
	fmt.Println("here")
	t.Error("Failing")
}
