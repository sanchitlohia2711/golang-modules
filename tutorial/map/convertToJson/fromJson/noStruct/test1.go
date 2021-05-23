func test(n int) {
	var char [255]string

	for i := 0; i < 255; i++ {
		char[i] = string(i)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 255; j++ {
			if i%j == 0 {
				fmt.Println(char[j])
			}
		}
	}

}
