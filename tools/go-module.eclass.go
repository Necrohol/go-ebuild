package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    switch os.Getenv("EAPI") {
    case "7", "8":
        // do nothing
    default:
        fmt.Fprintf(os.Stderr, "%s: EAPI %s not supported\n", os.Args[0], os.Getenv("EAPI"))
        os.Exit(1)
    }

    cmd := exec.Command("emerge", "-a", "dev-lang/go")
    if err := cmd.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "%s: failed to install dev-lang/go: %v\n", os.Args[0], err)
        os.Exit(1)
    }

    cmd = exec.Command("emerge", "-a", "app-arch/unzip")
    if err := cmd.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "%s: failed to install app-arch/unzip: %v\n", os.Args[0], err)
        os.Exit(1)
    }

    // TODO: implement src_unpack function

    os.Setenv("GO111MODULE", "on")
    os.Setenv("GOCACHE", os.Getenv("T")+"/go-build")
    os.Setenv("GOMODCACHE", os.Getenv("WORKDIR")+"/go-mod")
    os.Setenv("GOFLAGS", "-buildvcs=false -modcacherw -v -x")
}


// some the first few bits was trying to harvest some to aid in ebuild gneration.. 