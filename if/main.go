
if необязательнаяИнструкция1; логическоеВыражение1 {
 блок1
} else if необязательнаяИнструкция2; логическоеВыражение2 {
 блок2
} else {
 блок3
}

func main() {
    a, b := 4, 5
    if a < b {
        fmt.Println("a is less than b")
    } else if a > b {
        fmt.Println("a is greater than b")
    } else {
        fmt.Println("a is equal to b")
    } 
}


func main() {
    
    if a, b := 4, 5; a < b {
        fmt.Println("a is less than b")
    } else if a > b {
        fmt.Println("a is greater than b")
    } else {
        fmt.Println("a is equal to b")
    } 
}



func ArchiveFileList(file string) ([]string, error) {
	if suffix := Suffix(file); suffix == ".gz" {
		return GzipFileList(file)
	} else if suffix == ".tar" || suffix == ".tar.gz" || suffix == ".tgz" {
		return TarFileList(file)
	} else if suffix == ".zip" {
		return ZipFileList(file)
	}
	return nil, errors.New("unrecognized archive")
}

func Suffix(file string) string {
	file = strings.ToLower(filepath.Base(file))
	if i := strings.LastIndex(file, "."); i > -1 {
		if file[i:] == ".bz2" || file[i:] == ".gz" || file[i:] == ".xz" {
			if j := strings.LastIndex(file[:i], "."); j > -1 && strings.HasPrefix(file[j:], ".tar") {
				return file[j:]
			}
		}
	return file[i:]
	}
	return file
}