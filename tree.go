package brouter

import "net/http"

type Router struct {
	val     string
	len     uint8
	handler map[byte]*func(http.ResponseWriter, *http.Request)
	childs  *[]*Router
}

// and new path to the router and return the
// router node which is responsible for the
// given path
func (r *Router) add(s string, l uint8) *Router {
loop:
	if r.childs == nil {
		tmp := &Router{val: s, len: l}
		r.childs = &[]*Router{tmp}
		return tmp
	}

	for _, c := range *r.childs {
		i := comparePath(c, &s, l)
		if i == 0 {
			continue
		}

		if i == c.len {
			if i == l {
				return c
			}
			r = c
			s = s[i:]
			l -= i
			goto loop
		} else if i == l {
			c.childs = &[]*Router{{val: c.val[i:], len: c.len - i, handler: c.handler, childs: c.childs}}
			c.val = s
			c.len = l
			c.handler = nil
			return c
		}
		tmp := &Router{val: s[i:], len: l - i}
		c.childs = &[]*Router{{val: c.val[i:], len: c.len - i, handler: c.handler, childs: c.childs}, tmp}
		c.val = c.val[:i]
		c.len = i
		c.handler = nil
		return tmp
	}

	tmp := &Router{val: s, len: l}
	*r.childs = append(*r.childs, tmp)
	return tmp
}

func (r *Router) get(s string, l uint8) map[byte]*func(http.ResponseWriter, *http.Request) {
loop:
	if r.childs == nil {
		return nil
	}

	for _, c := range *r.childs {
		i := comparePath(c, &s, l)
		if i == 0 || c.len != i {
			continue
		}

		l = l - i
		if l == 0 {
			return c.handler
		}
		s = s[i:]
		r = c
		goto loop
	}

	return nil
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l := len(r.URL.Path)
	if l > 255 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	methodMap := router.get(r.URL.Path, uint8(l))
	if methodMap == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	method := INVALID_METHOD

	switch r.Method {
	case "GET":
		method = GET_METHOD
	case "POST":
		method = POST_METHOD
	case "PUT":
		method = PUT_METHOD
	case "DELETE":
		method = DELETE_METHOD
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	if methodMap[method] == nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	(*methodMap[method])(w, r)
}
