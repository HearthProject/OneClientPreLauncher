//go:generate goversioninfo -icon=icon.ico
package main

import (
	"fmt"
	"github.com/HearthProject/OneClientPreLauncher/utils"
	"github.com/shibukawa/configdir"
	"path/filepath"
	"github.com/go-toast/toast"
	"os/exec"
	"log"
	"os"
)


func main(){
	fmt.Println("Starting the oneclient!")


	versionJson, err := utils.GetString("http://hearthproject.uk/files/versions.json")
	if err != nil {
		println(err)
	}
	downloadUrl := utils.GetStringValue(versionJson, "downloadUrl")
	latestVersion := utils.GetStringValue(versionJson, "latestVersion")
	fmt.Println(downloadUrl)
	fmt.Println(latestVersion)

	saveDir := configdir.New("hearthproject", "oneclientprelauncher").QueryCacheFolder().Path
	if !utils.FileExists(saveDir){
		utils.MakeDir(saveDir)
	}

	jarFile := saveDir + string(filepath.Separator) + latestVersion + ".jar"

	if !utils.FileExists(jarFile){

		notification := toast.Notification{
			AppID: "One Client",
			Title: "Downloading One Client update!",
			Message: "OneClient is now being updated to version " + latestVersion,
		}
		notification.Push()

		fmt.Println("Downloading launcher update to " + jarFile)
		err := utils.DownloadFile(jarFile, downloadUrl)
		if err != nil {
 			print(err)
		}
	}

	if !IsJavaInstalled() {
		//TODO show an error
	}


	exec := exec.Command("java", "-jar", jarFile, "-native_launch")
	exec.Dir = saveDir
	er := exec.Run()
	if er != nil {
		log.Fatal(er)
		os.Exit(1)
	}
	fmt.Println("Goodbye")
}

func IsJavaInstalled() bool {
	_, err := exec.LookPath("java")
	return err != nil
}
