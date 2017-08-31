//go:generate goversioninfo -icon=icon.ico
package main

import (
	"fmt"
	"github.com/HearthProject/OneClientPreLauncher/utils"
	"github.com/shibukawa/configdir"
	"path/filepath"
	"os/exec"
	"log"
	"os"
)


func main(){
	fmt.Println("Starting One Client!")

	fmt.Println("Checking for new version!")
	versionJson, err := utils.GetString("http://hearthproject.uk/files/versions.json")
	if err != nil {
		println(err)
	}
	downloadUrl := utils.GetStringValue(versionJson, "downloadUrl")
	latestVersion := utils.GetStringValue(versionJson, "latestVersion")

	saveDir := configdir.New("hearthproject", "oneclientprelauncher").QueryCacheFolder().Path
	if !utils.FileExists(saveDir){
		utils.MakeDir(saveDir)
	}

	jarFile := saveDir + string(filepath.Separator) + latestVersion + ".jar"

	if !utils.FileExists(jarFile){
		fmt.Println("Downloading launcher update to " + jarFile)
		err := utils.DownloadFile(jarFile, downloadUrl)
		if err != nil {
 			print(err)
		}
	}

	fmt.Println("Starting...")
	exec := exec.Command("java", "-jar", jarFile, "-native_launch")
	exec.Dir = saveDir
	er := exec.Start()
	if er != nil {
		log.Fatal(er)
		os.Exit(1)
	}
}

