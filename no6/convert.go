package no6

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	// i的取值范围：0~numRows-1
	// flag: -1 | 1
	i, flag := 0, -1
	rows := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		rows[i] = make([]byte, 0, len(s)/2)
	}
	for _, c := range []byte(s) {
		rows[i] = append(rows[i], c)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}

	d := make([]byte, 0, len(s))
	for _, row := range rows {
		d = append(d, row...)
	}

	return string(d)
}
