//+build mage

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/magefile/mage/sh"
	"github.com/stevelacy/daz"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
)

type til struct {
	Name     string
	Contents string
	category string
	Path     string
	Date     time.Time
}

type tilCategory struct {
	Name    string
	Entries []*til
}

var markdownHeaderRegexp = regexp.MustCompile(`(?m)# (.+)\n?`)

// listTILs walks the current working directory and finds all valid TILs
func listTILs() (x []*tilCategory, numTILs int, err error) {

	tempTILs := make(map[string][]*til) // category as key

	err = filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// exclude files that aren't Markdown and directories that are hidden
			if lcp := strings.ToLower(path); strings.HasPrefix(lcp, ".") || !strings.HasSuffix(lcp, ".md") {
				return nil
			}

			splitC := strings.Split(path, string(os.PathSeparator))
			if len(splitC) == 1 { // in the base directory
				return nil
			}

			var name string
			var contents string
			{
				fcont, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}

				// attempt to extract title from Markdown
				if len(markdownHeaderRegexp.Find(fcont)) != 0 {
					subm := markdownHeaderRegexp.FindSubmatch(fcont)
					name = strings.TrimSpace(string(subm[1]))
					var doneOnce bool
					fcont = bytes.TrimSpace(markdownHeaderRegexp.ReplaceAllFunc(fcont, func(b []byte) []byte {
						if doneOnce {
							return b
						} else {
							doneOnce = true
							return nil
						}
					}))
				} else {
					// remove everything from the path after last `.` - in practise, this means removing the file extension
					xp := strings.Split(splitC[len(splitC)-1], ".")
					name = strings.Join(xp[0:len(splitC)-1], ".")
				}

				contents = string(fcont)
			}

			category := strings.Join(splitC[0:len(splitC)-1], "/")

			date, err := getFileModDate(path)
			if err != nil {
				return err
			}

			tempTILs[category] = append(tempTILs[category], &til{
				Name:     name,
				category: category,
				Path:     path,
				Date:     date,
				Contents: contents,
			})

			return nil
		})

	for key, contents := range tempTILs {
		x = append(x, &tilCategory{
			Name:    key,
			Entries: contents,
		})
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i].Name < x[j].Name
	})

	for _, y := range x {
		numTILs += len(y.Entries)
	}

	return
}

// rewriteTILPaths clones `x` and replaces all occurences of `old` in each TIL's path with `new`
func rewriteTILPaths(old, new string, x []*tilCategory) []*tilCategory {

	newX := make([]*tilCategory, len(x))

	for i, cat := range x {
		ne := make([]*til, len(cat.Entries))
		for j, catEnt := range cat.Entries {

			y := *catEnt
			ne[j] = &y

			ne[j].Path = strings.ReplaceAll(ne[j].Path, old, new)
		}

		y := *cat
		newX[i] = &y
		newX[i].Entries = ne
	}

	return newX
}

const tilDateFormat = "2006-01-02"

// makeTILMarkdown generates Markdown from a []*tilCategory to make a list of TILs
func makeTILMarkdown(tils []*tilCategory) (string, error) {

	const headerLevel = "###"

	var sb strings.Builder

	for _, category := range tils {

		sb.WriteString(headerLevel)
		sb.WriteRune(' ')
		sb.WriteString(category.Name)
		sb.WriteString("\n\n")

		for _, til := range category.Entries {
			sb.WriteString("* [")
			sb.WriteString(til.Name)
			sb.WriteString("](")
			sb.WriteString(til.Path)
			sb.WriteString(") - ")
			sb.WriteString(til.Date.Format(tilDateFormat))
			sb.WriteRune('\n')
		}

		sb.WriteString("\n")

	}

	return strings.TrimSpace(sb.String()), nil
}

var markdownImageRegexp = regexp.MustCompile(`(?m)!\[.+]\((.+)\)`)

// listMarkdownImages lists all images in a Markdown document
func listMarkdownImages(markdown string) []string {
	x := markdownImageRegexp.FindAllStringSubmatch(markdown, -1)
	if len(x) == 0 {
		return nil
	}
	var y []string
	for _, z := range x {
		y = append(y, z[1])
	}
	return y
}

// renderAnchor renders a HTML anchor tag
func renderAnchor(text, url string, newTab bool) func() string {
	attrs := daz.Attr{
		"href": url,
		"rel":  "noopener",
	}
	if newTab {
		attrs["target"] = "_blank"
	}
	return daz.H("a", attrs, daz.UnsafeContent(text))
}

// makeTILHTML generates HTML from a []*tilCategory to make a list of TILs
func makeTILHTML(tils []*tilCategory) (string, error) {

	const headerLevel = "h3"

	var parts []interface{}
	for _, category := range tils {

		header := daz.H(headerLevel, category.Name)

		var entries []daz.HTML
		for _, til := range category.Entries {

			x := daz.H("li", daz.UnsafeContent(renderAnchor(til.Name, til.Path, false)()), " - "+til.Date.Format(tilDateFormat))
			entries = append(entries, x)
		}

		parts = append(parts, []daz.HTML{header, daz.H("ul", entries)})
	}

	return daz.H("div", parts...)(), nil
}

//go:embed page.template.html
var htmlPageTemplate []byte

// renderHTMLPage renders a complete HTML page
func renderHTMLPage(title, titleBar, pageContent, extraHeadeContent string) ([]byte, error) {

	tpl, err := template.New("page").Parse(string(htmlPageTemplate))
	if err != nil {
		return nil, err
	}
	outputBuf := new(bytes.Buffer)

	tpl.Execute(outputBuf, struct {
		Title            string
		Content          string
		PageTitleBar     string
		ExtraHeadContent string
	}{Content: pageContent, PageTitleBar: titleBar, Title: title, ExtraHeadContent: extraHeadeContent})

	time.Sleep(time.Second) // ratelimit?

	return outputBuf.Bytes(), nil
}

var markdownRenderer = goldmark.New(goldmark.WithExtensions(extension.GFM, highlighting.NewHighlighting(
	highlighting.WithStyle("github"),
)))

// renderMarkdownToHTML renders GitHub flavoured Markdown to HTML
func renderMarkdownToHTML(markdown string) ([]byte, error) {
	output := new(bytes.Buffer)
	err := markdownRenderer.Convert([]byte(markdown), output)
	if err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}

// getFileModDate gets the latest modification date from a tracked Git file. If no tracked file is found, the current date is returned
func getFileModDate(file string) (time.Time, error) {

	output, err := sh.Output("git", "log", "-1", "--format=%cd", file)
	if err != nil {
		return time.Time{}, err
	} else if output == "" {
		return time.Now(), nil
	}

	return time.Parse("Mon Jan 2 15:04:05 2006 -0700", output)
}

func joinPath(x ...string) string {
	return strings.Join(x, string(os.PathSeparator))
}

func copyFile(src, new string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(new)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

//go:embed README.template.md
var readmeTemplate []byte

func GenerateReadme() error {

	tils, numTILs, err := listTILs()
	if err != nil {
		return err
	}

	var outputReadmeBuf *bytes.Buffer
	{
		tpl, err := template.New("readme").Parse(string(readmeTemplate))
		if err != nil {
			return err
		}
		outputReadmeBuf = new(bytes.Buffer)
		md, err := makeTILMarkdown(tils)
		if err != nil {
			return err
		}
		tpl.Execute(outputReadmeBuf, struct {
			NumTIL int
			TILs   string
		}{numTILs, md})
	}

	{
		err := ioutil.WriteFile("README.md", outputReadmeBuf.Bytes(), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateHTML() error {

	const outputDir = ".site"

	tils, numTILs, err := listTILs()
	if err != nil {
		return err
	}

	const pageTitle = "akp's TILs"

	head := daz.H(
		"div",
		daz.H("h1", pageTitle),
		daz.H(
			"p",
			daz.UnsafeContent(
				fmt.Sprintf(
					"There are currently %d TILs<br>Last modified %s<br>Repo: %s",
					numTILs,
					time.Now().Format(tilDateFormat),
					renderAnchor("<code>codemicro/til</code>", "https://github.com/codemicro/til", false)(),
				),
			),
		),
	)

	htmlTils := rewriteTILPaths(".md", ".html", tils)

	tilHTML, err := makeTILHTML(htmlTils)
	if err != nil {
		return err
	}

	outputContent, err := renderHTMLPage(pageTitle, head(), tilHTML, "")
	if err != nil {
		return err
	}

	_ = os.Mkdir(outputDir, os.ModeDir)

	err = ioutil.WriteFile(joinPath(outputDir, "index.html"), outputContent, 0644)
	if err != nil {
		return err
	}

	for _, tilCat := range tils {
		_ = os.Mkdir(joinPath(outputDir, tilCat.Name), os.ModeDir)

		for _, til := range tilCat.Entries {

			pathDir := filepath.Dir(til.Path)

			// copy images
			for _, relPath := range listMarkdownImages(til.Contents) {
				err = copyFile(joinPath(pathDir, relPath), joinPath(outputDir, tilCat.Name, relPath))
				if err != nil {
					return err
				}
			}

			// render and save as HTML
			mdHTML, err := renderMarkdownToHTML(til.Contents)
			if err != nil {
				return err
			}

			head := daz.H(
				"div",
				daz.H("h1", til.Name),
				daz.H(
					"p",
					daz.UnsafeContent(
						fmt.Sprintf(
							"%s<br>Date: %s<br>Category: %s",
							renderAnchor("Back to index", "../", false)(),
							til.Date.Format(tilDateFormat),
							tilCat.Name,
						),
					),
				),
			)

			renderedHTML, err := renderHTMLPage(fmt.Sprintf("%s - %s", til.Name, pageTitle), head(), string(mdHTML), "")
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(joinPath(outputDir, tilCat.Name, strings.ReplaceAll(filepath.Base(til.Path), ".md", ".html")), renderedHTML, 0644)
			if err != nil {
				return err
			}

		}

	}

	// Tasks:
	// 2. Make category directories
	// 3. Copy images to correct directories
	// 4. Render and save each TIL as HTML

	return nil
}
