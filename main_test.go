package main

import (
	"encoding/json"
	"fmt"
	"testing"
	//"strconv"
	//"bytes"
)

func pa(args []string, err error) {
	fmt.Println("[")
	for i, e := range args {
		fmt.Println(i, e)
	}
	fmt.Println("]")
}

func BenchmarkQuoteparse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quoteparse(`My name is 'Johhnnny' "qhat\'s yours" aaaaaaaa "Ur mum'\'\'" bbb hahaha "us them whatcha gonna do boy" "ending str" '"testing my quoteparse2"'`)
	}
}

func BenchmarkQuoteparse2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quoteparse2(`My name is 'Johhnnny' "qhat\'s yours" aaaaaaaa "Ur mum'\'\'" bbb hahaha "us them whatcha gonna do boy" \"testing\ my\ quoteparse2"`)
	}
}

func BenchmarkSpaceparse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		spaceparse(`My name is Johhnnny qhat's\ yours aaaaaaaa Ur\ mum'\\'\\' bbb hahaha us\ them\ whatcha\ gonna\ do\ boy ending\ str`)
	}
}

var Y []string

func BenchmarkJSONparse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f()
	}
}

func f() {
	//json.Unmarshal([]byte(`{"Data": ["datass", "hello", "whats \"up sugar pussy"]}`), &X)
	//X.Data = []string{}
	json.Unmarshal([]byte(`["My","name","is",     "Johhnnny","qhat\\'s yours","aaaaaaaa","Ur mum'\\'\\'","bbb","hahaha","us them whatcha gonna do boy"]`), &Y)
	Y = []string{}
}

func BenchmarkHGET(b *testing.B) {
	handle_message("HSET mykey myfield myvalue333")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handle_message("HGET mykey myfield")
	}
}

// func TestParse (t *testing.T) {
// 	f()
// 	pa(Y, nil)
// 	pa(quoteparse(`My name is 'Johhnnny' "qhat\'s yours" aaaaaaaa "Ur mum'\'\'" bbb hahaha "us them whatcha gonna do boy" "ending str" '"testing my quoteparse2"'`))
// 	pa(quoteparse2(`My name is 'Johhnnny' "qhat\'s yours" aaaaaaaa "Ur mum'\'\'" bbb hahaha "us them whatcha gonna do boy" "ending str" \"testing\ my\ quoteparse2"`))
// 	pa(spaceparse(`My name is Johhnnny qhat's\ yours aaaaaaaa Ur\ mum'\\'\\' bbb hahaha us\ them\ whatcha\ gonna\ do\ boy ending\ str`))
// }
