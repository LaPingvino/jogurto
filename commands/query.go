package commands

func LocalSearch(packages []string) error {
	packages = append(packages, `""`) //Trick to prevent barking on empty input. Will give a quoted empty string to grep so it shows everything
	return run("bash", "-c", []string{"dpkg --get-selections | grep " + packages[0]})
}
