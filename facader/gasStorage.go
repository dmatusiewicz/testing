package facader

import "errors"

type gasStorage struct {
	current int
	max     int
}

func (gs *gasStorage) addResource(v int) error {
	if (v + gs.current) > gs.max {
		return errors.New("Too much resources. Can't add.")
	} else {
		gs.current += v
	}

}
