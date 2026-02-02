package output

import (
	"fmt"
	"os"

	g "github.com/s0md3v/smap/internal/global"
)

var openedPairFile *os.File

func StartPair() {
	if g.PairFilename != "-" {
		openedPairFile = OpenFile(g.PairFilename)
	}
}

func ContinuePair(result g.Output) {
	thisString := ""
	// Use the user-provided hostname (domain) if available, otherwise use IP
	host := result.IP
	if result.UHostname != "" {
		host = result.UHostname
	}
	for _, port := range result.Ports {
		thisString += fmt.Sprintf("%s:%d\n", host, port.Port)
	}
	Write(thisString, g.PairFilename, openedPairFile)
}

func EndPair() {
}
