package main

import (
	"flag"
	"fmt"
	"github.com/Han-Ya-Jun/ini2go"
	"os"
)

/*
* @Author:hanyajun
* @Date:2019/7/27 23:13
* @Name:ini2go
* @Function:
 */
// handle flags that are string arrays
type stringArr []string

// needed to fulfill flag.Value
func (s *stringArr) String() string {
	var tmp string
	for i, v := range *s {
		tmp += v
		if i > 0 && i < len(*s) {
			tmp += " "
		}
	}
	return tmp
}

// Add the flag value to the slice.
func (s *stringArr) Set(v string) error {
	*s = append(*s, v)
	return nil
}

// Get the flag array as a slice.
func (s stringArr) Get() []string {
	return s
}

var (
	iniFileName string
	pkgName     string
	outputPath  string
	goFileName  string
	writeTag    bool
	tagKeys     stringArr
	help        bool
)

func init() {
	flag.StringVar(&iniFileName, "inifile", "app.ini", "the path of ini file")
	flag.StringVar(&iniFileName, "i", "", "the short flag for -inifile")
	flag.StringVar(&outputPath, "output", ".", "the path to the output file")
	flag.StringVar(&outputPath, "o", "stdin", "the short flag for -input")
	flag.StringVar(&goFileName, "gofilename", "default:the name of inifile", "path to the output file")
	flag.StringVar(&goFileName, "g", "default:the name of inifile", "the short flag for -gofilename")
	flag.StringVar(&pkgName, "pkgname", "", "the name of the package")
	flag.StringVar(&pkgName, "p", "", "the short flag for -pkg")
	flag.BoolVar(&writeTag, "writetag", false, "true: write tag")
	flag.BoolVar(&writeTag, "w", false, "the short flag for -writejson")
	flag.Var(&tagKeys, "tagkeys", "additional struct tag keys; can be used more than once")
	flag.Var(&tagKeys, "t", "the short flag for -tagkeys")
	flag.BoolVar(&help, "help", false, "ini2go help")
	flag.BoolVar(&help, "h", false, "the short flag for -help")
}
func main() {
	os.Exit(realMain())
}
func realMain() int {
	flag.Parse()
	args := flag.Args()
	// the only arg we care about is help.  This is in case the user uses
	// just help instead of -help or -h
	for _, arg := range args {
		if arg == "help" {
			help = true
			break
		}
	}
	if help {
		Help()
		return 0
	}
	if iniFileName == "" {
		fmt.Fprintln(os.Stderr, "\nthe ini file must be given")
		return 1
	}
	err := ini2go.Ini2Go(iniFileName, pkgName, goFileName, outputPath, writeTag, tagKeys.Get())
	if err != nil {
		fmt.Println(err)
		return 0
	}
	fmt.Println("ini2go success")
	return 0
}

func Help() {
	helpText := `
Usage: ini2go [options]
Ini2go is used to generate go struct

Options:

flag              default   description
---------------   -------   ------------------------------------------
-i  -inifile      app.ini   The path of ini file: required.
-o  -output       .         The Go srouce code output destination.
-p  -pkg          app       The name of the package.
-g  -gofilename   inifilename  The name of the go file.
-h  -help         false     Print the help text; 'help' is also valid.
-t  -tagkey                 Additional key to be added to struct tags.
                            For multiple keys, use one per key value.
-wt -writetag     fale      Write tag
`
	fmt.Println(helpText)
}
