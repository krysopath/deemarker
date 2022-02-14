package report

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

type customTime struct {
	time.Time
}

func (c *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	i, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
	if err != nil {
		return err
	}
	*c = customTime{time.Unix(i, 0)}
	return nil
}

func (c *customTime) UnmarshalXMLAttr(attr xml.Attr) error {
	i, err := strconv.ParseInt(strings.TrimSpace(attr.Value), 10, 64)
	if err != nil {
		return err
	}
	*c = customTime{time.Unix(i, 0)}
	return nil
}
