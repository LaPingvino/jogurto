package commands

func InstallLocal(packages []string) error {
	return run("dpkg", "-i", packages, true)
}

func InstallMissing(packages []string) error {
	return run("apt-get", "-f", []string{"install"}, true)
}
