//+build mage

package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/magefile/mage/sh"
)

type til struct {
	Name     string
	category string
	Path     string
	Date     time.Time
}

type tilCategory struct {
	Name    string
	Entries []*til
}

var mdHeaderRegexp = regexp.MustCompile(`(?m)# (.+)\n?`)

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
			{
				fcont, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}

				// attempt to extract title from Markdown
				if len(mdHeaderRegexp.Find(fcont)) != 0 {
					subm := mdHeaderRegexp.FindSubmatch(fcont)
					name = strings.TrimSpace(string(subm[1]))
				} else {
					// remove everything from the path after last `.` - in practise, this means removing the file extension
					xp := strings.Split(splitC[len(splitC)-1], ".")
					name = strings.Join(xp[0:len(splitC)-1], ".")
				}
			}

			category := strings.Join(splitC[0:len(splitC)-1], "/")
			categoryLower := strings.ToLower(category)

			date, err := getFileModDate(path)
			if err != nil {
				return err
			}

			tempTILs[categoryLower] = append(tempTILs[categoryLower], &til{
				Name:     name,
				category: category,
				Path:     path,
				Date: 	  date,
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

const tilDateFormat = "2006-01-02"

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

func getFileModDate(file string) (time.Time, error) {

	output, err := sh.Output("git", "log", "-1", "--format=%cd", file)
	if err != nil {
		return time.Time{}, err
	} else if output == "" {
		return time.Now(), nil
	}

	return time.Parse("Mon Jan 2 15:04:05 2006 -0700", output)
}

func GenerateReadme() error {

	tils, numTILs, err := listTILs()
	if err != nil {
		return err
	}

	var templateReadme string
	{
		fcont, err := ioutil.ReadFile("README.template.md")
		if err != nil {
			return err
		}
		templateReadme = string(fcont)
	}

	var outputReadmeBuf *bytes.Buffer
	{
		tpl, err := template.New("readme").Parse(templateReadme)
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
