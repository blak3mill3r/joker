package main

import (
	"io/ioutil"
	"strings"
)

var coreTemplate string = `// Generated by gen_data. Don't modify manually!

package core

var coreData = []byte("{content}")
`

const hextable = "0123456789abcdef"

func main() {

	coreContent, err := ioutil.ReadFile("data/core.joke")
	if err != nil {
		panic(err)
	}
	dst := make([]byte, len(coreContent)*4)
	for i, v := range coreContent {
		dst[i*4] = '\\'
		dst[i*4+1] = 'x'
		dst[i*4+2] = hextable[v>>4]
		dst[i*4+3] = hextable[v&0x0f]
	}

	ioutil.WriteFile("core_data.go", []byte(strings.Replace(coreTemplate, "{content}", string(dst), 1)), 0666)
}
