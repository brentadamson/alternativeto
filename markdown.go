package alternativeto

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func GenerateMainREADME() (err error) {
	var lines = []string{
		"# alternativeto",
		"A comprehensive list of alternatives to your favorite software.",
		"",
		"## Index",
	}

	// create the index
	sortedKeys := make([]string, len(allCompanies))
	i := 0
	for k := range allCompanies {
		sortedKeys[i] = k
		i++
	}
	sort.Strings(sortedKeys)
	for _, key := range sortedKeys {
		lines = append(lines, fmt.Sprintf("* [%s](#%s)", key, strings.ToLower(strings.ReplaceAll(key, " ", "-"))))
	}

	// add the companies
	for _, key := range sortedKeys {
		lines = append(lines, "")
		lines = append(lines, fmt.Sprintf("### %s", key))
		lines = append(lines, "Website | Open Source")
		lines = append(lines, "|---|---|")

		companies := allCompanies[key]
		sort.Slice(companies, func(i, j int) bool {
			return companies[i].Name < companies[j].Name
		})
		for _, company := range companies {
			var yesNo = "No"
			if company.OpenSource {
				yesNo = "Yes"
			}

			lines = append(lines, fmt.Sprintf(`| [%s](%s) | %s |`, company.Name, company.Website, yesNo))
		}

		lines = append(lines, "")
		lines = append(lines, "**[⬆ Back to Index](#index)**")
		lines = append(lines, "")
	}

	var contents = []byte(strings.Join(lines, "\n"))
	return os.WriteFile("../../README.md", contents, 0644)
}
