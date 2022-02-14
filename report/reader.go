package report

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var FeedbackRegexp *regexp.Regexp = regexp.MustCompile("(?is:<feedback.*</feedback>)")

func readReport(path string, reports map[string][]Feedback) error {
	var file io.Reader
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	validBytes := FeedbackRegexp.Find(fileBytes)

	fmt.Fprintf(os.Stderr, "'%s' is dmarc report\n", path)

	var q Query
	err = xml.Unmarshal(validBytes, &q.Feedback)
	if err != nil {
		return err
	}
	domain := q.Feedback.PolicyPublished.Domain
	domainReports, ok := reports[domain]
	if !ok {
		domainReports = make([]Feedback, 0)
		reports[domain] = append(domainReports, q.Feedback)
	}
	reports[domain] = append(domainReports, q.Feedback)
	return nil

}

func getVisitFunc(reports map[string][]Feedback) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".xml") {
			return readReport(path, reports)
		}
		return nil
	}
}

func ReadReports(reportPath string) (map[string][]Feedback, error) {
	reports := make(map[string][]Feedback)

	err := filepath.Walk(reportPath, getVisitFunc(reports))
	if err != nil {
		return nil, err
	}
	return reports, nil
}
