package strand

import "strings"

func ToRNA(dna string) string {
	var sb strings.Builder

	for _, nuc := range dna {
		switch nuc {
		case 'A':
			sb.WriteRune('U')
		case 'G':
			sb.WriteRune('C')
		case 'C':
			sb.WriteRune('G')
		case 'T':
			sb.WriteRune('A')
		}
	}

	return sb.String()
}
