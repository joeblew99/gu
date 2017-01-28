package parse

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	// shellComponentMarker defines the marker used to expose
	// to parser to check resource metas.
	shellComponentMarker = "shell:component"

	// shellGlobalMarker defines the marker used to expose
	// to parser to check resource metas.
	shellGlobalMarker = "shell:component:global"

	// resourceMarker defines the marker used to declare a resource
	// call.
	resourceMarker = "Resource"

	// resourceEnclosureBegin defines the marker used to define the
	// beginning of a resource
	resourceEnclosureBegin = "{"

	// resourceEnclosureEnd defines the marker used to define the
	// end of a resource
	resourceEnclosureEnd = "}"

	// resourceHeader combines the expect value for the resource declaration
	resourceHeader = resourceMarker + " " + resourceEnclosureBegin

	// resourceKVSeperator defines the token used seperate a resource field key
	// to it's value.
	resourceKVSeperator = ":"

	// resourceText defines the token used to enclose a text content and it
	// acts as the beginning and ending closure for the given text
	resourceText = "```"

	// resourceDefinitionSeperator defines the seperator token used
	// to define a new resource field.
	resourceDefinitionSeperator = ",\n"

	// multipleSpaces defines a regexp to match multiple spaces
	multipleSpaces = regexp.MustCompile("\\s+")

	// newLines defines a regexp for a newLine
	newLine = regexp.MustCompile("\n")

	// commentLine defines the giving tokens to expect for comments.
	commentLine = regexp.MustCompile("//|/*")

	// ErrInvalidResourceEnding defines the error returned when a giving
	// resource statement ends in a invalid state.
	ErrInvalidResourceEnding = errors.New("Invalid Resource Ending")

	// ErrInvalidCharacterFound defines the error returned when a charater is found
	// out of scope.
	ErrInvalidCharacterFound = errors.New("Invalid Resource Ending")

	// ErrInvalidFieldSet defines the errors returned when a field key value pair are
	// invalid.
	ErrInvalidFieldSet = errors.New("Invalid Resource Field Key Value set")

	// ErrInvalidEndCallWithNoResource defines the error returned when a ending
	// '}' symbol is found with no actual resource header initalized earlier.
	ErrInvalidEndCallWithNoResource = errors.New("Invalid End With No Resource Declaration")

	// ErrInvalidResourceField defines the error returned when a giving
	// resource statement ends in a invalid state.
	ErrInvalidResourceField = errors.New("Invalid Resource Field")

	// ErrInvalidTextWrap defines the error returned when a giving
	// resource statement text field is not properly delimited by a ```.
	ErrInvalidTextWrap = errors.New("Invalid Text Delimiter Wrap")
)

// GrabShellDoc retrieves all resource documentation which are detailed within
// the comments of components. It grabs all these into single collective and
// seperate strings which then gets sliced into actual Resource structs.
func GrabShellDoc(text string) ([]map[string]string, error) {
	resources := make([]map[string]string, 0)

	blocks, err := parseResourceMetas(text)
	if err != nil {
		return nil, err
	}

	for _, block := range blocks {
		block = strings.TrimPrefix(block, resourceHeader)
		block = strings.TrimSuffix(block, resourceEnclosureEnd)
		block = strings.TrimPrefix(block, "\n")
		block = strings.TrimSuffix(block, "\n")

		reader := strings.NewReader(block)
		bufReader := bufio.NewReader(reader)

		var textCollect bool
		var fieldName string
		var texts []string

		{
			res := make(map[string]string)

		readerLoop:
			for {
				scan, err := bufReader.ReadString('\n')
				if err == io.EOF {
					break readerLoop
				}

				nScan := strings.TrimSpace(scan)

				if textCollect {
					if nScan == "```" {
						textCollect = false
						res[capitalize(fieldName)] = strings.Join(texts, "")
						continue readerLoop
					}

					texts = append(texts, scan)
					continue readerLoop
				}

				nFields := strings.SplitN(nScan, ":", 2)
				if !textCollect && len(nFields) < 2 {
					return resources, ErrInvalidFieldSet
				}

				fname, fieldVal := nFields[0], strings.TrimSpace(nFields[1])

				if fieldVal == "```" {
					textCollect = true
					fieldName = fname
					continue readerLoop
				}

				res[capitalize(fname)] = fieldVal
			}

			resources = append(resources, res)
		}

	}

	return resources, nil
}

// parseResourceMetas retrieves all resource documentation which are detailed within
// the comments of components. It grabs all these into single collective slice
// of individual resource texts to be parsed.
func parseResourceMetas(text string) ([]string, error) {
	var blocks []string

	// split the text by newlines.
	docs := newLine.Split(text, -1)

	var preblocks [][]string

	for index, line := range docs {
		line = strings.TrimSpace(line)
		if newLine.MatchString(line) {
			continue
		}

		line = multipleSpaces.ReplaceAllString(line, " ")
		if line == " " || line == "" {
			continue
		}

		if line == shellComponentMarker || line == shellGlobalMarker {
			preblocks = append(preblocks, docs[index+1:])
			break
		}
	}

	var inResourceMode bool
	var block []string
	var blockStartCounter int
	var textCounter int
	var textblock bool

	for _, item := range preblocks {

	blockLoop:
		for _, pre := range item {
			npre := strings.TrimSpace(pre)

			if npre == "" {
				continue
			}

			switch npre {
			case resourceHeader:
				block = nil

				// If we are already collecting resource then explode
				// due to new resource header.
				if inResourceMode {
					return blocks, ErrInvalidResourceEnding
				}

				inResourceMode = true

				// We need to add back the newline back again has its the seperator format.
				block = append(block, pre+"\n")
				continue

			case "}", "}\n", "}\n\r":

				// If we are catching a case of a resource ending token then
				// explode with error.
				if !inResourceMode && blockStartCounter < 1 && !textblock {
					return blocks, ErrInvalidEndCallWithNoResource
				}

				if textblock {
					// We need to add back the newline back again has its the seperator format.
					block = append(block, pre+"\n")
					continue
				}

				// We need to add back the newline back again has its the seperator format.
				block = append(block, pre+"\n")
				if blockStartCounter > 0 {
					blockStartCounter--
					continue
				}

				blocks = append(blocks, strings.Join(block, ""))
				inResourceMode = false

				continue blockLoop
			}

			if !inResourceMode {
				return blocks, ErrInvalidEndCallWithNoResource
			}

			if textblock {
				// We need to add back the newline back again has its the seperator format.
				block = append(block, pre+"\n")

				if strings.Contains(npre, "```") {
					textCounter--

					if textCounter < 1 {
						textblock = false
					}
				}

				continue
			}

			// Do we have a some show of the {} closure characters?
			// Then increment and decrement count accordingly.
			if strings.Contains(npre, "{") {
				blockStartCounter++
			}

			if strings.Contains(npre, "}") {
				blockStartCounter--
			}

			if strings.Contains(npre, "```") {
				textCounter++
				textblock = true
			}

			// We need to add back the newline back again has its the seperator format.
			block = append(block, pre+"\n")
		}
	}

	return blocks, nil
}

var remoteSchemes = []string{"http", "https", "ssl", "git", "ftp"}

// toResources returns a slice of ResourceCollection from the provided
// map slices.
func toResources(res []map[string]string) ([]ResourceCollection, error) {
	var resc []ResourceCollection

	for _, rsc := range res {
		var r ResourceCollection

		if _, ok := rsc["Name"]; !ok {
			return nil, errors.New("Resource Map lacks 'Name' field")
		}

		if remote, err := strconv.ParseBool(rsc["Remote"]); err == nil {
			r.Remote = remote
			delete(rsc, "Remote")
		}

		if initd, err := strconv.ParseBool(rsc["Init"]); err == nil {
			r.Init = initd
			delete(rsc, "Init")
		} else {
			r.Init = true
		}

		if isglobal, err := strconv.ParseBool(rsc["Global"]); err == nil {
			r.Global = isglobal
			delete(rsc, "Global")
		}

		if localize, err := strconv.ParseBool(rsc["Localize"]); err == nil {
			r.Localize = localize
			delete(rsc, "Localize")
		}

		if enc, err := strconv.ParseBool(rsc["Encode"]); err == nil {
			r.Encode = enc
			delete(rsc, "Encode")
		} else {
			r.Encode = true
		}

		if encoded, err := strconv.ParseBool(rsc["Encoded64"]); err == nil {
			r.Encoded = encoded
			delete(rsc, "Encoded64")
		}

		r.Name = rsc["Name"]
		delete(rsc, "Name")

		r.Cache = rsc["Cache"]
		delete(rsc, "Cache")

		path := strings.TrimSpace(rsc["Path"])

		r.Path = path
		delete(rsc, "Path")

		if path != "" && !r.Remote {
			if purl, err := url.Parse(r.Path); err == nil {
				for _, scheme := range remoteSchemes {
					if purl.Scheme == scheme {
						r.Remote = true
						break
					}
				}
			}
		}

		content := rsc["Content"]
		r.Data = strings.TrimSpace(content)
		delete(rsc, "Content")

		r.HookName = rsc["Hook"]
		delete(rsc, "Hook")

		r.ContentSize = rsc["Size"]
		delete(rsc, "Size")

		r.Meta = rsc

		resc = append(resc, r)
	}

	return resc, nil
}

// ErrNotLocalPath is returned when the provided path is not an absolute file path,
// or relative path of a resource located within the local filesystem.
var ErrNotLocalPath = errors.New("Path is not a local resource but external URL")

var windowsAddr = regexp.MustCompile("^\\w$")

// getURLContent retrieves the giving url content for the provided path.
func getURLContent(path string, encodeb64 bool) ([]byte, error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var buff bytes.Buffer
	_, err = io.Copy(&buff, res.Body)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		if encodeb64 {
			return []byte(base64.StdEncoding.EncodeToString(buff.Bytes())), nil
		}

		return buff.Bytes(), nil
	}

	return nil, fmt.Errorf("Error Response: Status[%d] Content[%q]", res.StatusCode, buff.String())
}

// getFileContent pulls the content of a file path from the provided path
// used against the pkg path if it is a relative path.
func getFileContent(pkg string, path string, encodeb64 bool) ([]byte, error) {
	uri, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	if uri.Scheme != "" && uri.Scheme != "file" && !windowsAddr.MatchString(uri.Scheme) {
		return nil, ErrNotLocalPath
	}

	if uri.Scheme == "file" {
		if strings.HasPrefix(uri.Path, "/") {
			return getFileData(uri.Path[1:], encodeb64)
		}

		return getFileData(uri.Path, encodeb64)
	}

	if windowsAddr.MatchString(uri.Scheme) {
		return getFileData(uri.String(), encodeb64)
	}

	if uri.Scheme == "" && !strings.HasPrefix(path, "/") {
		return getFileData(filepath.Join(goSrcPath, pkg, path), encodeb64)
	}

	if uri.Scheme == "" && strings.HasPrefix(path, ".") {
		return getFileData(filepath.Join(goSrcPath, pkg, path), encodeb64)
	}

	return getFileData(path, encodeb64)
}

// getFileData retrieves the content of the provided file path.
func getFileData(path string, encodeb64 bool) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	_, err = io.Copy(&b, file)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if encodeb64 {
		return []byte(base64.StdEncoding.EncodeToString(b.Bytes())), nil
	}

	return b.Bytes(), nil
}

// capitalize returns the string capitalized.
func capitalize(m string) string {
	return strings.ToUpper(m[:1]) + m[1:]
}
