package main

import "fmt"

func main() {
	anagram := []string{"kita", "atik", "tika", "aku", "kua", "kia", "makan"}

	printAnagram(anagram)
}

func printAnagram(anagram []string) {
	anagramGroup := make([][]string, 0)
	anagramCpy := make([]string, 0)
	anagramMap := make(map[string][]int, 0)
	for _, v := range anagram {
		value := []byte(v)
		quickSort(value, 0, len(value)-1)
		anagramCpy = append(anagramCpy, string(value))
	}

	for k, v := range anagramCpy {
		if _, ok := anagramMap[v]; !ok {
			anagramMap[v] = make([]int, 0)
		}
		anagramMap[v] = append(anagramMap[v], k)
	}

	for _, v := range anagramMap {
		group := make([]string, 0)
		for _, i := range v {
			group = append(group, anagram[i])
		}
		anagramGroup = append(anagramGroup, group)
	}

	fmt.Println(anagramGroup)
}

func quickSort(characters []byte, low, high int) {
	if low < high {
		ref := partition(characters, low, high)
		quickSort(characters, low, ref)
		quickSort(characters, ref+1, high)
	}
}

func partition(data []byte, low, high int) int {

	pivot := data[low]
	i := low + 1

	j := high
	for ; ; j-- {
		if data[j] > pivot {
			continue
		}

		for {

			if i > j || data[i] >= pivot {
				break
			}
			i++
		}

		if j >= i {
			swap(&data[i], &data[j])
		} else {
			break
		}
	}
	if j > low {
		swap(&data[low], &data[j])
	}
	return j
}

func swap(i, j *byte) {
	t := *i
	*i = *j
	*j = t
}
