package main

import "bufio"
import "encoding/hex"
import "github.com/vbauerster/mpb"
import "hash"
import "io"
import "os"

func checksumCount(path string) int {
    var count = 0

    for checksumAlgo := range checksumLookupTable {
        checksumPath := xattrRoot + checksumAlgo
        checksumValue, _ := GetxattrHex(path, checksumPath)
        if len(checksumValue) > 0 {
            count += 1
        }
    }
    return count
}

func missingChecksums(path string) bool {
    return checksumCount(path) != len(checksumLookupTable)
}

func hasAllChecksums(path string) bool {
    return checksumCount(path) == len(checksumLookupTable)
}

func checksumPath(path string) (map[string]string, error) {
    var hashers = map[string]hash.Hash{}
    var hashes = map[string]string{}

    for checksumAlgo := range checksumLookupTable {
        hashers[checksumAlgo] = checksumLookupTable[checksumAlgo].New()
    }

    fp, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    defer fp.Close()

    fpStat, _ := fp.Stat()
    fileSize := fpStat.Size()
    var totalRead = 0
    var bar *mpb.Bar
    var rp *bufio.Reader

    if progressBar != nil {
        bar = progressBar.AddBar(fileSize).
                PrependName(path, 0, mpb.DwidthSync|mpb.DidentRight).
                PrependCounters("%3s / %3s", mpb.UnitBytes, 18, mpb.DwidthSync|mpb.DextraSpace).
                AppendPercentage(3, 0)
        rp = bufio.NewReader(bar.ProxyReader(fp))
    } else {
        rp = bufio.NewReader(fp)
    }

    buffer := make([]byte, bufferSize)

    for {
        amountRead, err := rp.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }
        totalRead += amountRead
        if bar != nil {
            bar.Incr(amountRead)
        }

        for checksumAlgo := range checksumLookupTable {
            hashers[checksumAlgo].Write(buffer)
        }
    }

    for checksumAlgo := range checksumLookupTable {
        hashes[checksumAlgo] = hex.EncodeToString(hashers[checksumAlgo].Sum(nil))
        Trace.Printf("%v: Computed hash %v: %v\n", path, checksumAlgo, hashes[checksumAlgo])
    }
    return hashes, nil
}
