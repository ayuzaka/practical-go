package chapter13

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		a        int
		b        int
		operator string
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "足し算",
			args: args{
				a:        10,
				b:        2,
				operator: "+",
			},
			want:    12,
			wantErr: false,
		},
		{
			name: "不正な演算子を指定",
			args: args{
				a:        10,
				b:        2,
				operator: "?",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Calc(tt.args.a, tt.args.b, tt.args.operator)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
