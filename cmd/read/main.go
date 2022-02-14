package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"krysopath.it/deemarker/report"
)

const reportXML = `
<?xml version="1.0"?>
<feedback xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <version>1.0</version>
  <report_metadata>
    <org_name>Outlook.com</org_name>
    <email>dmarcreport@microsoft.com</email>
    <report_id>ded22d897bfd494eb11fcd3818220f50</report_id>
    <date_range>
      <begin>1644451200</begin>
      <end>1644537600</end>
    </date_range>
  </report_metadata>
  <policy_published>
    <domain>verimi.de</domain>
    <adkim>s</adkim>
    <aspf>s</aspf>
    <p>none</p>
    <sp>none</sp>
    <pct>100</pct>
    <fo>1</fo>
  </policy_published>
  <record>
    <row>
      <source_ip>194.25.134.82</source_ip>
      <count>1</count>
      <policy_evaluated>
        <disposition>none</disposition>
        <dkim>pass</dkim>
        <spf>fail</spf>
      </policy_evaluated>
    </row>
    <identifiers>
      <envelope_to>outlook.de</envelope_to>
      <envelope_from>bnc3.mailjet.com</envelope_from>
      <header_from>verimi.de</header_from>
    </identifiers>
    <auth_results>
      <dkim>
        <domain>verimi.de</domain>
        <selector>mailjet</selector>
        <result>pass</result>
      </dkim>
      <spf>
        <domain>bnc3.mailjet.com</domain>
        <scope>mfrom</scope>
        <result>fail</result>
      </spf>
    </auth_results>
  </record>
</feedback>
`

func init() {
	flag.Parse()
}

func main() {
	r, err := report.ReadReports(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	jsonBytes, err := json.Marshal(r)
	fmt.Fprintf(os.Stdout, "%s\n", string(jsonBytes))
}
