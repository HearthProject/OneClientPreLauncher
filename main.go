//go:generate goversioninfo -icon=icon.ico

// go build -ldflags="-H windowsgui"

package main

import (
	"fmt"
	"github.com/HearthProject/OneClientPreLauncher/utils"
	"github.com/shibukawa/configdir"
	"path/filepath"
	"os/exec"
	"log"
	"os"
	"runtime"
	"net/http"
	"github.com/inconshreveable/go-update"
)


func main(){
	fmt.Println("Starting One Client!")

	platform := runtime.GOOS + "-" + runtime.GOARCH
	fmt.Println("Pre-Launcher version " + utils.Version + " running on " + platform)

	fmt.Println("Checking for pre-launcher update")
	checkForUpdate(utils.Version, platform)

	fmt.Println("Checking for oneclient update")
	versionJson, err := utils.GetString("http://fdn.redstone.tech/theoneclient/oneclient/versions.json")
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
	javaExec := "java"
	if runtime.GOOS == "windows" {
		javaExec = "javaw"
	}

	exec := exec.Command(javaExec, "-jar", jarFile, "-native_launch")
	exec.Dir = saveDir
	er := exec.Start()
	if er != nil {
		log.Fatal(er)
		os.Exit(1)
	}
}


func checkForUpdate(currentVersion string, platform string){
	json, err := utils.GetString("http://fdn.redstone.tech/theoneclient/oneclient/launcher/prelauncher_versions.json")
	if err != nil {
		println(err)
	}
	latestVersion := utils.GetStringValue(json, "version")
	if(currentVersion != latestVersion){
		fmt.Println("Downloading pre-launcher update " + latestVersion)
		versionJson, err := utils.GetString(utils.GetStringValue(json, "versionJsonURL"))
		if err != nil {
			println(err)
		}
		url, err := utils.GetQuery(versionJson).String("files", platform, "url")
		if err != nil {
			println(err)
		}
		hash, _ := utils.GetQuery(versionJson).String("files", platform, "sha256")
		fmt.Println("Updating to " + url)
		doUpdate(url, hash)
	}
}


func doUpdate(url string, hash string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{
		Checksum: []byte(hash),
	})
	if err != nil {
		println(err)
	}
	return err
}