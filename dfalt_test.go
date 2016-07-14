package dfalt_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stuart-warren/dfalt"
)

func TestEnvString(t *testing.T) {

	e1 := "my_string"
	a1 := dfalt.EnvString("NONEXIST", e1)
	if a1 != e1 {
		t.Errorf("Int: expected %d, actual %d", e1, a1)
	}

	e2 := "your_string"
	_ = os.Setenv("VAL_STR", e2)
	a2 := dfalt.EnvString("VAL_STR", "")
	if a2 != e2 {
		t.Errorf("Int: expected %d, actual %d", e2, a2)
	}
}

func TestEnvInt(t *testing.T) {

	e1 := 5
	a1 := dfalt.EnvInt("NONEXIST", e1)
	if a1 != e1 {
		t.Errorf("Int: expected %d, actual %d", e1, a1)
	}

	e2 := 9
	_ = os.Setenv("VAL_INT", strconv.Itoa(e2))
	a2 := dfalt.EnvInt("VAL_INT", 0)
	if a2 != e2 {
		t.Errorf("Int: expected %d, actual %d", e2, a2)
	}
}
