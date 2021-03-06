# Multicopy
## Quickly copy images and files uploaded to Salsa Classic

## Background
Salsa's clients store images and files on Salsa so that they can 
be served up using a secure ("https://") URL.  The files are uploaded to the "images and files" repository.  You can learn more
about uploading images and files to Salsay by [clicking here](https://help.salsalabs.com/hc/en-us/articles/223342607-Upload-images-or-files-to-Salsa).
You can learn more about the images and files repository itself by [clicking here](https://help.salsalabs.com/hc/en-us/articles/223342727-Managing-files-uploaded-to-Salsa). 

As you can see from the doc, the images and files repository is not accessible except on a file-by-file basis.  That _can_ be done manually, but it's tedious and takes hours.

This package contains a Go application that reads a client's images and files repository.
Files from the repository are written to disk, retaining then directory structure from the repository.
The older, `wget`-based process usually consumed a couple of hours.  `Multicopy` does the same work in minute or two.
# Prerequisites
* Login credentials for the client.

    If you are a Salsa client and no not have login credentials then see the person in your organization that manages Salsa Classic.

    If you are a Salsa staffer, then create campaign manager credentials for yourself in client's Salsa HQ.

* The [Go](https://golang.org/doc/install) programming language.

* The correct directory structure for Go.

    Go requires a very specific directory structure in order to run correctly.  The easiest way to do this is to open a console window, navigate to your home directory and create a structure exactly like this.
  ```
    $HOME
      |
      + go
        |
        + bin
        + pkg
        + src
  ```

* Assure that the Go `bin` directory in your path.  [Click here](https://superuser.com/questions/284342/what-are-path-and-other-environment-variables-and-how-can-i-set-or-use-them) for a detailed discussion on the topic.  In Linux and OSX, the command is

    ```export PATH=$HOME/go/bin:$PATH```
  
  See the article for steps to add the Go bin dir to the Windows environment.
## Installation

1. If you have not done so already, install the [Go programming language](https://golang.org/doc/install).
1. Open a console window. The rest of the installation commands need to be typed or copied/pasted.
1. Retrieve the `godig` package.  It provides access to the Salsa Classic API.  
    ```bash
    go get github.com/salsalabs/godig
    ```
1. Install `godig`:
    ```bash
    go install github.com/salsalabs/godig
    ```
1. Retrieve the `multicopy` package:
    ```bash
    go get github.com/salsalabs/multicopy
    ```
1. Install `multicopy`:
    ```bash
    go install github.com/salsalabs/multicopy
    ```
1. You'll know that the installation is complete if you type
```bash
multicopy --help
```
and see something like this.
```bash
multicopy --help
usage: multicopy --login=LOGIN [<flags>]

A command-line app to copy images and files from a Salsa HQ to your disk.

Flags:
  --help         Show context-sensitive help (also try --help-long and --help-man).
  --login=LOGIN  YAML file with login credentials
  --dir="."      Store contents starting in this directory.
  --count=20     Start this number of processors.
  ```

## Login credentials

The `multicopy` application looks for your login credentials in a YAML file.  You provide the filename as part of the execution.
 The easiest way to get started is to  copy the `sample_login.yaml` file and edit it.  Here's an example.
```yaml
host: wfc2.wiredforchange.com
email: chuck@echeese.bizi
password: extra-super-secret-password!
```
The `email` and `password` are the ones that you normally use to log in. The `host` can be found by using [this page](https://help.salsalabs.com/hc/en-us/articles/115000341773-Salsa-Application-Program-Interface-API-#api_host) in Salsa's documentation.

Save the new login YAML file to disk.  We'll need it when the run `multicopy`.

## Execution

The `multicopy` program is created in the Go `bin` dir during installation.  To run `multicopy`, you'll need to open a console window and navigate to the directory where you'd like the files to be stored.

Once you're there, then copy your login YAML file to the same directory.

Let's say you stored your login credentials in `boffo.yaml`.  Typing
```bash
multicopy --login boffo.yaml
```
copies images and files from your Salsa HQ into the current directory.

If you type 
```bash
multicopy --login boffo.yaml --dir /tmp
```
then the results will be stored in `/tmp`.

In all cases, the images will be stored in a directory structure like this

```
  [current directory or --dir]
    |
    + o
      |
      + [[organization_KEY]]
        |
        + images
          |
          + (files)
          + (more folders)
```
If you are extracting images and files for a chapter, then the chapter is inserted just after the organization key.
```
  [current directory or --dir]
    |
    + o
      |
      + [[organization_KEY]]
        |
        + chapter
          |
          + [[chapter_KEY]]
            |
            + images
              |
              + (files)
              + (more folders)
```
You can learn more about the images and files repository directory structure by [clicking  here](https://help.salsalabs.com/hc/en-us/articles/223342727-Managing-files-uploaded-to-Salsa).

To save on client confusion, zip up the contents of the tree starting at "images" (wherever it is).  Clients will see their stuff fairly quickly and won't wonder why we have all of the directories starting at "o".

## License
Read the `LICENSE` file in this repository.
## Questions?  Comments?
Please don't waste your time by contacting Salsalabs support. Use the [Issues](https://github.com/salsalabs/multicopy/issues) link
in GitHub. 
