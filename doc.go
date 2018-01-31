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
// libtcc yet. If there's a libtcc method that
// you need and is not implemented in this package,
// feel free to open an issue or submit a pull request.
//
// To use it you must have libtcc installed. You can
// download tcc from the website above. After that you
// can extract and run
//
//  ./configure
//  make
//  make install
//
// If everything works, you will have libtcc installed.
// Note that you need a C compiler like gcc for the above
// to work.
package tcc
