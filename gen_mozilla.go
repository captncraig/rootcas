package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

func GetMozillaCerts() []rootCa {
	// mozilla publishes a nice csv file with complete PEM blocks. Thanks!
	resp, err := http.Get("https://ccadb-public.secure.force.com/mozilla/IncludedCACertificateReportPEMCSV")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)
	dat, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	roots := []rootCa{}
	for _, row := range dat[1:] {
		roots = append(roots, rootCa{
			Name:    row[3],
			Comment: fmt.Sprintf("%s Valid %s to %s, SHA256:%s", row[9], row[7], row[8], row[5]),
			PemData: strings.Trim(row[28], "'"),
		})
	}
	return roots
}
