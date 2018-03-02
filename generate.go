package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"text/template"
)

type rootCa struct {
	Name    string
	Comment string
	PemData string
}

func main() {
	dat := genGoSource("mozilla", GetMozillaCerts())
	if err := ioutil.WriteFile(filepath.Join("mozilla", "mozilla.go"), dat, 0644); err != nil {
		panic(err)
	}
}

func genGoSource(pkg string, roots []rootCa) []byte {
	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, map[string]interface{}{
		"pkg":   pkg,
		"roots": roots,
	}); err != nil {
		panic(err)
	}
	dat, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	return dat
}

var tpl = template.Must(template.New("").Parse(`package {{.pkg}}

	import (
		"crypto/x509"
	)
	
	var CaPool *x509.CertPool
	
	func init() {
		CaPool = x509.NewCertPool()
		{{range .roots}}
		// {{.Name}}
		// {{.Comment}}
		CaPool.AppendCertsFromPEM([]byte(` + "`{{.PemData}}`" + `))
	{{ end }}  }
	`))
