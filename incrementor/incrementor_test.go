package inctrementor

import (
	"testing"
)

func TestIncrementNumberOnce(t *testing.T) {
	inc := NewMyIncrementor()
	inc.IncrementNumber()

	expected := 1
	if observed := inc.currentNum; observed != expected {
		t.Fatalf("currentNum = %v, expected %v", observed, expected)
	}
}

func TestIncrementNumberManyTimes(t *testing.T) {
	inc := NewMyIncrementor()
	for i := 0; i < 100; i++ {
		inc.IncrementNumber()
	}

	expected := 100
	if observed := inc.currentNum; observed != expected {
		t.Fatalf("currentNum = %v, expected %v", observed, expected)
	}
}

func TestInitialMaximumValueEqualsMaximumInt64(t *testing.T) {
	inc := NewMyIncrementor()

	maxInt64value := 9223372036854775807
	if observed := inc.maxValue; observed != maxInt64value {
		t.Fatalf("currentNum = %v, expected %v", observed, maxInt64value)
	}
}

func TestSetMaximumValue(t *testing.T) {
	inc := NewMyIncrementor()
	inc.SetMaximumValue(1000)

	expected := 1000
	if observed := inc.maxValue; observed != expected {
		t.Fatalf("currentNum = %v, expected %v", observed, expected)
	}
}

func TestCurrentNumberIsZeroUponReachingMaximumValue(t *testing.T) {
	inc := NewMyIncrementor()
	inc.SetMaximumValue(5)

	for i := 0; i < 5; i++ {
		inc.IncrementNumber()
	}

	if observed := inc.currentNum; observed != 0 {
		t.Fatalf("currentNum = %v, expected %v", observed, 0)
	}
}

func TestTrySetNegativeMaxValue(t *testing.T) {
	inc := NewMyIncrementor()
	initMax := inc.maxValue

	inc.SetMaximumValue(-1)

	if observed := inc.maxValue; observed != initMax {
		t.Fatalf("currentNum = %v, expected %v", observed, initMax)
	}
}

func BenchmarkIncrementNumber(b *testing.B) {
	inc := NewMyIncrementor()
	for i := 0; i < b.N; i++ {
		inc.IncrementNumber()
	}
}
