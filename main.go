package main

import (
	"io/ioutil"
	"os"
	"log"
	"bytes"

	flag "github.com/spf13/pflag"
)

type Args struct {
	ShowAll         bool
	NumberNonblank  bool
	ShowEnds        bool
	Number          bool
	SqueezeBlank    bool
	ShowTabs        bool
	ShowNonprinting bool
	Help            bool
	Version         bool
}

var args Args

func init() {
	flag.BoolVarP(&args.ShowAll, "show-all", "A", false, "equivalent to -vET")
	flag.BoolVarP(&args.NumberNonblank, "number-nonblank", "b", false, "number nonempty output lines, overrides -n")
	flag.BoolVarP(&args.ShowEnds, "show-ends", "E", false, "display $ at end of each line")
	flag.BoolVarP(&args.Number, "number", "n", false, "number all output lines")
	flag.BoolVarP(&args.SqueezeBlank, "squeeze-blank", "s", false, "suppress repeated empty output lines")
	flag.BoolVarP(&args.ShowTabs, "show-tabs", "T", false, "display TAB characters as ^I")
	flag.BoolVarP(&args.ShowNonprinting, "show-nonprinting", "v", false, "use ^ and M- notation, except for LFD and TAB")
	flag.BoolVarP(&args.Help, "help", "h", false, "display this help and exit")
	flag.BoolVarP(&args.Version, "version", "V", false, "output version information and exit")
	flag.Parse()
}


func main() {
	for _, i := range flag.Args() {
	    b, err := ioutil.ReadFile(i)
        if err != nil {
            log.Fatal(err)
        }
		bytes.replace
		os.Stdout.Write(b)
	}
}
