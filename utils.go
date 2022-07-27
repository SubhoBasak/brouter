package brouter

import "fmt"

const (
	INVALID_METHOD = byte(iota)
	GET_METHOD
	POST_METHOD
	PUT_METHOD
	DELETE_METHOD
)

// compare the strings of router and the given path
// and return the length of common starting portion
func comparePath(n *Router, s *string, l uint8) uint8 {
	mx := n.len
	if mx > l {
		mx = l
	}

	for i := uint8(0); i < mx; i++ {
		if n.val[i] != (*s)[i] {
			return i
		}
	}

	return mx
}

// check weather the given string is a valid path or
// or not and return the lenth of the string if it is
// a valid path
func validatePath(s *string) uint8 {
	if *s == "" {
		panic("Path can't be empty")
	} else if (*s)[0] != '/' {
		panic("Path must start with /")
	}
	l := len(*s)
	if l > 255 {
		panic(fmt.Sprintf("Path %s length greater than 255", *s))
	}

	for _, c := range (*s)[1:] {
		if c < '0' || (c > '9' && c < 'A') || (c > 'Z' && c < 'a') || c > 'z' {
			panic(fmt.Sprintf("Path %s contains invalid character", *s))
		}
	}

	return uint8(l)
}
