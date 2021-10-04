# Test Number 3
Please refactor the code below to make it more concise, efficient and readable with good logic flow.
```text
Answer : 
func findFirstStringInBracket(str string) string {
	if str == "" {
		return ""
	}

	indexFirstBracketFound := strings.Index(str, "(")
	indexClosingBracketFound := strings.Index(str, ")")

	if indexFirstBracketFound < 0 || indexClosingBracketFound < 0 {
		return ""
	}
	
	return str[indexFirstBracketFound + 1:indexClosingBracketFound]
}

func findFirstStringInBracket(str string) string {
	if str == "" {
		return ""
	}

	if strings.Index(str, "(") < 0 || strings.Index(str, ")") < 0 {
		return ""
	}

	return str[strings.Index(str, "(")+1 : strings.Index(str, ")")]
}
```