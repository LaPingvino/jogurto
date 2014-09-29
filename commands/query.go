package commands

func LocalSearch(packages []string) error {
	return run("apt-cache", "pkgnames", nil)
}
