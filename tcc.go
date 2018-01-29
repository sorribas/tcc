package tcc

// #cgo LDFLAGS: libtcc.a -ldl
// #include "libtcc.h"
//
// typedef void (*golang_tcc_error_func_type) (void*, const char*);
// char* golang_tcc_error_message;
//
// void golang_tcc_error_func (void* opaque, char *msg) {
//   golang_tcc_error_message = msg;
// }
import "C"
import "errors"

type Tcc struct {
	ctcc *C.TCCState
}

// output will be run in memory (default)
const OUTPUT_MEMORY = 1

// executable file
const OUTPUT_EXE = 2

// dynamic library
const OUTPUT_DLL = 3

// object file
const OUTPUT_OBJ = 4

// only preprocess (used internally)
const OUTPUT_PREPROCESS = 5

// create a new TCC compilation context
func NewTcc() *Tcc {
	ctcc := C.tcc_new()
	C.tcc_set_error_func(ctcc, nil, C.golang_tcc_error_func_type(C.golang_tcc_error_func))
	r := &Tcc{ctcc}
	return r
}

// free a TCC compilation context
func (tcc *Tcc) Delete() {
	C.tcc_delete(tcc.ctcc)
}

// set CONFIG_TCCDIR at runtime
func (tcc *Tcc) SetLibPath(path string) {
	C.tcc_set_lib_path(tcc.ctcc, C.CString(path))
}

// set options as from command line (multiple supported)
func (tcc *Tcc) SetOptions(opts string) {
	C.tcc_set_options(tcc.ctcc, C.CString(opts))
}

// compile a string containing a C source.
func (tcc *Tcc) CompileString(src string) error {
	C.golang_tcc_error_message = nil
	r := C.tcc_compile_string(tcc.ctcc, C.CString(src))
	if r == -1 {
		return errors.New(C.GoString(C.golang_tcc_error_message))
	}

	return nil
}

// set output type. MUST BE CALLED before any compilation
func (tcc *Tcc) SetOutputType(outputType int) {
	C.tcc_set_output_type(tcc.ctcc, C.int(outputType))
}

// output an executable, library or object file.
func (tcc *Tcc) OutputFile(path string) {
	C.tcc_output_file(tcc.ctcc, C.CString(path))
}

// link and run main() function and return its value.
func (tcc *Tcc) Run(argv []string) (int, error) {
	cstrings := []*C.char{}
	for _, arg := range argv {
		cstrings = append(cstrings, C.CString(arg))
	}

	C.golang_tcc_error_message = nil
	var r int
	if len(cstrings) > 0 {
		r = int(C.tcc_run(tcc.ctcc, C.int(len(cstrings)), &cstrings[0]))
	} else {
		r = int(C.tcc_run(tcc.ctcc, C.int(0), nil))
	}
	if r == -1 && C.golang_tcc_error_message != nil {
		return r, errors.New(C.GoString(C.golang_tcc_error_message))
	}

	return r, nil
}
