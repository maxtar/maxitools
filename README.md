# maxitools
Different tools for different IT tasks

---
## checksum
Tool for calculating different hash.

Now available the following hash algorithms:
* md5 (by default)
* sha1
* sha256
* sha512

### How to use
`checksum -f <path to file> [-algo ]`

Arguments order doesn't matter.

For example:

`cheksum -algo sha256 -f d:\usr\tmp.xml`

---
## wintime
Tool created for the linux time purposes - calculate program executing time.

### How to use
`wintime <command to run>`

---
## stubwebserver
Stub server for receieve any request and return parameters to output response.

---
todo: client for post
todo: HTTP POST any size files