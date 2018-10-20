package main

import (
	"C"
)

//export Sum
func Sum(a, b int) int {
	return a + b
}

// $ go help
// Go is a tool for managing Go source code.

// Usage:

//         go command [arguments]

// The commands are:

//         build       compile packages and dependencies
//         clean       remove object files
//         doc         show documentation for package or symbol
//         env         print Go environment information
//         bug         start a bug report
//         fix         run go tool fix on packages
//         fmt         run gofmt on package sources
//         generate    generate Go files by processing source
//         get         download and install packages and dependencies
//         install     compile and install packages and dependencies
//         list        list packages
//         run         compile and run Go program
//         test        test packages
//         tool        run specified go tool
//         version     print Go version
//         vet         run go tool vet on packages

// Use "go help [command]" for more information about a command.

// Additional help topics:

//         c           calling between Go and C
//         buildmode   description of build modes
//         filetype    file types
//         gopath      GOPATH environment variable
//         environment environment variables
//         importpath  import path syntax
//         packages    description of package lists
//         testflag    description of testing flags
//         testfunc    description of testing functions

// Use "go help [topic]" for more information about that topic.

// $ go help build
// usage: go build [-o output] [-i] [build flags] [packages]

// Build compiles the packages named by the import paths,
// along with their dependencies, but it does not install the results.

// If the arguments to build are a list of .go files, build treats
// them as a list of source files specifying a single package.

// When compiling a single main package, build writes
// the resulting executable to an output file named after
// the first source file ('go build ed.go rx.go' writes 'ed' or 'ed.exe')
// or the source code directory ('go build unix/sam' writes 'sam' or 'sam.exe').
// The '.exe' suffix is added when writing a Windows executable.

// When compiling multiple packages or a single non-main package,
// build compiles the packages but discards the resulting object,
// serving only as a check that the packages can be built.

// When compiling packages, build ignores files that end in '_test.go'.

// The -o flag, only allowed when compiling a single package,
// forces build to write the resulting executable or object
// to the named output file, instead of the default behavior described
// in the last two paragraphs.

// The -i flag installs the packages that are dependencies of the target.

// The build flags are shared by the build, clean, get, install, list, run,
// and test commands:

//         -a
//                 force rebuilding of packages that are already up-to-date.
//         -n
//                 print the commands but do not run them.
//         -p n
//                 the number of programs, such as build commands or
//                 test binaries, that can be run in parallel.
//                 The default is the number of CPUs available.
//         -race
//                 enable data race detection.
//                 Supported only on linux/amd64, freebsd/amd64, darwin/amd64 and windows/amd64.
//         -msan
//                 enable interoperation with memory sanitizer.
//                 Supported only on linux/amd64,
//                 and only with Clang/LLVM as the host C compiler.
//         -v
//                 print the names of packages as they are compiled.
//         -work
//                 print the name of the temporary work directory and
//                 do not delete it when exiting.
//         -x
//                 print the commands.

//         -asmflags 'flag list'
//                 arguments to pass on each go tool asm invocation.
//         -buildmode mode
//                 build mode to use. See 'go help buildmode' for more.
//         -compiler name
//                 name of compiler to use, as in runtime.Compiler (gccgo or gc).
//         -gccgoflags 'arg list'
//                 arguments to pass on each gccgo compiler/linker invocation.
//         -gcflags 'arg list'
//                 arguments to pass on each go tool compile invocation.
//         -installsuffix suffix
//                 a suffix to use in the name of the package installation directory,
//                 in order to keep output separate from default builds.
//                 If using the -race flag, the install suffix is automatically set to race
//                 or, if set explicitly, has _race appended to it. Likewise for the -msan
//                 flag. Using a -buildmode option that requires non-default compile flags
//                 has a similar effect.
//         -ldflags 'flag list'
//                 arguments to pass on each go tool link invocation.
//         -linkshared
//                 link against shared libraries previously created with
//                 -buildmode=shared.
//         -pkgdir dir
//                 install and load all packages from dir instead of the usual locations.
//                 For example, when building with a non-standard configuration,
//                 use -pkgdir to keep generated packages in a separate location.
//         -tags 'tag list'
//                 a space-separated list of build tags to consider satisfied during the
//                 build. For more information about build tags, see the description of
//                 build constraints in the documentation for the go/build package.
//         -toolexec 'cmd args'
//                 a program to use to invoke toolchain programs like vet and asm.
//                 For example, instead of running asm, the go command will run
//                 'cmd args /path/to/asm <arguments for asm>'.

// All the flags that take a list of arguments accept a space-separated
// list of strings. To embed spaces in an element in the list, surround
// it with either single or double quotes.

// For more about specifying packages, see 'go help packages'.
// For more about where packages and binaries are installed,
// run 'go help gopath'.
// For more about calling between Go and C/C++, run 'go help c'.

// Note: Build adheres to certain conventions such as those described
// by 'go help gopath'. Not all projects can follow these conventions,
// however. Installations that have their own conventions or that use
// a separate software build system may choose to use lower-level
// invocations such as 'go tool compile' and 'go tool link' to avoid
// some of the overheads and design decisions of the build tool.

// See also: go install, go get, go clean.

// $ go help buildmode
// The 'go build' and 'go install' commands take a -buildmode argument which
// indicates which kind of object file is to be built. Currently supported values
// are:

//         -buildmode=archive
//                 Build the listed non-main packages into .a files. Packages named
//                 main are ignored.

//         -buildmode=c-archive
//                 Build the listed main package, plus all packages it imports,
//                 into a C archive file. The only callable symbols will be those
//                 functions exported using a cgo //export comment. Requires
//                 exactly one main package to be listed.

//         -buildmode=c-shared
//                 Build the listed main package, plus all packages it imports,
//                 into a C shared library. The only callable symbols will
//                 be those functions exported using a cgo //export comment.
//                 Requires exactly one main package to be listed.

//         -buildmode=default
//                 Listed main packages are built into executables and listed
//                 non-main packages are built into .a files (the default
//                 behavior).

//         -buildmode=shared
//                 Combine all the listed non-main packages into a single shared
//                 library that will be used when building with the -linkshared
//                 option. Packages named main are ignored.

//         -buildmode=exe
//                 Build the listed main packages and everything they import into
//                 executables. Packages not named main are ignored.

//         -buildmode=pie
//                 Build the listed main packages and everything they import into
//                 position independent executables (PIE). Packages not named
//                 main are ignored.

//         -buildmode=plugin
//                 Build the listed main packages, plus all packages that they
//                 import, into a Go plugin. Packages not named main are ignored.

// go build -buildmode=c-shared -o libcommon.so test/testcgo/common
func main() {

}
