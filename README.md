# tcc

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/sorribas/tcc)

Go bindings for [libtcc](https://bellard.org/tcc/).

## Usage

Here's how to compile a C program in memory and run it.

```go
package main

import "github.com/sorribas/tcc"

func main () {
	program := `
	extern int printf(char* c,...);
	int main () {
		printf("Hello world!\n");
		return 0;
	}
	`
	tcc := NewTcc()
	defer tcc.Delete()
	tcc.SetOutputType(OUTPUT_MEMORY)

	// tcc must contain the libtcc1.a file that includes the TCC runtime.
	tcc.SetLibPath("/path/to/tcc")
	err := tcc.CompileString(program)
	if err != nil {
		// compilation error
	}
	// run the program
	i := tcc.Run([]string{})
}
```

You can use this library as a backend for a compiler or simply
to compile and run dynamically generated code.

To build a project using this library, you need to have a `libtcc.a`
for your platform in the cwd at build time.

## License

MIT License
