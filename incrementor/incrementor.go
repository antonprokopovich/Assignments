package inctrementor

import (
	"fmt"
	"math"
)

const (
	// Дефолтное максимальное допустимое значение числа инкрементора.
	// Равно максимальному значению 64-битного целого числа (тип int64).
	maxNumValue = math.MaxInt64
)

// Интерфейс, который должен быть реализован конкретным инкрементером.
type Incrementor interface {
	// GetNumber возвращает текущее число.
	GetNumber() int

	// IncrementNumber увеличивает текущее число.
	IncrementNumber()

	// SetMaximumValue устанавливает максимальное значение текущего числа.
	// Хранимое число не может превышать установленное максимальное значение.
	SetMaximumValue(maxValue int)
}

// MyIncrementor реализует Incrementor.
type MyIncrementor struct {
	// currentNum хранит текущее число инкрементора
	currentNum int
	// maxValue хранит максимальное допустимое значение текущего числа
	maxValue int
}

// GetNumber реализует Incrementor.GetNumber.
func (inc *MyIncrementor) GetNumber() int {
	return inc.currentNum
}

// GetNumber реализует Incrementor.IncrementNumber,
// увеличивая число на один.
// Если при очередном увеличении чило достигает максимального
// допустимого значения, то оно сбрасывается в ноль.
func (inc *MyIncrementor) IncrementNumber() {
	if inc.currentNum < inc.maxValue - 1 {
		inc.currentNum += 1
	} else if inc.currentNum == inc.maxValue - 1 {
		inc.currentNum = 0
	}
}

// SetMaximumValue реализует Incrementor.SetMaximumValue.
// Не позволяет устанавливать число меньше нуля в качестве
// максимального, выводя предупреждающее сообщение.
func (inc *MyIncrementor) SetMaximumValue(maxValue int) {
	if maxValue >= 0 && maxValue > inc.currentNum {
		inc.maxValue = maxValue
	} else if maxValue >= 0 && maxValue <= inc.currentNum {
		inc.maxValue = maxValue
		inc.currentNum = 0
	} else {
		fmt.Printf(
			"Cant't set maximum value to %d. " +
				"The value can not be less then 0", maxValue,
		)
	}
}

// NewMyIncrementor создает новый экземпляр MyIncrementor.
// Задает начальные значения текущего и максимального числа.
func NewMyIncrementor() MyIncrementor {
	inc := MyIncrementor{
		currentNum: 0,
		maxValue: maxNumValue,
	}
	return inc
}
