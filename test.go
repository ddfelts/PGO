package main

import (
	"fmt"
	"os"

	"github.com/sbinet/go-python"
)

func doPlugin(e string, b string) string {
	python.Initialize()
	defer python.Finalize()
	fooModule := python.PyImport_ImportModule(e)
	if fooModule == nil {
		panic("Error importing module")
	}
	helloFunc := fooModule.GetAttrString("run")
	if helloFunc == nil {
		panic("Error importing function")
	}
	pystr := python.PyString_FromString(b)
	tup := python.PyTuple_New(1)
	python.PyTuple_SetItem(tup, 0, pystr)
	d := helloFunc.Call(tup, python.PyDict_New())
	ngostr := python.PyString_AsString(d)
	return ngostr
}

func main() {
	os.Setenv("PYTHONPATH", "plugins/")
	d := doPlugin("foo", "[{\"Main\":\"Tester\"}]")
	fmt.Println(d)
}
