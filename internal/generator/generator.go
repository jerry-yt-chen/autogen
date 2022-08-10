package generator

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	pkgErr "github.com/pkg/errors"
)

type ProjectType int

const (
	TypeAPI ProjectType = iota
	TypeComponent
)

var projectTypeMap = map[string]ProjectType{
	"category": TypeAPI,
	"comp":     TypeComponent,
}

var projectFolderMap = map[ProjectType]string{
	TypeAPI:       "api",
	TypeComponent: "comp",
}

type templateSet struct {
	templateFilePath string
	templateFileName string
	genFilePath      string
}

type data struct {
	RootPath    string
	ProjectType string
	ProjectPath string //The Go import project path (eg:github.com/fooOrg/foo)
	ProjectName string //The project name which want to generate
	Signal      string
}

type ProjectGenerator struct {
	f embed.FS
	data
}

func NewProjectGenerator(f embed.FS, appName string, projType ProjectType) ProjectGenerator {
	rootPath, _ := os.Getwd()
	return ProjectGenerator{
		f: f,
		data: data{
			ProjectType: projectFolderMap[projType],
			ProjectPath: "github.com/17media/" + appName,
			ProjectName: appName,
			RootPath:    filepath.Join(rootPath, appName),
			Signal:      "<-sig",
		},
	}
}

func (g *ProjectGenerator) Generate() error {
	templateSets, err := g.getTemplates()
	if err != nil {
		log.Println("getTemplates: ", err.Error())
		return err
	}

	for _, tmpl := range templateSets {
		if err = g.generate(tmpl); err != nil {
			return err
		}
	}
	return nil
}

func (g *ProjectGenerator) getTemplates() ([]templateSet, error) {
	var sets []templateSet
	tempRoot := filepath.Join("templates", getTemplateName(g.ProjectType))
	if err := fs.WalkDir(g.f, tempRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println("walk fail: ", err.Error())
			return err
		}
		relPath, _ := filepath.Rel(tempRoot, path)
		if strings.Contains(relPath, "cmd/main.go") {
			relPath = fmt.Sprintf("cmd/%s/main.go.tmpl", g.ProjectName)
		}
		if !d.IsDir() {
			ext := filepath.Ext(path)
			fileName := filepath.Base(path)
			tmpl := templateSet{
				templateFilePath: path,
				templateFileName: fileName,
				genFilePath:      filepath.Join(filepath.Dir(relPath), wrapFileName(fileName, ext)),
			}
			sets = append(sets, tmpl)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return sets, nil
}

func (g *ProjectGenerator) genFromTemplate(templateSets []templateSet) error {
	for _, tmpl := range templateSets {
		if err := g.generate(tmpl); err != nil {
			return err
		}
	}
	return nil
}

func (g *ProjectGenerator) generate(tmplSet templateSet) error {
	tmpl := template.New(tmplSet.templateFileName)
	tmpl = tmpl.Funcs(template.FuncMap{"unescaped": func(x string) interface{} {
		return template.HTML(x)
	}})

	tmpl, err := tmpl.ParseFS(g.f, tmplSet.templateFilePath)
	if err != nil {
		return pkgErr.WithStack(err)
	}

	relateDir := filepath.Dir(tmplSet.genFilePath)

	distRelFilePath := filepath.Join(relateDir, filepath.Base(tmplSet.genFilePath))
	distAbsFilePath := filepath.Join(g.data.RootPath, distRelFilePath)

	if err := os.MkdirAll(filepath.Dir(distAbsFilePath), os.ModePerm); err != nil {
		return pkgErr.WithStack(err)
	}

	dist, err := os.Create(distAbsFilePath)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	defer dist.Close()

	log.Printf("Create %s\n", distAbsFilePath)
	return tmpl.Execute(dist, g.data)
}

func ProjectTypeMap() map[string]ProjectType {
	return projectTypeMap
}
