//multicopy is a multi-threaded URL retriever.  You provide
//a list of URLs to copy, multicopy reads each URL and stores
//the contents in the specified directory.  Multicopy retains
//the directory structure for each URL.  For example, a URL
//like
//
// https://whatever.where.bizi/og/blarf/5/bijj.png
//
// will appear in
//
// [dir]/og/blarf/5/bijj.png
//
// Installation:
//
// go get github.com/salsalabs/multicopy
//
// go install github.com/salsalabs/multicopy
//
// Execution:
//
// multicopy [options] data-file
//
// Help:
//
// multicopy --help
//
package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"

	"github.com/salsalabs/godig"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	//RepTemplate is the URL temlate for retrieving the contents of a dir.
	RepTemplate = `https://hq.salsalabs.com/salsa/include/fck2.5.1/editor/filemanager/browser/default/connectors/jsp/connector?Command=GetFoldersAndFiles&Type=Image&CurrentFolder=/%s`
)

//Connector is the wrapper for the rest of the XML-based structure.
type Connector struct {
	Command      string        `xml:"command,attr"`
	ResourceType string        `xml:"resourceType,attr"`
	Current      CurrentFolder `xml:"CurrentFolder"`
	Dirs         Folders       `xml:"Folders"`
	Entries      Files         `mxl:"Files"`
}

//CurrentFolder is the current folder being parsed.
type CurrentFolder struct {
	Path string `xml:"path,attr"`
	URL  string `xml:"url,attr"`
}

//Folders represents a list of folders.  Can be empty.
type Folders struct {
	XMLName xml.Name `xml:"Folders"`
	Entries []Folder `xml:"Folder"`
}

//Files represents a list of fileds. Can be empty.
type Files struct {
	XMLName xml.Name `xml:"Files"`
	Entries []Folder `xml:"File"`
}

//Folder represents a folder.  No contents, just the folder.
type Folder struct {
	Name string `xml:"name,attr"`
}

//File represents a file in the current folder.
type File struct {
	Name string `xml:"name,attr"`
	Size string `xml:"size,attr"`
}

//Load reads a URL and pushes file URLs onto the provided channel.
//avaialble from a Salsa Classic images and files repository. Load
//calls itself re-entrantly to process directores in the repository.
//Load returns when all directories have been examined.
func Load(api *godig.API, dir string, c chan string) error {
	u := fmt.Sprintf(RepTemplate, dir)
	log.Println(u)
	req, err := http.NewRequest("GET", u, nil)
	if err == nil {
		// Salsa's API needs these cookies to verify authentication.
		for _, c := range api.Cookies {
			req.AddCookie(c)
		}
		resp, err := api.Client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				//log.Printf("\n%v\n\n", string(body))
				var v Connector
				err := xml.Unmarshal(body, &v)
				if err != nil {
					return err
				}
				log.Printf("%+v\n", v)
			}
		}
	}
	return err
}

//Run reads URLs from a channel and writes them to disk.
//Errors are logged and are not fatal.  Processing continues
//until the done channel has contents or is closed.
func Run(c chan string, dir string, done chan bool) {
	var errLog = log.New(os.Stderr, "", log.LstdFlags)
	var stdLog = log.New(os.Stdout, "", log.LstdFlags)
	for {
		select {
		case u := <-c:
			_, err := Store(u, dir)
			if err != nil {
				errLog.Printf("Error: %v %s\n", err, u)
			} else {
				stdLog.Printf("%s\n", u)
			}
		case <-done:
			return
		default:
		}
	}
}

//Store saves a URL to disk.  The contents are stored starting in
//the provided directory keeping the URL's directory structure
//intact.
func Store(link string, dir string) (int64, error) {
	r, err := http.Get(link)
	if err != nil {
		return 0, err
	}
	if r.StatusCode != 200 {
		m := fmt.Sprintf("%v (%v)", r.StatusCode, http.StatusText(r.StatusCode))
		return 0, errors.New(m)
	}
	defer r.Body.Close()

	u, err := url.Parse(link)
	if err != nil {
		return 0, err
	}
	p := path.Join(dir, u.Path)
	d := path.Dir(p)
	err = os.MkdirAll(d, os.ModePerm)
	f, err := os.Create(p)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n, err := io.Copy(f, r.Body)
	if err != nil {
		return 0, err
	}
	f.Sync()
	return int64(n), nil
}

//main is the application.  Gathers arguments, starts listeners, reads
//URLs and processes them.
func main() {
	var (
		app   = kingpin.New("multicopy", "A command-line app to copy the contents of a list of URLs to a dir.")
		login = app.Flag("login", "YAML file with login credentials").Required().String()
		dir   = app.Flag("dir", "Store contents starting in this directory.").Default(".").String()
		count = app.Flag("count", "Start this number of processors.").Default("20").Int()
		//data  = app.Arg("data", "File of URLs to store, one per line.").Required().String()
	)
	app.Parse(os.Args[1:])

	api, err := (godig.YAMLAuth(*login))
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	var wg sync.WaitGroup
	c := make(chan string)
	done := make(chan bool)

	// Start processors.
	for i := 1; i <= *count; i++ {
		go func(i int) {
			wg.Add(1)
			defer wg.Done()
			Run(c, *dir, done)
		}(i)
	}

	// Queue up urls.  No buffering means that the URL
	// is not queued until there's a listener.
	repDir := ""
	err = Load(api, repDir, c)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	// Tells the processors that we're through.
	close(done)

	// Wait for everything to finish.
	wg.Wait()
}
