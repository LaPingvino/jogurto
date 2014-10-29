package commands

func Remove(packages []string) error {
	return run("apt-get", "remove", packages, true)
}

func Purge(packages []string) error {
	return run("apt-get", "purge", packages, true)
}

func Autoremove(packages []string) error {
	return run("apt-get", "autoremove", nil, true)
}
