package domain

import "net"

func CheckDomain(domain string) {
	var hasSPF, hasMX, hasDMARC bool
	var spfRecord, dmarRecord string

	net.LookupMX(domain)
}
