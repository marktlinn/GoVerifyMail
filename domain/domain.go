package domain

import (
	"fmt"
	"net"
	"strings"
)

type Domain struct {
	domain      string
	hasSPF      bool
	hasMX       bool
	hasDMARC    bool
	spfRecord   string
	dmarcRecord string
}

func (d *Domain) setMX() error {
	mxRecords, err := net.LookupMX(d.domain)
	if err != nil {
		return fmt.Errorf("failed to lookup MX records: %w", err)
	}
	if len(mxRecords) > 0 {
		d.hasMX = true
	}
	return nil
}

func (d *Domain) setSPF() error {
	spfRecords, err := net.LookupTXT(d.domain)
	if err != nil {
		return fmt.Errorf("failed to lookup SPF records: %s\n", err)
	}

	for _, record := range spfRecords {
		if strings.HasPrefix(record, "v=spf1") {
			d.hasSPF = true
			d.spfRecord = record
			break
		}
	}
	return nil
}

func (d *Domain) setDMARC() error {
	dmarcRecord, err := net.LookupTXT("_dmarc." + d.domain)
	if err != nil {
		return fmt.Errorf("failed to lookup DMARC records: %s\n", err)
	}

	for _, record := range dmarcRecord {
		if strings.HasPrefix(record, "v=DMARC1") {
			d.hasDMARC = true
			d.dmarcRecord = record
			break
		}
	}
	return nil
}

func CheckDomain(domainName string) error {
	domain := Domain{domain: domainName}

	if err := domain.setMX(); err != nil {
		return fmt.Errorf("failed to set MX: %w", err)
	}

	if err := domain.setSPF(); err != nil {
		return fmt.Errorf("failed to set SPF: %w", err)
	}

	if err := domain.setDMARC(); err != nil {
		return fmt.Errorf("failed to set DMARC: %w", err)
	}
	fmt.Printf("Searched %s\n", domainName)
	fmt.Println("Found:")
	fmt.Printf("hasSPF:    %t\n", domain.hasSPF)
	fmt.Printf("hasMX:     %t\n", domain.hasMX)
	fmt.Printf("hasDMARC:  %t\n", domain.hasDMARC)
	fmt.Printf("spfRecord: %s\n", domain.spfRecord)
	fmt.Printf("dmarcRecord: %s\n", domain.dmarcRecord)
	return nil
}
