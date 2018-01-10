package util

// CheckTrailingSlash checks the presence of trailing slash in
// case the argument is a directory, and appends on if absent.
func CheckTrailingSlash(dir string) string {
	if dir[len(dir) - 1] != '/' {
		dir += "/"
	}
	return dir
}
