package streaming

import (
	"fmt"
	"os"

	slice_utils_module "github.com/pseudoelement/golang-utils/src/utils/slices"
)

func checkValidQuality(value string) error {
	if !slice_utils_module.Contains(QUALITY, value) {
		return fmt.Errorf("%v is invalid quality.", value)
	}

	return nil
}

func checkFileNameExists(quality, fileName string) error {
	pwd, _ := os.Getwd()
	path := pwd + "/streaming/video/" + quality

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, e := range entries {
		if e.Name() == fileName {
			return nil
		}
	}

	return fmt.Errorf("fileName %s doesn't exist in directory %s.", fileName, quality)
}
