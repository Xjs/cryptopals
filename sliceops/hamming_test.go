package sliceops

import "testing"

func TestHammingDistance(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"ex5",
			args{
				[]byte("this is a test"),
				[]byte("wokka wokka!!!"),
			},
			37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HammingDistance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("HammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
