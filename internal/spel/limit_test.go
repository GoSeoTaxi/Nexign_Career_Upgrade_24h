package spel

import (
	"Nexign_Career_Upgrade_24h/internal/etc"
	"fmt"
	"reflect"
	"testing"
)

func TestLimitReq(t *testing.T) {
	type args struct {
		arT []string
	}

	//	abc := []string{"Очень большой и важный текст", "Ощибка", "Земля", "Марс", "Юпитер", "Сатурн", "Уран", "Нептун"}

	//	args1 := args{arT: []string{"Очень большой и важный текст", "Ощибка", "Земля", "Марс", "Юпитер", "Сатурн", "Уран", "Нептун"}}

	//	fmt.Println(args1)

	var a123 = etc.AnsJ{Word: "Ощибка", S: []string{"Ошибка"}}
	var want1 []etc.AnsJ
	want1 = append(want1, a123)

	tests := []struct {
		name string
		args args
		want []etc.AnsJ
	}{
		{
			name: "1",
			args: args{arT: []string{"важный текст", "Ощибка", "Земля", "Марс", "Юпитер", "Сатурн", "Уран", "Нептун"}},
			want: want1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LimitReq(tt.args.arT); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LimitReq() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
				fmt.Println(tt.want)
			}
		})
	}
}
