package commands

func Install(packages []string) error {
	return run("apt-get", "install", packages)
}

func Update(packages []string) error {
	return run("apt-get", "update", nil)
}

func Upgrade(packages []string) error {
	return run("apt-get", "upgrade", nil)
}

func Search(packages []string) error {
	return run("apt-cache", "search", packages)
}