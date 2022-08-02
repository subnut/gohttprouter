package gohttprouter

// RFC 3986
// 	pchar       =  unreserved / pct-encoded / sub-delims / ":" / "@"
// 	unreserved  =  ALPHA / DIGIT / "-" / "." / "_" / "~"
// 	sub-delims  =  "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / ";" / "="
const rfc3986_DIGIT = "0123456789"
const rfc3986_ALPHA = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const rfc3986_subdelims = "!" + "$" + "&" + "'" + "(" + ")" + "*" + "+" + "," + ";" + "="
const rfc3986_unreserved = rfc3986_ALPHA + rfc3986_DIGIT + "-" + "." + "_" + "~"
const rfc3986_pchar = rfc3986_unreserved + rfc3986_subdelims + ":" + "@"

