package main

import (
	"bufio"
	"errors"
<<<<<<< HEAD
	"fmt"
=======
	"io/ioutil"
	"log"
>>>>>>> e25142959c73225243eb6b8308c13eb34cca6acb
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

<<<<<<< HEAD
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
=======
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".tf") && !file.IsDir() {
			/*********************************************************************************************************************************************************************/
			f, err := os.Open(file.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()

				// If the line is a comment, skip it
				if strings.HasPrefix(line, "#") {
					continue
				}

				if strings.HasPrefix(line, "resource") {
					fileNameParts := strings.Split(line, " ")
					if len(fileNameParts) == 4 {
						newFileName := fileNameParts[0] + "." + strings.Replace(fileNameParts[1], "\"", "", -1) + "." + strings.Replace(fileNameParts[2], "\"", "", -1) + ".tf"
						/* Skip if the file is named correctly already ***********************************************************************************************************/
						if newFileName == file.Name() {
							PrintSuccess("Named correctly: " + file.Name() + ". Skipping...")
							break
						}
						/* Make sure file with the same name doesn't already exist ***********************************************************************************************/
						if _, err := os.Stat(newFileName); err == nil {
							// path/to/whatever exists
							PrintError("A file with the name \"" + newFileName + "\" exists already. Unable to rename \"" + file.Name() + "\". Skipping...")
							break
						} else if errors.Is(err, os.ErrNotExist) {
							// path/to/whatever does *not* exist
							// Rename file
							if file.Name() != newFileName {
								e := os.Rename(file.Name(), newFileName)
								if e != nil {
									log.Fatal(e)
								}
								PrintWarning("Renamed file from " + file.Name() + " to " + newFileName)
							}
						} else {
							// File may or may not exist. See err for details.
							// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
							log.Fatal(err)
						}
					}
					break // skip rest of the file once first resource is found
>>>>>>> e25142959c73225243eb6b8308c13eb34cca6acb
				}
			}
		}

	}

	readFile.Close()
}
