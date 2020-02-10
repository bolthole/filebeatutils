
/* attempt to write something to check if 
 * filebeat is "done" with a file
 *
 */

package main


import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"golang.org/x/sys/unix"
)

/* Based on file beat 5.6 registry format */
type FileStateOS struct {
	Inode uint64
	Device uint64 /* devid num */
}
type FilebeatEntry struct {
	Source string
	Offset int64
	FileStateOS FileStateOS
	// Timestamp string  /* we dont actually care about this yet */
	// Ttl int /* meh */
}

var jsonbytes []byte
var regentries []FilebeatEntry

/* theoretical exit codes:
*    0 - file found and fully processed
*    1 - file found but NOT fully processed
*    2 - file not found
*/

func main() {
	var Registrypath = "/var/lib/filebeat/registry"

	if(len(os.Args) < 2) {
		fmt.Println("Need a filename")
		os.Exit(1)
	}
	targetfile := os.Args[1]
	var targetstat unix.Stat_t

	fmt.Println ("Debug: target file is ",targetfile)

	unix.Lstat(targetfile, &targetstat)


	fmt.Println ("Debug: dev, inode, length:",
		targetstat.Dev, targetstat.Ino, targetstat.Size)

	jsonbytes,_ := ioutil.ReadFile(Registrypath)
	// jsonbytes,_ := ioutil.ReadFile("registry.json")


	json.Unmarshal(jsonbytes,  &regentries)

	//fmt.Printf("All the strings\n")
	//fmt.Printf("%+v",regentries)

	for _, l := range regentries {
		// fmt.Println( l.Source, l.Offset)

		if((l.FileStateOS.Device == targetstat.Dev) && (l.FileStateOS.Inode == targetstat.Ino)) {
			fmt.Println("File found in registry")

			if(l.Offset == targetstat.Size){
				fmt.Println("File is fully processed")
				os.Exit(0)
			} else {
				fmt.Println("File %s is NOT fully processed!",targetfile)
				os.Exit(1)
			}
		}
	}
	fmt.Println("File not found in registry!")
	os.Exit(1)

}
