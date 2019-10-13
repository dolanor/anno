package anno

import (
	"bytes"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

// Punctuation is a string of common punctuation and quotation characters.
var Punctuation = `.?!'",;`

// TrimPunctuation trims off Punctuation characters.
func TrimPunctuation(s []byte) []byte {
	return bytes.Trim(s, Punctuation)
}

var tlds = [][]byte{
	[]byte(".com"),
	[]byte(".dev"),
	[]byte(".net"),
	[]byte(".org"),
	[]byte(".edu"),
	[]byte(".co"),
	[]byte(".io"),
	[]byte(".uk"),
	[]byte(".us"),
	[]byte(".ca"),
	[]byte(".de"),
	[]byte(".jp"),
	[]byte(".fr"),
	[]byte(".au"),
	[]byte(".ru"),
	[]byte(".ch"),
	[]byte(".it"),
	[]byte(".nl"),
	[]byte(".se"),
	[]byte(".no"),
	[]byte(".es"),
	[]byte(".mil"),
}

// Emails finds email addresses.
var Emails = FieldFunc("email", func(s []byte) (bool, []byte) {
	trimmedS := TrimPunctuation(s)
	if !bytes.Contains(s, []byte("@")) { // not email address
		return false, s
	}
	for _, tld := range tlds {
		if bytes.HasSuffix(s, tld) {
			return true, trimmedS
		}
	}
	return false, s
})

// URL finds web addresses.
var URLs = FieldFunc("url", func(s []byte) (bool, []byte) {
	trimmedS := TrimPunctuation(s)
	if bytes.Contains(s, []byte("@")) { // email address
		return false, s
	}
	if bytes.HasPrefix(trimmedS, []byte("http")) || bytes.HasPrefix(trimmedS, []byte("www")) {
		return true, trimmedS
	}
	for _, tld := range tlds {
		if bytes.HasSuffix(trimmedS, tld) {
			return true, trimmedS
		}
	}
	return false, s
})

// Mentions finds @twitter style mentions.
var Mentions = FieldFunc("mention", func(s []byte) (bool, []byte) {
	trimmedS := TrimPunctuation(s)
	return bytes.HasPrefix(trimmedS, []byte("@")), trimmedS
})

// Hashtags finds #hashtags.
var Hashtags = FieldFunc("hashtag", func(s []byte) (bool, []byte) {
	trimmedS := TrimPunctuation(s)
	return bytes.HasPrefix(trimmedS, []byte("#")), trimmedS
})

var Dates = FieldFunc("date", func(s []byte) (bool, []byte) {
	trimmedS := TrimPunctuation(s)
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	r, err := w.Parse(string(trimmedS), time.Now())
	if err != nil || r == nil {
		return false, trimmedS
	}

	return true, trimmedS
})
