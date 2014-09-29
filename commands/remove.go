package commands

func Remove(packages []string) error {
	return run("apt-get", "remove", packages)
}

func Purge(packages []string) error {
	return run("apt-get", "purge", packages)
}
