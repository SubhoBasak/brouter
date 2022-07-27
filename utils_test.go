package brouter

import "testing"

func TestComparePath(t *testing.T) {
	type tData struct{
		s1 string
		s2 string
		ans int
	}

	testData := [...]tData{
		{"/abc", "/abcdef", 4},
		{"/123", "/123/456", 4},
		{"helloWorld", "HelloWorld", 0},
		{"TESTdata", "TEST", 4},
		{"/common/Different", "/common/different", 8},
	}

	for _, data := range testData {
		node := &Router{val: data.s1, len: uint8(len(data.s1))}

		l1 := uint8(data.ans)
		l2 := comparePath(node, &data.s2, uint8(len(data.s2)))

		if l1 != l2{
			t.Fatalf("Wrong answer. Expect %d got %d", l1, l2)
		}
	}
}

func TestValidatePath(t *testing.T) {
	validPaths := [...]string{
		"/",
		"/category",
		"/customers",
		"/allProducts",
		"/newTrendingTopics",
		"/HappyNewYear2022",
		"/class12",
		"/12345",
	}

	for _, path := range validPaths {
		l1 := uint8(len(path))
		l2 := validatePath(&path)

		if l1 != l2 {
			t.Fatalf("Wrong path length for %s. Expected %d got %d", path, l1, l2)
		}
	}

	invalidPaths := [...]string{
		"",
		"//",
		"customers/",
		"_allProducts",
		"new_TrendingTopics",
		"HappyNewYear@2022",
		"class12",
		"/12345/",
	}

	for _, path := range invalidPaths {
		func() {
			defer func() {
				if err := recover(); err == nil {
					t.Fatalf("%s did not panic", path)
				}
			}()
			validatePath(&path)
		}()
	}
}
