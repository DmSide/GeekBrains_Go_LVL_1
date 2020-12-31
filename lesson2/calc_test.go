package lesson2

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		name    string
		input1  string
		input2  string
		input3  string
		want    float64
		wantErr error
	}{
		{name: "positive_+", input1: "10", input2: "3", input3: "+", want: 13, wantErr: nil},
		{name: "positive_-", input1: "10", input2: "3", input3: "-", want: 7, wantErr: nil},
		{name: "positive_*", input1: "10", input2: "3", input3: "*", want: 13, wantErr: nil},
		{name: "positive_/", input1: "10", input2: "2", input3: "/", want: 5, wantErr: nil},
		{name: "negative_/_0", input1: "10", input2: "0", input3: "/", want: 0, wantErr: fmt.Errorf("ой, Деление на ноль невозможно")},
		{name: "positive_^", input1: "10", input2: "3", input3: "^", want: 1000, wantErr: nil},
		{name: "negative_wrong_input1", input1: "string", input2: "0", input3: "+", want: 0, wantErr: fmt.Errorf("ой, Ошибка чтения первого числа")},
		{name: "negative_wrong_input2", input1: "10", input2: "string", input3: "+", want: 0, wantErr: fmt.Errorf("ой, Ошибка чтения второго числа")},
		{name: "negative_wrong_input3", input1: "10", input2: "0", input3: "string", want: 0, wantErr: fmt.Errorf("ой, Ошибка чтения арифмметической операции")},
		{name: "negative_wrong_op", input1: "10", input2: "0", input3: "op", want: 0, wantErr: fmt.Errorf("ой, Операция выбрана неверно")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var stdin bytes.Buffer

			stdin.Write([]byte(tc.input1 + "\n" + tc.input2 + "\n" + tc.input3 + "\n"))

			got, err := Calculator(&stdin)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}

			if !reflect.DeepEqual(tc.wantErr, err) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}

}
