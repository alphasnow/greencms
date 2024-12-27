// @author AlphaSnow

package utils

import "testing"

func TestPrivacyEmail(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "163Email", args: args{str: "mynotice@163.com"}, want: "myn****@163.com"},
		{name: "qqEmail", args: args{str: "15263548521@qq.com"}, want: "152****@qq.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrivacyEmail(tt.args.str); got != tt.want {
				t.Errorf("PrivacyEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrivacyPhone(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "T15200126329", args: args{str: "15200126329"}, want: "152****6329"},
		{name: "T17922550022", args: args{str: "17922550022"}, want: "179****0022"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrivacyPhone(tt.args.str); got != tt.want {
				t.Errorf("PrivacyPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrivacyUsername(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "6str", args: args{str: "wuli15"}, want: "wu****15"},
		{name: "10str", args: args{str: "xiakekeQOa"}, want: "xi****Oa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrivacyUsername(tt.args.str); got != tt.want {
				t.Errorf("PrivacyUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
