package ascii

func MakeMap(s string) map[rune]string {
	simvol := make(map[rune]string)
	temp := ""
	count := 0
	j := rune(32)
	for _, data := range s {
		temp += string(data)
		if data == '\n' {
			count++
		}
		if count == 9 {
			count = 0
			simvol[j] = temp[1 : len(temp)-1]
			temp = ""
			j++
		}
	}
	return simvol
}
