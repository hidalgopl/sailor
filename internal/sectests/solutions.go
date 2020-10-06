package sectests

import "fmt"

const (
	SEC0001 = "SEC0001: X-Content-Type-Options: no-sniff\n" +
		"The server should send an X-Content-Type-Options: nosniff \n" +
		"to make sure the browser does not try to detect a different Content-Type \n" +
		"than what is actually sent (as this can lead to XSS)\n" +
		"Learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options\n"
	SEC0002 = "SEC0002: X-Frame-Options: deny or sameorigin\n" +
		"The server should send the X-Frame-Options security header with deny or sameorigin value,\n" +
		"to protect against drag'n drop clickjacking attacks in older browsers.\n" +
		"Sites can use this to avoid clickjacking attacks, by ensuring that their content \n" +
		"is not embedded into other sites.\n" +
		"Learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options\n"
	SEC0003 = "SEC0003: X-XSS-Protection: 1; mode=block\n" +
		"This header enables the Cross-site scripting (XSS) filter in your browser. 1; mode=block Filter enabled.\n" +
		"Rather than sanitize the page, when a XSS attack is detected,\n" +
		"the browser will prevent rendering of the page.\n" +
		"Learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection"
	SEC0004 = "SEC0004: Content-Security-Policy: policy\n" +
		"Content Security Policy (CSP) is an added layer of security\n" +
		"that helps to detect and mitigate certain types of attacks, including Cross Site Scripting (XSS)\n" +
		"and data injection attacks. These attacks are used for everything\n" +
		"from data theft to site defacement to distribution of malware.\n" +
		"Learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP\n"
	SEC0005 = "SEC0005: Fingerprint headers\n" +
		"Looks like your API returns some fingerprint headers, such as \n" +
		"\"X-Powered-By\", \"X-Generator\", \"Server\", \"X-ASPNet-Version\", \"X-ASPNETMVC-version\".\n" +
		"As each web framework has it's own Common Vulnerabilities and Exposures\n" +
		"(e.g. https://www.cvedetails.com/vulnerability-list/vendor_id-12043/product_id-22568/Rubyonrails-Ruby-On-Rails.html )\n" +
		"you shouldn't expose such information. Remove those headers and this tests would be green.\n" +
		"Learn more: https://github.com/OWASP/wstg/blob/master/document/4-Web_Application_Security_Testing/01-Information_Gathering/02-Fingerprint_Web_Server.md\n"
	SEC0006 = "SEC0006: Access-Control-Allow-Origin\n" +
		"This test is failing if your API returns Access-Control-Allow-Origin: \"*\" which tells the browser\n" +
		"to allow code from any origin to access a resource. In most cases (if your API isn't public facing one)\n" +
		"it's way too broad clause.\n" +
		"Learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin\n"
	SEC0007 = "SEC0007: Strict-Transport-Security: max-age=(age in seconds); (other options)\n" +
		"This header lets a web site tell browsers that it should only be accessed using HTTPS, instead of using HTTP.\n" +
		"Learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security\n"
	SEC0008 = "SEC0008: Set-Cookie header should contain Secure and HttpOnly options\n" +
		"Setting Secure option ensures cookie is only sent to the server when a request is made with the https: scheme.\n" +
		"HttpOnly option forbids JavaScript from accessing the cookie.\n" +
		"This mitigates attacks against cross-site scripting (XSS).\n" +
		"Learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie\n"
	SEC0009 = "SEC0009: Set Cache-Control or Expires header\n" +
		"The Expires header contains the date/time after which the response is considered stale.\n" +
		"If there is a Cache-Control header with the max-age or s-maxage directive in the response,\n" +
		"the Expires header is ignored.\n" +
		"Learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expires\n"
)

var (
	SEC_TEST_KEYS = []string{
		"SEC0001",
		"SEC0002",
		"SEC0003",
		"SEC0004",
		"SEC0005",
		"SEC0006",
		"SEC0007",
		"SEC0008",
		"SEC0009",
	}
	SEC_TEST_SOLUTIONS = map[string]string{
		"SEC0001": SEC0001,
		"SEC0002": SEC0002,
		"SEC0003": SEC0003,
		"SEC0004": SEC0004,
		"SEC0005": SEC0005,
		"SEC0006": SEC0006,
		"SEC0007": SEC0007,
		"SEC0008": SEC0008,
		"SEC0009": SEC0009,
	}
)

func PrintExplanation(args []string) {
	for _, secTestKey := range args {
		fmt.Println("---------------------------------------------------------------------------------------------")
		fmt.Println(SEC_TEST_SOLUTIONS[secTestKey])
		fmt.Println("---------------------------------------------------------------------------------------------")
	}
}

func PrintSummary(failed int, total int)  {
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("|                    TOTAL: %v | FAILED: %v | PASSED: %v                                       |\n", total, failed, total-failed)
	fmt.Println("---------------------------------------------------------------------------------------------")
}