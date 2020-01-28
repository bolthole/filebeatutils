
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
)

/* Based on file beat 5.6 registry format */
type OSFileRef struct {
	Inode uint64
	Device uint64 /* devid num */
}
type FilebeatEntry struct {
	Source string
	Offset int64
	OSFileRef OSFileRef
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

	jsonbytes,_ := ioutil.ReadFile(Registrypath)

	json.Unmarshal(jsonbytes,  &regentries)


	for _, l := range regentries {
		 fmt.Println( l.Source, l.Offset)
	}
	fmt.Println("End of ", Registrypath)
	os.Exit(1)

}
