package quic_go_utils

import (
	"fmt"
	"regexp"
)

func AppendPortIfNotSpecified(address string, port int) string {
	hasPort, _ := regexp.MatchString("^.*:\\d+$", address)
	if hasPort {
		return address
	}
	return fmt.Sprintf("%s:%d", address, port)
}
