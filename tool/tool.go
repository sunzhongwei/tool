/*
Copyright 2013 Zhongwei Inc.

Author: Zhongwei Sun
EMail:  zhongwei.sun2008@gmail.com
*/

package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static/tool/")))
	http.ListenAndServe(":8888", nil)
}
