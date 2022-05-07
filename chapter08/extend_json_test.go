package chapter08

import "testing"

func TestDelayDecode(t *testing.T) {
	type args struct {
		fileName string
	}

	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "response1",
			args: args{
				fileName: "response1.json",
			},
			want: Message{
				ID:      "123456",
				UserID:  "ABC",
				Message: "Hello World!",
			},
		},
		{
			name: "response2",
			args: args{
				fileName: "response2.json",
			},
			want: Sensor{
				ID:       "09876",
				DeviceID: "YJKR67",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DelayDecode(tt.args.fileName)
			if got != tt.want {
				t.Errorf("DelayDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
