package commands

func Install(packages []string) error {
	return run("apt-get", "install", packages, true)
}

func Update(packages []string) error {
	return run("apt-get", "update", nil, true)
}

func Upgrade(packages []string) error {
	return run("apt-get", "upgrade", nil, true)
}

func DistUpgrade(packages []string) error {
	return run("apt-get", "dist-upgrade", nil, true)
}

func Search(packages []string) error {
	return run("apt-cache", "search", packages, false)
}
