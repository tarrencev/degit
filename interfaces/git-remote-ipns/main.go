package main

import (
	"log"
	"os"

	"github.com/Persper/degit/core"
)

func main() {

	if len(os.Args) < 3 {
		core.Err_print("Usage: git-remote-ipns remote-name url\n")
		return
	}

	// os.args[0] [1] [2]
	// git cloone: $GOROOT/bin/git-remote-ipns ipns::
	// git push: $GOROOT/bin/git-remote-ipns ipns::  ipns-value
	// git pull: $GOROOT/bin/git-remote-ipns ipns::  ipns-value

	/* Transform the alias to the ipns hash */
	if len(os.Args[2]) > 0 && os.Args[2][0:2] != "Qm" {
		hash, err := core.Read_IPNS_alias(os.Args[2])
		if err != nil {
			log.Fatal(err)
			return
		}
		os.Args[2] = hash
	}

	/*
	 * For example: git clone ipns::QmS5mHovjz7soFc7joLu2smafRdNg2QDvBGu4s7EKm29Qv
	 * QmS5mHovjz7soFc7joLu2smafRdNg2QDvBGu4s7EKm29Qv: the ipns key value
	 * IPNS_Key -> IPFS_Key -> Git_Commit_Hash
	 */
	if len(os.Args[2]) > 0 && os.Args[2][0:2] == "Qm" {
		repo_ipfs_hash, _ := core.Transform_ipns_to_ipfs(os.Args[2])
		os.Args[2] = core.Transform_ipfs_to_git(repo_ipfs_hash)
	}

	if err := core.Main(true); err != nil {
		log.Fatal(err)
	}
}
