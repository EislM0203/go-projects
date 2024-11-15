package conversion

import "strconv"

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for stringIndex, stringValue := range(strings) {
		float, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			return nil, err
		}
		floats[stringIndex] = float
	}
	return floats, nil
}