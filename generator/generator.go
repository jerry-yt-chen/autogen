package generator

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	pkgErr "github.com/pkg/errors"
)

type templateSet struct {
	templateFilePath string
	templateFileName string
	genFilePath      string
}

type data struct {
	AbsGenProjectPath string // The Abs Gen Project Path
	ProjectPath       string //The Go import project path (eg:github.com/fooOrg/foo)
	ProjectName       string //The project name which want to generated
	Signal            string
}

type ProjectGenerator struct {
	f embed.FS
}

func NewProjectGenerator(f embed.FS) *ProjectGenerator {
	return &ProjectGenerator{f: f}
}

func (g *ProjectGenerator) Generate() error {
	currPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	projectName := filepath.Base(currPath)
	if err != nil {
		return err
	}
	fmt.Println(currPath)
	templateSets, err := g.getTemplates()
	if err != nil {
		fmt.Println("getTemplates: ", err.Error())
		return err
	}
	//fmt.Println("templateSets: ", templateSets)
	d := data{
		AbsGenProjectPath: currPath,
		ProjectPath:       "github.com/17media/" + projectName,
		ProjectName:       projectName,
		Signal:            "<-sig",
	}

	for _, tmpl := range templateSets {
		if err = g.gen(tmpl, d); err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}

func (g *ProjectGenerator) getTemplates() ([]templateSet, error) {
	fmt.Println("getTemplates")
	var sets []templateSet
	if err := fs.WalkDir(g.f, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("walk fail: ", err.Error())
			return err
		}
		relPath, _ := filepath.Rel("templates", path)
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

func (g *ProjectGenerator) genFromTemplate(templateSets []templateSet, d data, f embed.FS) error {
	fmt.Println("genFromTemplate")
	for _, tmpl := range templateSets {
		fmt.Println(tmpl)
		if err := g.gen(tmpl, d); err != nil {
			return err
		}
	}
	return nil
}

func (g *ProjectGenerator) gen(tmplSet templateSet, d data) error {
	fmt.Println("gen")
	tmpl := template.New(tmplSet.templateFileName)
	tmpl = tmpl.Funcs(template.FuncMap{"unescaped": func(x string) interface{} {
		return template.HTML(x)
	}})

	fmt.Println("tmplSet.templateFilePath: ", tmplSet.templateFilePath)
	dat, err := os.ReadFile(tmplSet.templateFilePath)
	fmt.Println("abc: ", string(dat))

	tmpl, err = tmpl.ParseFS(g.f, tmplSet.templateFilePath)
	if err != nil {
		fmt.Println("fuckoff")
		return pkgErr.WithStack(err)
	}

	relateDir := filepath.Dir(tmplSet.genFilePath)

	distRelFilePath := filepath.Join(relateDir, filepath.Base(tmplSet.genFilePath))
	distAbsFilePath := filepath.Join(d.AbsGenProjectPath, distRelFilePath)

	fmt.Printf("distRelFilePath:%s\n", distRelFilePath)
	fmt.Printf("distAbsFilePath:%s\n", distAbsFilePath)

	if err := os.MkdirAll(filepath.Dir(distAbsFilePath), os.ModePerm); err != nil {
		return pkgErr.WithStack(err)
	}

	dist, err := os.Create(distAbsFilePath)
	if err != nil {
		return pkgErr.WithStack(err)
	}
	defer dist.Close()

	fmt.Printf("Create %s\n", distRelFilePath)
	return tmpl.Execute(dist, d)
}
func wrapFileName(path string, ext string) string {
	switch ext {
	case ".ops":
		return strings.TrimSuffix(path, ext)
	case ".tmpl":
		return strings.TrimSuffix(path, ext) + ".go"
	default:
		return path
	}
}
