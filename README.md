# Multicopy
## Quickly copy the contents of a list of URLs to disk.

## Background
Salsa's clients store images and files on Salsa so that they can 
be served up using a secure ("https://") URL.  You can learn more
about the images and file directory by [clicking here](https://help.salsalabs.com/hc/en-us/articles/223342607-Upload-images-or-files-to-Salsa).
There is a detailed description of how the images and files repository works [here](https://help.salsalabs.com/hc/en-us/articles/223342727-Managing-files-uploaded-to-Salsa). 

Salsa has a process for retrieving a client's images and files.
The process is described in [this GitHub repository](https://gist.github.com/salsalabs/7c1c69f9cae6280a5a8f).  That
process uses `wget` to retrieve files from Salsa.  The constraint
is that `wget` can only retrieve one file at a time due to the way
that Salsa stores files.  (Hint: Salsa does not allow access to
the directories that contain the files.)

This app speeds things up by having several processors save images
and files at the same time.
# Prerequisites
* The [Go](https://golang.org/doc/install) programming language.
* The correct directory structure.  Here's a sample.
```
$HOME
  + go
    + bin
    + pkg
    + src
```
* The `bin` directory in your path.  For example.
```bash
export PATH=$HOME/go/bin:$PATH
```
* A list of URLs to retrieve in a file, one URL per line, no spaces
or commas (or semicolons or pipes or tabs.)

## Installation

Retrieve the `multicopy` package:
```bash
go get github.com/salsalabs/multicopy
```
Install `multicopy`:
```bash
go install github.com/salsalabs/multicopy
```
You'll know that the installation is complete if you type
```bash
multicopy --help
```
and see something like this.
```bash
usage: multicopy [<flags>] <data>

A command-line app to copy the contents of a list of URLs to a dir.

Flags:
  --help      Show context-sensitive help (also try --help-long and --help-man).
  --dir="."   Store contents starting in this directory.
  --count=20  Start this number of processors.

Args:
  <data>  File of URLs to store, one per line.
```
## Execution
Let's say you stored the list of urls into `boffo.txt`.  Here's how
to store the contents of the URLs in `boffo.txt` to the current 
directory.
```bash
multicopy boffo.txt
```
Here's how to store the contents into `/tmp`.
```bash
multicopy --dir /tmp boffo.txt
```
The output files will be stored starting in the current directory or directory that you choose.

Here's an example.  Let's say that this is the image URL.

`https://a.b.bizi/ochre/ogre/image.png`

and this is the `multicopy` command line.
```bash
multicopy --dir /home/me/mine boffo.txt
```
`Multicopy` will store the file in

`/home/me/mine/ochre/ogre/image.png`

Note that `multicopy` will automatically any directories needed
to store stuff.

## License
Read the `LICENSE` file in this repository.
## Questions?  Comments?
Please don't waste your time by contacting Salsalabs support. Use the [Issues](https://github.com/salsalabs/multicopy/issues) link
in GitHub. 
