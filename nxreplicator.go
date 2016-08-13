package main

import (
	"fmt"
	"os"
	"os/user"
	"path"

	bp "github.com/nexustix/boilerplate"
	nrc "github.com/nxReplicator/nxReplicatorCommon"
)

//nxreplicator install amazingBulk /home/sam/amazingBulk

func main() {
	version := "V.0-1-0"
	fmt.Printf("<-> NxReplicator Version: %s\n", version)

	args := os.Args

	usr, err := user.Current()
	bp.FailError(err)
	workingDir := usr.HomeDir
	//atomDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", "atoms")
	//moleculeDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", "molecules")
	bulkDir := nrc.InitWorkFolder(workingDir, ".nxreplicator", "bulks")

	//atomManager := nrc.AtomManager{WorkingDir: atomDir}
	//molecule := nrc.Molecule{}

	action := bp.StringAtIndex(1, args)
	bulkID := bp.StringAtIndex(2, args)
	destination := bp.StringAtIndex(3, args)
	//atomID := bp.StringAtIndex(4, args)

	fmt.Printf("%s | %s | %s\n", action, bulkID, destination)

	switch action {
	case "install":
		installBulk(bulkID, bulkDir, destination)
	}
}

func installBulk(bulkID, bulkDir, destination string) {
	bulkPath := path.Join(bulkDir, bulkID+".nxrb")
	tmpBulk := nrc.Bulk{}
	tmpBulk.LoadFromFile(bulkPath)

	for k, v := range tmpBulk.BulkItems {
		//fmt.Printf("(%v) %s %s\n", k, v.Download.Filename, v.Download.URL)
		fmt.Printf("(%v of %v) %s \n", k+1, len(tmpBulk.BulkItems), v.Download.Filename)
		err := os.MkdirAll(path.Join(destination, v.RelativePath), 0777)
		bp.FailError(err)
		bp.DownloadRemoteFile(path.Join(destination, v.RelativePath, v.Download.Filename), v.Download.URL)
	}

	//bp.DownloadRemoteFile("./test", "https://upload.wikimedia.org/wikipedia/commons/5/50/SSTV_reception.png")
}
