package goblin

import ("testing"
        "testing/quick"
	"fmt"
	"math"
	"strconv"
)

// TODO: install github.com/stretchr/testify
// and make these unit tests use something other than just ==


func TestRoundTripFloat(t *testing.T) {
	f := func(flt float64) bool {
		flt = math.Abs(flt)
		needed := fmt.Sprintf("%f", flt)
		gotten := TestExpr(needed)
		result, _ := strconv.ParseFloat(gotten["value"].(string), 64)
		return flt == result
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestIota(t *testing.T) {
	gotten := TestExpr("iota")
	val := gotten["value"].(map[string]interface{})
	if val["type"] != "IOTA" {
		t.Error("Didn't parse iota as a literal")
	}

}

func TestTrue(t *testing.T) {
	gotten := TestExpr("true")
	val := gotten["value"].(map[string]interface{})
	if val["type"] != "BOOL" || val["value"] != "true" {
		t.Error("Didn't parse 'true' as true")
	}
}

func TestFalse(t *testing.T) {
	gotten := TestExpr("false")
	val := gotten["value"].(map[string]interface{})
	if val["type"] != "BOOL" || val["value"] != "false" {
		t.Error("Didn't parse 'false' as false")
	}
}

func TestProvidedImag(t *testing.T) {
	gotten := TestExpr("1.414i")
	if gotten["type"] != "IMAG" || gotten["value"] != "1.414i" {
		t.Error("Imaginary numbers not parsing correctly")
	}
}

func TestProvidedFloat(t *testing.T) {
	gotten := TestExpr("3.14")
	if gotten["value"].(string) != "3.14" {
		t.Error("Floats not parsing correctly")
	}
}

func TestCall(t *testing.T) {
	gotten := TestExpr("foo(bar)")
	if gotten["type"].(string) != "call" {
		t.Error("Function calls not parsing correctly")
	}
}


func TestRoundTripUInt(t *testing.T) {
	f := func(int uint64) bool {
		needed := fmt.Sprintf("%d", int)
		gotten := TestExpr(needed)
		return needed == gotten["value"]
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}