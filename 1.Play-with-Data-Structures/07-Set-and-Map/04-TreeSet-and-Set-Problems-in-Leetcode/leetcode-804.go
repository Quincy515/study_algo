package _4_TreeSet_and_Set_Problems_in_Leetcode

import "bytes"

func uniqueMorseRepresentations(words []string) int {
	codes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	buffer := bytes.Buffer{}
	uniqueWord := make(map[string]bool)
	for _, word := range words {
		buffer.Reset()
		for _, letter := range word {
			buffer.WriteString(codes[letter-'a'])
		}
		uniqueWord[buffer.String()] = true
	}
	return len(uniqueWord)
}

func uniqueMorseRepresentations1(words []string) int {
	var dict = [...]string{".-", "-...", "-.-.", "-..", ".", "..-.",
		"--.", "....", "..", ".---", "-.-", ".-..", "--", "-.",
		"---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-",
		".--", "-..-", "-.--", "--.."}
	var res = make(map[string]string)
	var conv = func(S string) string {
		var rets = ""
		for _, v := range S {
			rets += dict[v-'a']
		}
		return rets
	}
	for _, v := range words {
		res[conv(v)] = v
	}

	return len(res)
}

func uniqueMorseRepresentations2(words []string) int {
	table := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	hash := map[string]bool{}
	//cnt := 0
	for _, w := range words {
		t := ""
		for _, x := range w {
			t += table[x-'a']
		}
		//if !hash[t] {
		//cnt++
		hash[t] = true
		//}
	}
	return len(hash)
}
