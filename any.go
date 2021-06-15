package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unsafe"
)

func main() {

	f, err := os.Open("any")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	type T struct {
		A int
		B uint8
		C uint16
		D int8
		E int16
		F int32
		G int64
		K float32
		L float64
		M complex64
		N complex128
		O bool
		P string
	}

	type kv struct {
		Key   string
		Value uintptr
	}

	t := T{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var array [10][]string
	for index, element := range lines {
		array[index] = strings.Split(element, " ")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m := make(map[string]uintptr)
	for i := 3; i < 9; i++ {
		var a uintptr
		switch array[i][2] {
		case "uint", "int", "*uint", "*int":
			a = unsafe.Sizeof(t.A)
			break
		case "uint8", "byte", "*uint8", "*byte":
			a = unsafe.Sizeof(t.B)
			break
		case "uint16", "*uint16":
			a = unsafe.Sizeof(t.C)
			break
		case "int8", "*int8":
			a = unsafe.Sizeof(t.D)
			break
		case "int16", "*int16":
			a = unsafe.Sizeof(t.E)
			break
		case "int32", "rune", "*int32", "*rune":
			a = unsafe.Sizeof(t.F)
			break
		case "int64", "*int64":
			a = unsafe.Sizeof(t.G)
			break
		case "float32", "*float32":
			a = unsafe.Sizeof(t.K)
			break
		case "float64", "*float64":
			a = unsafe.Sizeof(t.L)
			break
		case "complex64", "*complex64":
			a = unsafe.Sizeof(t.M)
			break
		case "complex128", "*complex128":
			a = unsafe.Sizeof(t.N)
			break
		case "bool", "*bool":
			a = unsafe.Sizeof(t.O)
			break
		case "string", "*string":
			a = unsafe.Sizeof(t.P)
			break
		}
		m[array[i][2]] = a
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	var result []string
	var maxSize = ss[0].Value
	var end = 1
	for index, kv := range ss {

		result = append(result, kv.Key)
		var sum = kv.Value
		var beforeEnd = 1
		for i := len(ss) - end; i >= index; i-- {
			if sum+ss[i].Value > maxSize {
				break
			} else {
				sum += ss[i].Value
				result = append(result, ss[i].Key)
			}
			beforeEnd++
		}
		end = beforeEnd

		if index == (len(ss) / 2)-1 {
			break
		}

	}

	fmt.Println(result)

}
