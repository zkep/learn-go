package design

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 简单工厂
func TestNewIRuleConfigParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 工厂方法
func TestNewIRuleConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParserFactory
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParserFactory{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParserFactory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParserFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 抽象工厂

func Test_jsonConfigParserFactory_CreateRuleParser(t *testing.T) {
	tests := []struct {
		name string
		want IRuleConfigParser
	}{
		{
			name: "json",
			want: jsonRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonConfigParserFactory{}
			if got := j.CreateRuleParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRuleParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonConfigParserFactory_CreateSystemParser(t *testing.T) {
	tests := []struct {
		name string
		want ISystemConfigParser
	}{
		{
			name: "json",
			want: jsonSystemConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonConfigParserFactory{}
			if got := j.CreateSystemParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSystemParser() = %v, want %v", got, tt.want)
			}
		})
	}
}




// A 依赖关系 A -> B -> C
type A struct {
	B *B
}

// NewA NewA
func NewA(b *B) *A {
	return &A{
		B: b,
	}
}

// B B
type B struct {
	C *C
}

// NewB NewB
func NewB(c *C) *B {
	return &B{C: c}
}

// C C
type C struct {
	Num int
}

// NewC NewC
func NewC() *C {
	return &C{
		Num: 1,
	}
}

func TestContainer(t *testing.T) {
	container := New()
	err := container.Provide(NewA)
	assert.Nil(t,err)

	err = container.Provide(NewB)
	assert.Nil(t,err)

	err = container.Provide(NewC)
	assert.Nil(t,err)

	err = container.Invoke(func(a *A) {
		fmt.Printf("%+v: %d", a, a.B.C.Num)
	})
	assert.Nil(t,err)
}