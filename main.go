package main

import (
	 "log"
	 "net/http"
	 "path"
	 "strings"
	 "strconv"
	 "encoding/json"
)

func main() {
	if err := http.ListenAndServe(":7070", http.HandlerFunc(getPrimefactors)); err != nil {
		log.Fatal(err)
	}
}

func getPrimefactors(res http.ResponseWriter, req *http.Request) {
	log.Println("request received")
	var head string
	var pfs []int

	head, req.URL.Path = shiftPath(req.URL.Path)

	var n int
	var err error
	if n, err = strconv.Atoi(head); err != nil {
		http.Error(res, "not an integer: " + head, http.StatusOK)
	}

	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	js, err := json.Marshal(pfs)

	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}