package cmd

import "main/modules/users"

var modules = make([]func(), 0)

func init() {
	New(users.Init)
}

func New(fn func()) {

	modules = append(modules, fn)
}

func Execute() {
	for _, fn := range modules {
		fn()
	}
}
