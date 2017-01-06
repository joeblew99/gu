package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gu-io/gu/shell"
	"github.com/gu-io/gu/shell/parse"

	"gopkg.in/urfave/cli.v2"
)

var (
	version  = "0.0.1"
	commands = []*cli.Command{}

	usage = `Provides a CLi tool which provides different commands to build and prepare
Gu based projects for testing, deployment and push.`

	aferoTemplate = `// Package %s is auto-generated and should not be modified by hand.
// This package contains a virtual file system for generate resources which are not accessed
// through a remote endpoint (i.e those resources generated from the manifests that are local in the
// filesystem and are not marked as remote in access).
package %s

import (
	"path/filepath"

	"github.com/spf13/afero"
)

// AppFS defines the global handler for which all assets generated from manifests
// files which are not remote resources are provided as binary embedded assets.
var AppFS = afero.NewMemMapFs()

// addFile adds a giving file name into the file system.
func addFile(path string, content []byte){
	dir, _ := filepath.Split(path)
	if dir != "" {
		AppFS.MkdirAll(dir,0755)
	}

	afero.WriteFile(AppFS, path, content, 0644)
}

func init(){
%+s
}

`
)

func main() {
	initCommands()

	app := &cli.App{}
	app.Name = "Gushell"
	app.Version = version
	app.Commands = commands
	app.Usage = usage

	app.Run(os.Args)
}

func generateAddFile(name string, content []byte) string {
	return fmt.Sprintf(`
		addFile(%q, []byte(%+q))
`, name, content)
}

func initCommands() {
	commands = append(commands, &cli.Command{
		Name:        "generate-vfs",
		Usage:       "gushell generate-vfs <PackageName>",
		Description: "Generate-VFS generates a new package which loads the resources loaded from the package, creating a new package which can be loaded and used to virtually serve the resources through a virtual filesystem",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output-dir",
				Aliases: []string{"dir"},
			},
			&cli.StringFlag{
				Name:    "packageName",
				Aliases: []string{"pkg"},
			},
		},
		Action: func(ctx *cli.Context) error {
			cdir, err := os.Getwd()
			if err != nil {
				return err
			}

			dirName := ctx.String("packageName")
			packageName := ctx.String("packageName")

			if packageName == "" {
				dirName = "manifests"
				packageName = "manifests"
			}

			res, err := parse.ShellResources(cdir)
			if err != nil {
				return err
			}

			var bu bytes.Buffer
			for _, rs := range res {
				manifest, merr := rs.GenManifests()
				if merr != nil {
					return merr
				}

				for _, attr := range manifest.Manifests {
					if !attr.Remote && attr.Content != "" {
						bu.WriteString(generateAddFile(attr.Name, []byte(attr.Content)))
					}
				}
			}

			contents := fmt.Sprintf(aferoTemplate, packageName, packageName, bu.String())

			outdir := ctx.String("output-dir")
			if outdir == "" {
				outdir = cdir
			}

			if outdir != "" && !filepath.IsAbs(outdir) {
				outdir = filepath.Join(cdir, outdir)
			}

			if merr := os.MkdirAll(filepath.Join(outdir, dirName), 0755); merr != nil {
				return merr
			}

			manifestFile, err := os.Create(filepath.Join(outdir, dirName, "manifest.go"))
			if err != nil {
				return err
			}

			defer manifestFile.Close()

			total, err := manifestFile.Write([]byte(contents))
			if err != nil {
				return err
			}

			if total != len(contents) {
				return errors.New("Data written is incomplete")
			}

			return nil
		},
	})

	commands = append(commands, &cli.Command{
		Name:        "generate",
		Usage:       "gushell generate",
		Description: "Generate parses the current directory which it assumes is a Go pkg directory and creates a manifest.json file to contain all generated resources from the meta comments within the package and it's imports",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output-dir",
				Aliases: []string{"dir"},
			},
		},
		Action: func(ctx *cli.Context) error {
			cdir, err := os.Getwd()
			if err != nil {
				return err
			}

			res, err := parse.ShellResources(cdir)
			if err != nil {
				return err
			}

			var manifests []*shell.AppManifest

			for _, rs := range res {
				manifest, merr := rs.GenManifests()
				if merr != nil {
					return merr
				}

				manifests = append(manifests, manifest)
			}

			manifestJSON, err := json.MarshalIndent(manifests, "", "\t")
			if err != nil {
				return err
			}

			outdir := ctx.String("output-dir")
			if outdir == "" {
				outdir = cdir
			}

			if outdir != "" && !filepath.IsAbs(outdir) {
				outdir = filepath.Join(cdir, outdir)
			}

			manifestFile, err := os.Create(filepath.Join(outdir, "manifest.json"))
			if err != nil {
				return err
			}

			defer manifestFile.Close()

			total, err := manifestFile.Write(manifestJSON)
			if err != nil {
				return err
			}

			if total != len(manifestJSON) {
				return errors.New("Data written is incomplete")
			}

			return nil
		},
	})
}
