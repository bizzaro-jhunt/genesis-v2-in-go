package kit

import (
	"os"
)

func DevKit() (Kit, error) {
	f, err := os.Open(DevDirectory+"/"+KitMetadataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return Kit{}, NotFoundError{IsDev: true}
		}
		return Kit{}, err
	}
	defer f.Close()

	k := Kit{IsDev: true}
	err = k.load(f)
	if err != nil {
		return k, err
	}

	return k, nil
}
