package brouter

import (
	"net/http"
	"testing"
)

func TestAdd(t *testing.T) {
	validData := []string{
		"/category",
		"/customer",
		"/product",
		"/products",
		"/product/all",
		"/product/pid",
		"/a/a/a",
		"/a/a/b",
		"/a/a/c",
		"/a/b/a",
		"/a/b/b",
		"/a/b/c",
		"/a/c/a",
		"/a/c/b",
		"/a/c/c",
	}

	router := Router{}
	for _, data := range validData {
		router.add(data, uint8(len(data)))
	}

	for _, data := range validData {
		node := router.add(data, uint8(len(data)))
		if data[uint8(len(data))-node.len:] != node.val {
			t.Fatalf("Got wrong node for path %s", data)
		}
	}
}

func TestGet(t *testing.T) {
	validData := []string{
		"/category",
		"/customer",
		"/product",
		"/products",
		"/product/all",
		"/product/pid",
		"/a/a/a",
		"/a/a/b",
		"/a/a/c",
		"/a/b/a",
		"/a/b/b",
		"/a/b/c",
		"/a/c/a",
		"/a/c/b",
		"/a/c/c",
	}

	handlerFunc := func(w http.ResponseWriter, r *http.Request) {}

	router := Router{}
	for i, data := range validData {
		node := router.add(data, uint8(len(data)))
		if node.handler == nil {
			node.handler = make(map[byte]*func(http.ResponseWriter, *http.Request))
		}

		switch i % 4 {
		case 0:
			node.handler[GET_METHOD] = &handlerFunc
			node.handler[PUT_METHOD] = &handlerFunc
		case 1:
			node.handler[POST_METHOD] = &handlerFunc
			node.handler[DELETE_METHOD] = &handlerFunc
		case 2:
			node.handler[GET_METHOD] = &handlerFunc
			node.handler[DELETE_METHOD] = &handlerFunc
		case 3:
			node.handler[POST_METHOD] = &handlerFunc
			node.handler[PUT_METHOD] = &handlerFunc
		}
	}

	for i, data := range validData {
		handler := router.get(data, uint8(len(data)))

		switch i % 4 {
		case 0:
			if handler[GET_METHOD] == nil ||
				handler[POST_METHOD] != nil ||
				handler[PUT_METHOD] == nil ||
				handler[DELETE_METHOD] != nil {
				t.Fatalf("Got wrong handler for path %s", data)
			}
		case 1:
			if handler[GET_METHOD] != nil ||
				handler[POST_METHOD] == nil ||
				handler[PUT_METHOD] != nil ||
				handler[DELETE_METHOD] == nil {
				t.Fatalf("Got wrong handler for path %s", data)
			}
		case 2:
			if handler[GET_METHOD] == nil ||
				handler[POST_METHOD] != nil ||
				handler[PUT_METHOD] != nil ||
				handler[DELETE_METHOD] == nil {
				t.Fatalf("Got wrong handler for path %s", data)
			}
		case 3:
			if handler[GET_METHOD] != nil ||
				handler[POST_METHOD] == nil ||
				handler[PUT_METHOD] == nil ||
				handler[DELETE_METHOD] != nil {
				t.Fatalf("Got wrong handler for path %s", data)
			}
		}
	}

	invalidData := []string{
		"/user",
		"/home",
		"/contact",
		"/about",
		"/blog/new",
		"/blog/123",
	}

	for _, data := range invalidData {
		if router.get(data, uint8(len(data))) != nil {
			t.Fatalf("Got wrong answer for %s", data)
		}
	}
}
