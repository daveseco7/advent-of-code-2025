package util

import (
	"log"
	"regexp"
)

// MustGetFromCaptureGroups Extracts the capture groups of the provided string line.
// This may not be the best approach performance wise, but it is not important for this situation.
// NOTE: The provided string must contain all the capture groups, or else a panic will be triggered.
func MustGetFromCaptureGroups(regEx, line string) map[string]string {
	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(line)
	names := compRegEx.SubexpNames()

	captureGroups := make(map[string]string)
	for i, name := range names {
		if i > 0 && i <= len(match) {
			captureGroups[name] = match[i]
		}
	}

	if len(match) != len(names) {
		log.Fatalf("the provided string does not contain all regex's capture groups")
	}

	return captureGroups
}
