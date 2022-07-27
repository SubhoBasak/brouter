package brouter

import "net/http"

// Add a sub router to the router
func (r *Router) Router(p string, l *Router) {
	s := validatePath(&p)
	if s == 0 {
		panic("Parent router can't be only /")
	}

	node := r.add(p, s)

	if node.childs == nil {
		node.childs = &[]*Router{(*l.childs)[0]}
		return
	}

	for _, c := range *node.childs {
		if c.val == "/" {
			panic("Duplicate path")
		}
	}

	*node.childs = append(*node.childs, (*l.childs)[0])
}

// Add a GET method to the leaf router
func (r *Router) GET(p string, c func(http.ResponseWriter, *http.Request)) {
	s := validatePath(&p)
	node := r.add(p, s)
	if node.handler == nil {
		node.handler = make(map[byte]*func(http.ResponseWriter, *http.Request))
	} else if node.handler[GET_METHOD] != nil {
		panic("Duplicate path")
	}

	node.handler[GET_METHOD] = &c
}

// Add a POST method to the leaf router
func (r *Router) POST(p string, c func(http.ResponseWriter, *http.Request)) {
	s := validatePath(&p)
	node := r.add(p, s)
	if node.handler == nil {
		node.handler = make(map[byte]*func(http.ResponseWriter, *http.Request))
	} else if node.handler[POST_METHOD] != nil {
		panic("Duplicate path")
	}

	node.handler[POST_METHOD] = &c
}

// Add a PUT method to the leaf router
func (r *Router) PUT(p string, c func(http.ResponseWriter, *http.Request)) {
	s := validatePath(&p)
	node := r.add(p, s)
	if node.handler == nil {
		node.handler = make(map[byte]*func(http.ResponseWriter, *http.Request))
	} else if node.handler[PUT_METHOD] != nil {
		panic("Duplicate path")
	}

	node.handler[PUT_METHOD] = &c
}

// Add a DELETE method to the leaf router
func (r *Router) DELETE(p string, c func(http.ResponseWriter, *http.Request)) {
	s := validatePath(&p)
	node := r.add(p, s)
	if node.handler == nil {
		node.handler = make(map[byte]*func(http.ResponseWriter, *http.Request))
	} else if node.handler[DELETE_METHOD] != nil {
		panic("Duplicate path")
	}

	node.handler[DELETE_METHOD] = &c
}
