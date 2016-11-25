package main

import (
    "fmt"
    "os"
    "flag"
    "encoding/gob"
    "log"
    "io"
    "path/filepath"
    "crypto/md5"
)

type MetaFile struct {
	Path	string
//	Stat	os.FileInfo
	Size	int64
	Hash	[]byte
}
type MetaFileList map[string]MetaFile

type Mode int

const (
	ModeInit  Mode = iota
	ModeDedup      = iota
)

func processFile(fileName string) (*MetaFile, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, nil
	}
	defer file.Close()

	hash := md5.New()
	io.Copy(hash, file)

	file.Close()

	return &MetaFile{
		Path: fileName,
//		Stat: stat,
		Size: stat.Size(),
		Hash: hash.Sum(nil),
	}, nil
}

var (
	flagVerbose = flag.Bool("verbose", false, "verbose logging")
)

func main() {
	var flagSilos = flag.String("silos", ".silos", "silos path for md5 sums")
	var flagInit = flag.Bool("init", false, "init silos")
	var flagDedup = flag.Bool("dedup", false, "dedup stuff to silos")

	flag.Parse()

	if (len(flag.Args()) != 1) {
		flag.Usage()
		log.Fatal("bleh")
	}

	var searchDir = flag.Args()[0]
	silosFile := *flagSilos // + "/silos.go"

	fmt.Println("==>", searchDir, silosFile, *flagVerbose)

	if (*flagInit == false && *flagDedup == false) {
		flag.Usage()
	}

	if (*flagInit) {
		silosOp(silosFile, searchDir, ModeInit)
	} else if (*flagDedup) {
		silosOp(silosFile, searchDir, ModeDedup)
	}
}

func silosFileIterator(path string, f os.FileInfo, err error, ma *MetaFileList, whichMode Mode) error {
	if (f.IsDir()) {
		return nil
	}
	var metaFile, _ = processFile(path)
	var key = fmt.Sprintf("%x", metaFile.Hash)
	_, key_exists := (*ma)[key]
	if (key_exists) {
		fmt.Println(path, "EXISTS")
	} else {
		(*ma)[key] = *metaFile
		if (*flagVerbose) {
			fmt.Println(path)
		}
	}
	return nil
}

func silosOp(silosFile string, searchPath string, whichMode Mode) {
	var metaFileList = make(MetaFileList)
	if (whichMode == ModeDedup) {
		metaFileList = silosRead(silosFile)
	}
	err := filepath.Walk(searchPath, func(path string, f os.FileInfo, err error) error {
		return silosFileIterator(path, f, err, &metaFileList, whichMode)
	})
	if (err != nil) {
		log.Fatal(err)
	}
	silosWrite(silosFile, metaFileList)
}

func silosWrite(silosFile string, metaFileList MetaFileList) {
	encodeFile, err := os.Create(silosFile)
	if err != nil {
		panic(err)
	}

	encoder := gob.NewEncoder(encodeFile)
	if err := encoder.Encode(metaFileList); err != nil {
		panic(err)
	}
	encodeFile.Close()
}

func silosRead(silosFile string) MetaFileList {
	decodeFile, err := os.Open(silosFile)
	if err != nil {
		panic(err)
	}
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)
	metaFileList := make(map[string]MetaFile)
	decoder.Decode(&metaFileList)
	return metaFileList
}
