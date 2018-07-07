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
* A list of URLs to retrieve in a file.  [Click here]((https://gist.github.com/salsalabs/7c1c69f9cae6280a5a8f)) to see how to retrieve the list of URLs.  The list is actually a bash script.
Open the bash script with an editor and do the following.
    * Sort the contents of the file.
    * Remove duplicate lines (typically there's a `Unique lines` command in the editor).
    * Remove all lines that do not start with `${CMD}`.
    * Remove "`${CMD}` " from each line.  (There's a trailing space there...)
    * Save the file in the current directory.

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
``bash
multicopy boffo.txt
```
Here's now to store the contents into `/tmp`.
```bash
multicopy --dir/tmp boffo.txt
```
## License
Read the `LICENSE` file in this repository.
## Questions?  Comments?
Use the [Issues](https://github.com/salsalabs/multicopy/issues) link
in GitHub.
