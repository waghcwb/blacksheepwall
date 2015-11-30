package bsw

import (
	"strings"
)

// NS returns the A record for any NS records for a domain.
func NS(domain, serverAddr string) *Tsk {
	t := newTsk("ns")
	servers, err := LookupNS(domain, serverAddr)
	if err != nil {
		t.SetErr(err)
		return t
	}
	for _, s := range servers {
		ip, err := LookupName(s, serverAddr)
		if err != nil || ip == "" {
			continue
		}
		t.AddResult(ip, strings.TrimRight(s, "."))
	}
	return t
}
