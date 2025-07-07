package utils

import (
	"github.com/niklasfasching/go-org/org"
	"regexp"
	"strings"
)

var startPropertiesBlock *regexp.Regexp = regexp.MustCompile(`\s*:PROPERTIES:`)
var titleSection *regexp.Regexp = regexp.MustCompile(`\s*:TITLE:\s.*`)
var descSection *regexp.Regexp = regexp.MustCompile(`\s*:DESCRIPTION:\s.*`)
var endPropertiesBlock *regexp.Regexp = regexp.MustCompile(`\s*:END:`)

func ParseProperties(nodes []org.Node) map[string]string {
	properties := map[string]string{}
	for _, node := range nodes {
		if m := startPropertiesBlock.FindStringSubmatch(node.String()); m != nil {
			subnodes := strings.Split(node.String(), "\n")[1:]
			for _, subnode := range subnodes {
				if m := titleSection.FindStringSubmatch(subnode); m != nil {
					properties["TITLE"] = strings.TrimLeft(subnode, ":TITLE: ")
				}
				if m := descSection.FindStringSubmatch(subnode); m != nil {
					properties["DESCRIPTION"] = strings.TrimLeft(subnode, ":DESCRIPTION: ")
				}
				if m := endPropertiesBlock.FindStringSubmatch(subnode); m != nil {
					break
				}
			}
		}
	}
	return properties
}
