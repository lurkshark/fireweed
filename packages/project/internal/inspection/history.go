package inspection

import "errors"

func LoadAt(start, revision string) (Data, error) {
	return Data{}, errors.New("baseline snapshot is not implemented")
}
