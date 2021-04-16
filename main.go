package main

import(
    "fmt"
    "flag"
)


func showHelp() {
    fmt.Println(`

Usage: CLI Template [OPTIONS]

Options:
    -s, --string     Prints string input.       (default='some string').
    -i, --int        Prints int.                 (defualt=0).
    -e, --error      Prints a costum made error.
    -w, --warning    Prints the warning you entered.

    -h, --help       prints this help info.


    `)
}


func setFlag(flag *flag.FlagSet) {
    flag.Usage = func() {
        showHelp()
    }
}

func main() {
    var sHelp bool
    flag.BoolVar(&sHelp, "h", false, "")
    flag.BoolVar(&sHelp, "help", false, "")

    setFlag(flag.CommandLine)

    flag.Parse()

    if sHelp {
        showHelp()
        return
    }


}