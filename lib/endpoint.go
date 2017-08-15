package proxy

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

// GetCalcEndpoint function  the calc service endpoint from DNS
func GetCalcEndpoint() (string, error) {
	var addrs []*net.SRV
	var err error
	if _, addrs, err = net.LookupSRV("", "", "test-api.servicediscovery.internal"); err != nil {
		return "", err
	}
	for _, addr := range addrs {
		return strings.TrimRight(addr.Target, ".") + ":" + strconv.Itoa(int(addr.Port)), nil
	}
	return "", errors.New("No record found")
}
