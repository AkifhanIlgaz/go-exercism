package protein

import "errors"

func FromRNA(rna string) ([]string, error) {
	var proteins []string

	for i := 0; i < len(rna)-2; i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err != nil {
			if err == ErrStop {
				break
			} else {
				return nil, err
			}
		}

		proteins = append(proteins, protein)
	}

	return proteins, nil
}

var ErrInvalidBase = errors.New("invalid base")
var ErrStop = errors.New("stop")

func FromCodon(codon string) (string, error) {
	var protein string

	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}

	if protein == "" {
		return "", ErrInvalidBase
	}

	return protein, nil
}
