// Written by Rob Fuller / mubix
// Source: https://github.com/vmware/go-nfs-client
// Using: https://github.com/go-nfs/nfsv3

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-nfs/nfsv3/nfs"
	"github.com/go-nfs/nfsv3/nfs/rpc"
	"github.com/go-nfs/nfsv3/nfs/util"
	"github.com/jedib0t/go-pretty/v6/table"
)

func ls(v *nfs.Target, path string) ([]*nfs.EntryPlus, error) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	dirs, err := v.ReadDirPlus(path)
	if err != nil {
		return nil, fmt.Errorf("readdir error: %s", err.Error())
	}

	util.Infof("dirs:")
	t.AppendHeader(table.Row{"Filename", "UID", "GID", "Mode", "Size"})
	for _, dir := range dirs {
		t.AppendRow(table.Row{dir.FileName, dir.Attr.Attr.UID, dir.Attr.Attr.GID, dir.Attr.Attr.Mode, dir.Attr.Attr.Size()})
	}
	t.Render()
	return dirs, nil
}

func uploadFile(v *nfs.Target, path string, dest string) {
	f, err := os.Open(path)
	util.Infof("Opening %s", path)
	if err != nil {
		util.Errorf("Could not open local file: %s", err.Error())
		return
	}
	fileinfo, err := f.Stat()
	if err != nil {
		util.Errorf("read fail: %s", err.Error())
		return
	}
	filesize := fileinfo.Size()
	wr, err := v.OpenFile(dest, 0777)
	if err != nil {
		util.Errorf("write fail: %s", err.Error())
		return
	}

	// calculate the sha
	h := sha256.New()
	t := io.TeeReader(f, h)

	// Copy filesize
	n, err := io.CopyN(wr, t, int64(filesize))
	if err != nil {
		util.Errorf("error copying: n=%d, %s", n, err.Error())
		return
	}
	expectedSum := h.Sum(nil)

	if err = wr.Close(); err != nil {
		util.Errorf("error committing: %s", err.Error())
		return
	}

	//
	// get the file we wrote and calc the sum
	rdr, err := v.Open(dest)
	if err != nil {
		util.Errorf("read error: %v", err)
		return
	}

	h = sha256.New()
	t = io.TeeReader(rdr, h)

	_, err = ioutil.ReadAll(t)
	if err != nil {
		util.Errorf("readall error: %v", err)
		return
	}
	actualSum := h.Sum(nil)

	if bytes.Compare(actualSum, expectedSum) != 0 {
		log.Fatalf("sums didn't match. actual=%x expected=%s", actualSum, expectedSum) //  Got=0%x expected=0%x", string(buf), testdata)
	}

	log.Printf("Sums match %x %x", actualSum, expectedSum)
}

func downloadFile(v *nfs.Target, path string) {
	wr, err := v.Open(path)
	if err != nil {
		util.Errorf("read fail: %s", err.Error())
		return
	}
	pathsplit := strings.Split(path, "/")
	filename := pathsplit[len(pathsplit)-1]

	fileInfo, _, err := wr.Lookup(path)
	filesize := fileInfo.Size()
	util.Infof("FIle size: %d", filesize)

	if err != nil {
		util.Errorf("read fail: %s", err.Error())
		return
	}
	util.Infof("Name: %s", filename)
	f, err := os.OpenFile(filename, os.O_CREATE, 0777)

	if err != nil {
		util.Errorf("read fail: %s", err.Error())
		return
	}

	// calculate the sha
	h := sha256.New()
	t := io.TeeReader(wr, h)

	// Copy filesize
	n, err := io.CopyN(f, t, int64(filesize))
	if err != nil {
		util.Errorf("error copying: n=%d, %s", n, err.Error())
		return
	}
	expectedSum := h.Sum(nil)

	if err = wr.Close(); err != nil {
		util.Errorf("error committing: %s", err.Error())
		return
	}
	f.Close()

	//
	// get the file we wrote and calc the sum
	rdr, err := os.Open(filename)
	if err != nil {
		util.Errorf("read error: %v", err)
		return
	}

	h = sha256.New()
	t = io.TeeReader(rdr, h)

	_, err = ioutil.ReadAll(t)
	if err != nil {
		util.Errorf("readall error: %v", err)
		return
	}
	actualSum := h.Sum(nil)

	if bytes.Compare(actualSum, expectedSum) != 0 {
		log.Fatalf("sums didn't match. actual=%x expected=%s", actualSum, expectedSum) //  Got=0%x expected=0%x", string(buf), testdata)
	}

	log.Printf("Sums match %x %x", actualSum, expectedSum)
	wr.Close()
	f.Close()
}

func removeFile(v *nfs.Target, path string) {
	err := v.Remove(path)
	if err != nil {
		log.Fatalf("rm of %s err: %s", path, err.Error())
	}
}

func makeDirectory(v *nfs.Target, path string) {
	if _, err := v.Mkdir(path, 0775); err != nil {
		log.Fatalf("mkdir error: %v", err)
	}
}

func removeDirectory(v *nfs.Target, path string) {
	if err := v.RmDir(path); err != nil {
		log.Fatalf("mkdir error: %v", err)
	}
}

func main() {
	util.DefaultLogger.SetDebug(true)
	util.Infof(strconv.Itoa(len(os.Args)))
	if len(os.Args) <= 3 {
		util.Infof("%s <host>:<target path> <access level root:0:0> <command ls/up/down/rm/mkdir/rmdir> <path if required> <dest if upload>", os.Args[0])
		os.Exit(-1)
	}
	fmt.Println(os.Args)
	b := strings.Split(os.Args[1], ":")
	c := strings.Split(os.Args[2], ":")

	host := b[0]
	target := b[1]
	user := c[0]
	uidstring := c[1]
	gidstring := c[2]
	cmd := os.Args[3]
	path := "."
	dest := ""
	if len(os.Args) > 4 {
		path = os.Args[4]
	}
	if len(os.Args) > 5 {
		dest = os.Args[5]
	}

	// convert uid to int
	uid64, err := strconv.Atoi(uidstring)
	uid := uint32(uid64)
	if err != nil {
		log.Fatalf("UID needs to be an integer - %v", err)
	}

	// convert uid to int
	gid64, err := strconv.ParseUint(gidstring, 10, 32)
	gid := uint32(gid64)
	if err != nil {
		log.Fatalf("GID needs to be an integer - %v", err)
	}

	util.Infof("host=%s target=%s command=%s\n", host, target, cmd)

	//connect
	mount, err := nfs.DialMount(host, true)
	if err != nil {
		log.Fatalf("unable to dial MOUNT service: %v", err)
	}
	defer mount.Close()

	//auth
	auth := rpc.NewAuthUnix(user, uid, gid)

	v, err := mount.Mount(target, auth.Auth())
	if err != nil {
		log.Fatalf("unable to mount volume: %v", err)
	}
	defer v.Close()

	switch cmd {
	case "ls":
		ls(v, path)
	case "down":
		downloadFile(v, path)
	case "up":
		uploadFile(v, path, dest)
	case "rm":
		removeFile(v, path)
	case "mkdir":
		makeDirectory(v, path)
	case "rmdir":
		removeDirectory(v, path)
	default:
		util.Infof("No command given, just running an directory listing")
	}

	if err = mount.Unmount(); err != nil {
		log.Fatalf("unable to umount target: %v", err)
	}
	mount.Close()
	util.Infof("Completed tests")
}
