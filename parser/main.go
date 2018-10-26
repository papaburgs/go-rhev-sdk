/*Package parser parses a red hat documentation page

I couldn't find a usable version of the api in xml format so wrote this
`little` script. I don't plan to use it often. I wrote it for version 4.0 and 4.2.

Usage

Assuming this repo has been cloned to your local machine, go to this directory
Find the version of the documentation you need and save the Types section as a single page html file.
Do not need anything more than the html.

Put that file in this directory, make sure it is named `types.html`
Then run this command: `go run main.go`

Output goes to stdout, you can redirect that to a file, ie, `go run main.go > types.go.generated`

Issues

I do not know if this will be of any use. Seems even on the first run, there are differences between
the generated code and the actual output API.
If you find a better way to do this, let me know.

*/
package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Sirupsen/logrus"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type dataAttributes struct {
	name    string
	section string
	attrs   []attribute
}

// String outputs a go friendly version of structs.
// This does not do proper formatting, I let the gofmt tool do that
func (d dataAttributes) String() string {
	var out string
	out = fmt.Sprintf("// type %s from section %s of the Redhat api guide\ntype %s struct {\n",
		d.name, d.section, d.name)
	for _, x := range d.attrs {
		out = out + fmt.Sprintf("%s %s `json:\"%s\"`\n", x.goName, x.attType, x.jsonName)
	}
	out = out + "}"
	return out
}

type attribute struct {
	jsonName string
	attType  string
	goName   string
}

var (
	// ErrEOF represents and error that we made it to the end of the file
	ErrEOF = errors.New("EOF")
	// ErrFound is passed back when something was found that it shouldn't have
	ErrFound = errors.New("Found")
	// ErrNoName is passed when a name isn't found
	ErrNoName = errors.New("Could find name in expected location")
	// ErrBadTable is used when a table doesn't match was is expectec
	ErrBadTable = errors.New("Invalid table")
)

func nextToken(z *html.Tokenizer) (html.Token, error) {
	tt := z.Next()
	if tt == html.ErrorToken {
		return html.Token{}, ErrEOF
	}
	return z.Token(), nil
}

// return is name, section, error
func findNameInSection(z *html.Tokenizer) (string, string, error) {
	for {
		t, err := nextToken(z)
		if err != nil {
			return "", "", err
		}
		if t.Type == html.StartTagToken && t.DataAtom == atom.H2 {
			// next token is the content we want
			t, err = nextToken(z)
			if err != nil {
				return "", "", err
			}
			// the token Data will have content like "6.233 Weight"
			// want the number is one variable and the rest in another
			re := regexp.MustCompile(`(\d\.\d+)[\W]+(.*)\s*`)
			match := re.FindStringSubmatch(t.Data)
			if len(match) != 3 {
				return "", "", ErrNoName
			}
			return strings.Trim(match[2], " "),
				strings.Trim(match[1], " "),
				nil

		}
	}
}

// function findOpenTag advances the tokenizer until if finds what we want
func findOpenTag(z *html.Tokenizer, tag atom.Atom) error {
	for {
		t, err := nextToken(z)
		if err != nil {
			return err
		}
		if t.Type == html.StartTagToken && t.DataAtom == tag {
			return nil
		}
	}
}

// same as findOpenTag, but will stop if it finds the before
// use case is for looking for tr tags, we can give up if we see a tbody or table tag
func findOpenTagBefore(z *html.Tokenizer, tag atom.Atom, before atom.Atom) error {
	for {
		t, err := nextToken(z)
		if err != nil {
			return err
		}
		if t.Type == html.StartTagToken && t.DataAtom == tag {
			return nil
		}
		if t.DataAtom == before {
			return ErrFound
		}

	}
}

// function convertAttr is a stringy mess that normalizes fields
// this includes changing munging names for Types and making
// names that are Go compatible instead of snakey case
func convertAttr(in attribute) attribute {
	var (
		prepend bool
		y       string
		a       attribute
	)
	a = attribute{
		jsonName: strings.Trim(in.jsonName, " "),
	}
	prepend = false
	// fix type
	y = strings.Trim(in.attType, " ")
	if len(y) == 0 {
		// this can happen on those 2 column tables, just set to strings to be safe
		y = "string"
	}
	if len(y) > 2 && y[len(y)-2:] == "[]" {
		prepend = true
		y = y[:len(y)-2]
	}
	if y == "Boolean" {
		y = "bool"
	}
	if y == "Integer" {
		y = "int"
	}
	if y == "String" {
		y = "string"
	}
	if prepend {
		y = "[]" + y
	}
	a.attType = y

	// convert json name to a golang name
	y = ""
	capNext := true
	for _, v := range in.jsonName {
		switch {
		case v >= 'A' && v <= 'Z':
			y += string(v)
			capNext = false

		case v >= '0' && v <= '9':
			y += string(v)
			capNext = true

		case v >= 'a' && v <= 'z':
			if capNext {
				y += strings.ToUpper(string(v))
				capNext = false
			} else {
				y += string(v)
			}
		case v == ' ' || v == '-' || v == '_':
			capNext = true
		}
	}
	a.goName = y
	return a
}

// function analyzeTable looks at the headers to see how many columns
// expecting either 2 or 3, if we have something different, stuff
// will likely fail
func analyzeTable(z *html.Tokenizer) int {
	headers := 0
	err := findOpenTag(z, atom.Thead)
	if err != nil {
		return headers
	}
	_ = findOpenTag(z, atom.Tr)
	for {
		err = findOpenTagBefore(z, atom.Th, atom.Tr)
		if err != nil {
			break
		}
		headers++
	}
	return headers

}

// function mapTable was the worst to build
// it is not very extensible and really iterative, nothing elegant about it
// but it works
func mapTable(z *html.Tokenizer, cols int) ([]attribute, error) {
	var err error
	var a attribute
	var atts []attribute
	if cols < 2 || cols > 3 {
		// only have logic in place for tables with 2 or three columns
		return atts, ErrBadTable
	}
	// at this point we are in a tbody, so we parse rows until we find the end of tbody
	for {
		a = attribute{}
		// find start of row
		err := findOpenTagBefore(z, atom.Tr, atom.Tbody)
		if err != nil {
			if err == ErrFound {
				// reset error, as this is ok, just no more rows
				err = nil
			}
			break
		}
		// find first cell in the row, this is a code literal with json name
		err = findOpenTag(z, atom.Td)
		if err != nil {
			break
		}
		// find that code tag
		err = findOpenTag(z, atom.Code)
		if err != nil {
			break
		}
		t, _ := nextToken(z)
		a.jsonName = t.Data
		if cols == 2 { // coninue to next row unless I want to pull in summary
			atts = append(atts, convertAttr(a))
			continue
		}
		// Now look for the type attribute (which doesn't exist if there is only 2 cols)
		err = findOpenTag(z, atom.Td)
		if err != nil {
			break
		}
		// it is in a p block, in and an anchor block (it is a link)
		err = findOpenTag(z, atom.A)
		if err != nil {
			break
		}
		t, _ = nextToken(z)
		a.attType = t.Data
		l.Debug("appending: ", convertAttr(a))
		atts = append(atts, convertAttr(a))
	}
	return atts, err
}

// global logger - deal with it
var l = logrus.New()

func main() {
	var (
		da  dataAttributes
		err error
	)
	//l.SetLevel(logrus.DebugLevel)

	// hard coded file name = Greatness!
	typeFile, err := os.Open("types.html")
	if err != nil {
		l.Error(err)
		return
	}
	z := html.NewTokenizer(typeFile)

	for {
		da = dataAttributes{}
		err = findOpenTag(z, atom.Section)
		if err != nil {
			break
		}
		da.name, da.section, err = findNameInSection(z)
		if err != nil {
			l.Error("spot2 ", err)
			return
		}
		l.Debug("have: ", da.name, " - ", da.section)

		cols := analyzeTable(z)
		if cols == 0 {
			l.Error("problems finding table")
		}

		err = findOpenTag(z, atom.Tbody)
		if err != nil {
			l.Error("premature ending")
			return
		}
		da.attrs, err = mapTable(z, cols)
		if err != nil {
			l.Error("spot3 ", err)
			return
		}
		// each loop gives us a section, so just print it out, could
		// save this into a slice and do some more stuff on it if you wanted
		fmt.Println(da)
	}

}
