package main

import (
	"bytes"
	"fmt"

	//"log"
	"strings"
	"testing"
)

var charac = "s"

func BenchmarkConcat(t *testing.B) {
	t.ResetTimer()
	var str string
	for n := 0; n < t.N; n++ {
		str += charac
	}
	t.StopTimer()
	logIt("BenchmarkConcat", str, t.N)
}
func logIt(funcName, str string, n int) {
	//log.Printf("functionName: %v\n", funcName)
	//log.Printf("value of N %v\n", n)
	//log.Printf("str: %v\n", str)
	//log.Printf("--------------\n")
}
func BenchmarkSprintF(t *testing.B) {
	t.ResetTimer()
	var str string
	for n := 0; n < t.N; n++ {
		str = fmt.Sprintf("%v%v", str, charac)
	}
	t.StopTimer()
	logIt("BenchmarkSprintF", str, t.N)
}

func BenchmarkBuffer(t *testing.B) {
	t.ResetTimer()
	var buffer bytes.Buffer
	for n := 0; n < t.N; n++ {
		buffer.WriteString(charac)
	}
	t.StopTimer()
	logIt("BenchmarkBuffer", buffer.String(), t.N)
}

func BenchmarkCopy(t *testing.B) {
	t.ResetTimer()
	byteString := make([]byte, t.N)
	byteLength := 0

	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		byteLength += copy(byteString[byteLength:], charac)
	}
	t.StopTimer()
	logIt("BenchmarkCopy", string(byteString), t.N)
}

func BenchmarkStringBuilder(t *testing.B) {
	t.ResetTimer()
	var strBuilder strings.Builder
	for n := 0; n < t.N; n++ {
		strBuilder.WriteString(charac)
	}
	t.StopTimer()
	logIt("BenchmarkStringBuilder", strBuilder.String(), t.N)
}

func BenchmarkAppendStringArray(t *testing.B) {
	t.ResetTimer()
	var result string
	data := make([]string, 0)
	for n := 0; n < t.N; n++ {
		data = append(data, charac)
	}
	result = strings.Join(data, "")
	t.StopTimer()
	logIt("BenchmarkAppendStringArray", result, t.N)
}

func BenchmarkAppendStringArrayCap(t *testing.B) {
	t.ResetTimer()
	var result string
	data := make([]string, 0, t.N)
	for n := 0; n < t.N; n++ {
		data = append(data, charac)
	}
	result = strings.Join(data, "")
	t.StopTimer()
	logIt("BenchmarkAppendStringArray", result, t.N)
}
