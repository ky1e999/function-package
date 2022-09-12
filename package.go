package main

import (
	"fmt"
	"strings"
)

func playfair(kEy string, msG string) string {
	key := strings.Replace(kEy, " ", "", -1)
	msg := strings.Replace(msG, " ", "", -1)
	var ans []byte
	ans = make([]byte, len(msg))
	var arr [][]byte
	arr = make([][]byte, 5)
	letter := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'k',
		'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	Key := norepeat([]byte(key))
	for i := 0; i < 25; i++ {
		if !IsContain(letter[i], Key) {
			Key = append(Key, letter[i])
		}
	}
	for i := 0; i < 5; i++ {
		arr[i] = Key[5*i : 5*i+5]
	}
	Msg := []byte(msg)
	for i := 0; i < len(Msg)-1; i = i + 2 {
		if Msg[i] == Msg[i+1] {
			Msg = append(Msg, Msg[len(Msg)-1])
			for z := len(Msg) - 2; z > i+1; z-- {
				Msg[z] = Msg[z-1]
			}
			Msg[i+1] = 'x'
			i++
		}
	}
	if len(Msg)%2 != 0 {
		Msg = append(Msg, 'x')
	}
	for i := 0; i < len(Msg)-1; i = i + 2 {
		x1, y1 := LocateArr(arr, Msg[i])
		x2, y2 := LocateArr(arr, Msg[i+1])
		if x1 == x2 {
			if y1 == 4 {
				ans = append(ans, arr[x1][0])
			} else {
				ans = append(ans, arr[x1][y1+1])
			}
			if y2 == 4 {
				ans = append(ans, arr[x1][0])
			} else {
				ans = append(ans, arr[x1][y2+1])
			}
		}
		if y1 == y2 {
			if x1 == 4 {
				ans = append(ans, arr[0][y1])
			} else {
				ans = append(ans, arr[x1+1][y1])
			}
			if x2 == 4 {
				ans = append(ans, arr[0][y1])
			} else {
				ans = append(ans, arr[x2+1][y1])
			}
		}
		if (x1 != x2) && (y1 != y2) {
			ans = append(ans, arr[x1][y2])
			ans = append(ans, arr[x2][y1])
		}

	}
	Ans := string(ans[:])
	return Ans
}
func LocateArr(arr [][]byte, ch byte) (int, int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if arr[x][y] == ch {
				return x, y
			}
		}
	}
	return 0, 0
}
func IsContain(ch byte, list []byte) bool {
	for _, x := range list {
		if x == ch {
			return true
		}
	}
	return false
}
func norepeat(str []byte) []byte {
	n := len(str)
	for i := 0; i < n; i++ {
		if str[i] == 'i' || str[i] == 'j' {
			str[i] = 'i'
		}
	}
	for q := 0; q < n; q++ {
		base := str[q]
		for p := q + 1; p < n; p++ {
			if base == str[p] {
				for x := p; x < n-1; x++ {
					str[x] = str[x+1]
				}
				n--
			}
		}
	}
	return str
}
func Locate(list []byte, ch byte) int {
	for i := 0; i < len(list); i++ {
		if ch == list[i] {
			return i
		}
	}
	return len(list)
}
func Vigenere(msG string, kEy string) string {
	key := strings.ToLower(strings.Replace(kEy, " ", "", -1))
	msg := strings.ToLower(strings.Replace(msG, " ", "", -1))
	letter := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k',
		'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	Key := []byte(key)
	Msg := []byte(msg)
	var ans []byte
	for i := 0; i < len(Msg); i++ {
		p := Locate(letter, Msg[i])
		q := Locate(letter, Key[i%len(Key)])
		ans = append(ans, letter[(p+q+1)%26])
	}
	Ans := string(ans[:])
	return Ans
}
func quicksort(str []int, beg int, end int) []int {
	if beg < end {
		key := str[beg]
		i := beg
		j := end
		for {
			if i >= j {
				break
			}
			for {
				if i >= j || str[j] <= key {
					break
				}
				j--
			}
			if i < j {
				str[i] = str[j]
				i++
			}
			for {
				if i >= j || str[i] >= key {
					break
				}
				i++
			}
			if i < j {
				str[j] = str[i]
				j--
			}
		}
		str[i] = key
		quicksort(str, beg, i-1)
		quicksort(str, i+1, end)
	}
	return str
}
func main() {
	key := "playfair is a digram cipher"
	msg := "best"
	fmt.Println(Vigenere(key, msg))
}
