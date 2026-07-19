package profile

import "os"

func SaveConfig(
	path string,
	config string,
) error {

	return os.WriteFile(
		path,
		[]byte(config),
		0600,
	)
}
