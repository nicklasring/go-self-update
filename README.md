# go-self-update

A program that can update itself by replacing its executable file with a new version.

Using `syscall.Exec` to replace parent process so no new childs are spawned.

# Installation
go get http://github.com/nicklasring/go-self-update

make (Generates two binaries with different version numbers)

./go-self-update (Replaces Version 1 with Version 2)
