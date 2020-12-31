package lesson1

import (
	"reflect"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{name: "positive", input: "RU", want: "Привет, мир!", wantErr: nil},
		{name: "negative lower case", input: "ru", want: "", wantErr: &HelloWorldError{}},
		{name: "negative wrong country", input: "ZZ", want: "", wantErr: &HelloWorldError{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := HelloWorld(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}

			if !reflect.DeepEqual(tc.wantErr, err) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}

}
