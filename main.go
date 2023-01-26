package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func getTerraformFileName(resourceNameLine string) (string, error) {
	newFileName := ""
	fileNameParts := strings.Split(resourceNameLine, " ")
	if len(fileNameParts) == 4 {
		newFileName = fileNameParts[0] + "." + strings.Replace(fileNameParts[1], "\"", "", -1) + "." + strings.Replace(fileNameParts[2], "\"", "", -1) + ".tf"
	} else {
		err := errors.New("Unable to parse file name from \"" + resourceNameLine + "\". Skipping...")
		return newFileName, err
	}
	return newFileName, nil
}

func stringToFile(terrformBlock string) error {

	// Create a file
	resourceNameLine := strings.Split(terrformBlock, "\n")[0]
	terraformFileName, err := getTerraformFileName(resourceNameLine)
	if err != nil {
		return err
	}
	file, err := os.Create(terraformFileName)
	if err != nil {
		return err
	}
	_, err = file.WriteString(terrformBlock)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func main() {

	readFile, err := os.Open("netappfiles.tf")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	// Iterate through the Terraform file line by line
	// if the line starts with resource, start harvesting
	// if a line includes { only at a single line, this is the end of the block

	// define a stirng called resourceBlock

	resourceBlock := ""

	for fileScanner.Scan() {
		if strings.HasPrefix(fileScanner.Text(), "resource") {
			resourceBlock = fileScanner.Text() + "\n"
			for fileScanner.Scan() {
				resourceBlock = resourceBlock + fileScanner.Text() + "\n"
				if fileScanner.Text() == "}" {
					// End of the blcok
					stringToFile(resourceBlock)
					break
				}
			}
		}

	}

	readFile.Close()
}
