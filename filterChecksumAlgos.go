package main

import "crypto"
import "strings"

// Checksum algos
import (
    _ "crypto/md5"
    _ "crypto/sha1"
    _ "crypto/sha256"
    _ "crypto/sha512"
)

// Man, why don't people allow their table to be exported...
var checksumLookupTable = map[string]crypto.Hash{
    "md5":       crypto.MD5,
    "sha1":      crypto.SHA1,
    "sha256":    crypto.SHA256,
    "sha512":    crypto.SHA512,
}

var checksumAlgos = map[string]crypto.Hash{}

func filterChecksumAlgos() {
    i := strings.Split(checksums, ",")
    var j = map[string]crypto.Hash{}
    for _, checksum := range i {
        if checksumLookupTable[checksum].Available() == false {
            Error.Fatalln("Unsupported checksum algorithm: " + checksum)
        }
        j[checksum] = checksumLookupTable[checksum]
    }
    checksumLookupTable = j
}
