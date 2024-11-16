package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

// Spec
// 1. Read input into mem
// 2. Specify output sink
// 3. Render
// 4. (preview only) Open OS's web page tool for preview
// 5. (preview only) Delete the temporary file which contains render output for preview
func main() {
	// By default, read from stdin and output to stdout
	// TODO support multiple input files
	inPath := flag.String("i", "-", "Input file path")
	// TODO use temporary filesystem
	// In preview mode we open the rendered file w/ OS's default web page viewer tool (usually a web browser)
	// After that we remove the rendered file
	previewOnly := flag.Bool("preview", false, "Preview only")

	flag.Parse()
	var mdTxtReader io.Reader = os.Stdin
	if p := *inPath; p != "" && p != "-" {
		f, err := os.Open(p)
		if err != nil {
			panic(fmt.Errorf("error opening input file %s: %w", p, err))
		}
		mdTxtReader = f
	}
	mdTxt, err := io.ReadAll(mdTxtReader)
	if err != nil {
		panic(fmt.Errorf("error reading all Markdown content from input: %w", err))
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)
	var sink io.Writer = os.Stdout
	// path to the temp file which contains markdown render output
	var tmpOut string
	if *previewOnly {
		tmpDir, err := os.MkdirTemp("", "rmd")
		if err != nil {
			panic(fmt.Errorf("error creating temp directory: %w", err))
		}
		// clean up upon exit
		defer func() {
			if err := os.RemoveAll(tmpDir); err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("error removing temporary directory %s: %w", tmpDir, err))
			}
		}()

		tmpOut = path.Join(tmpDir, "out.html")
		if f, err := os.Create(tmpOut); err != nil {
			panic(fmt.Errorf("error creating temp directory: %w", err))
		} else {
			sink = f
		}
	}
	// output rendered data to stdout to leverage existing shell tooling
	if err := md.Convert(mdTxt, sink); err != nil {
		panic(fmt.Errorf("error rendering Markdown: %w", err))
	}

	if *previewOnly {
		// only support OSX for now; TODO cover Linux as well
		if err := exec.Command("open", tmpOut).Run(); err != nil {
			panic(fmt.Errorf("error opening OS's default web page viewer: %w", err))
		}
		// NOTE this a hack to let the external tool read the rendered data before we perform cleanup
		// which is prone to race condition; Any better way to eliminate the race condition?
		time.Sleep(1 * time.Second)
	}
}
