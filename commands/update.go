package commands

func InstallLocal(packages []string) error {
	return run("dpkg", "-i", packages)
}

func InstallMissing(packages []string) error {
	return run("apt-get", "-f", []string{"install"})
}
