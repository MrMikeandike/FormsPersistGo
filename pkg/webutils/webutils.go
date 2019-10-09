package webutils

//ShiftPath returns first path element and tail of path
//the head will will never start with a forward slash
//tail will always begin with a forward slash and have no trailing slash
func ShiftPath(p string) (head string, tail string) {
	if p == "" {
		return "", ""
	}

	b := []byte(p)
	var r int = 0
	var n int = len(b)
	if b[0] == '/' {
		if n == 1 {
			return "/", ""
		}
		r = 1
	}
	if b[n-1] == '/' {
		n--
	}

	for i := r; n > i; i++ {
		if b[i] == '/' {
			return string(b[r:i]), string(b[i:n])
		}

	}

	return string(b[r:n]), ""
}
