# filebeatutils

Some useful utils, for older versions of Elastic Filebeat, that keep
its repository in a single file, 
/var/lib/filebeat/registry

dumpbeatregistry obviously just dumps the registry in human readable form.
Basically it gives you a list of files being tracked

checkbeat checks to see if a file is in the registry, AND up to date
 (ie: fully archived)
Note that it does so by device and inode lookup, so it properly handles renames
just like filebeat itself


These are "go" programs. To use, install a golang compiler or runtime, 
and use either

  go run checkbeat.go

or

  go build checkbeat.go 
  
  ./checkbeat  ....args here...




