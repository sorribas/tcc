// Package tcc offers go bindings to libtcc.
//
// https://bellard.org/tcc/
//
// The libtcc library allows you to call the tcc
// compiler programatically. It is available as a
// static C library. This package has go bindings
// to it.
//
// It doesn't implement every method available in
// libtcc yet. If there's a libtcc function that
// you need and is not implemented in this package,
// feel free to open an issue or submit a pull request.
//
// To use it, you must have a libtcc.a binary in the
// cwd while building. You can download tcc in the link
// above, and after compiling it, you'll get a libtcc.a
// for your platform.
package tcc
