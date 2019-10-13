package LeetCode

func isIsomorphic1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashMap := make(map[byte]byte)
	map_t := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, ok := map_t[t[i]]; ok && v != s[i] {
			return false
		}
		hashMap[s[i]] = t[i]
		map_t[t[i]] = s[i]
	}

	var tmp []byte
	for i := 0; i < len(s); i++ {
		v, ok := hashMap[s[i]]
		if !ok {
			return false
		} else {
			tmp = append(tmp, v)
		}
	}

	if string(tmp) != t {
		return false
	}
	return true
}

func isIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashMap := make(map[byte]byte, 256)
	mapped := make(map[byte]bool, 256)

	for i := 0; i < len(s); i++ {
		if hashMap[s[i]] == 0 {
			if mapped[t[i]] {
				return false
			}
			hashMap[s[i]] = t[i]
			mapped[t[i]] = true
		} else if hashMap[s[i]] != t[i] {
			return false
		}
	}
	return true
}
