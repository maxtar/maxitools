# maxitools
Different tools for different IT tasks

---
## checksum
Tool for calculating different hashes.

Now available the following hash algorithms:
* md5 (by default)
* sha1
* sha256
* sha512

### How to use
`checksum -f <path to file> [-algo ]`

Arguments order doesn't matter.

For example:

`checksum -algo sha256 -f d:\usr\tmp.xml`

---
## wintime
Tool created for the linux "time" purposes - calculate program executing time.

### How to use
`wintime <command to run>`

---
## stubwebserver
Stub server receieve any request and return parameters to output response.

### Command line arguments
* **p** - server port. By default: 8080.
* **logdir** - directory where will be saved log-file with requests. If *none* (default value) requests will not be saved.
* **stdout** - enable print requests to standart output stream. *true* by default.
* **pfd** - directory where will be saved files from POST multidata-form requests. If *none* (default value) files will not be saved.

For example:

`stubserver -p 8080 -logdir "../log" -stdout`

todo: add proxy mode

---
## udpstubserver
Server listen UDP port and write received packet to console.

### Command line arguments
* **p** - server port. By default: 8125.

For example:

`udpstubserver -p 8125`

---
todo: client for post

todo: HTTP POST any size files