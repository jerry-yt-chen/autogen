package generator

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	pkgErr "github.com/pkg/errors"
)

type CategoryGenerator struct {
}

func New() *CategoryGenerator {
	return &CategoryGenerator{}
}

func (g *CategoryGenerator) Generate(path string) error {
	genAbsDir, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	projectName := filepath.Base(genAbsDir)
	d := data{
		AbsGenProjectPath: genAbsDir + "/example",
		ProjectPath:       "github.com/17media/" + projectName,
		ProjectName:       projectName,
		Signal:            "<-sig",
	}
	fmt.Println("projectName: ", projectName)
	if err := g.genFromTemplate(getTemplateSets(), d); err != nil {
		return err
	}
	return nil
}

type templateSet struct {
	templateFilePath string
	templateFileName string
	genFilePath      string
}

type templateEngine struct {
	Templates []templateSet
	currDir   string
}

func getTemplateSets() []templateSet {
	tt := templateEngine{}
	templatesFolder := filepath.Join(".", "template/")
	fmt.Printf("walk:%s\n", templatesFolder)

	filepath.Walk(templatesFolder, tt.visit)
	fmt.Println("tt.Templates: ", tt.Templates)
	return tt.Templates

}
func (e *templateEngine) visit(path string, f os.FileInfo, err error) error {
	fmt.Println("path: ", path)
	if err != nil {
		fmt.Println("err: ", err)
		return err
	}

	if ext := filepath.Ext(path); ext == ".tmpl" {
		templateFileName := filepath.Base(path)

		genFileBaeName := strings.TrimSuffix(templateFileName, ".tmpl") + ".go"
		genFileBasePath, err := filepath.Rel(filepath.Join("", "", "template"), filepath.Join(filepath.Dir(path), genFileBaeName))
		if err != nil {
			return pkgErr.WithStack(err)
		}

		templ := templateSet{
			templateFilePath: path,
			templateFileName: templateFileName,
			genFilePath:      filepath.Join(e.currDir, genFileBasePath),
		}

		e.Templates = append(e.Templates, templ)

	} else if mode := f.Mode(); mode.IsRegular() {
		templateFileName := filepath.Base(path)

		basepath := filepath.Join("", "", "template")
		targpath := filepath.Join(filepath.Dir(path), templateFileName)
		genFileBasePath, err := filepath.Rel(basepath, targpath)
		if err != nil {
			return pkgErr.WithStack(err)
		}

		templ := templateSet{
			templateFilePath: path,
			templateFileName: templateFileName,
			genFilePath:      filepath.Join(e.currDir, genFileBasePath),
		}

		e.Templates = append(e.Templates, templ)
	}

	return nil
}

type data struct {
	AbsGenProjectPath string // The Abs Gen Project Path
	ProjectPath       string //The Go import project path (eg:github.com/fooOrg/foo)
	ProjectName       string //The project name which want to generated
	Signal            string
}

func (g *CategoryGenerator) genFromTemplate(templateSets []templateSet, d data) error {
	for _, tmpl := range templateSets {
		fmt.Println(tmpl)
		if err := g.tmplExec(tmpl, d); err != nil {
			return err
		}
	}
	return nil
}

func unescaped(x string) interface{} { return template.HTML(x) }

func (g *CategoryGenerator) tmplExec(tmplSet templateSet, d data) error {
	tmpl := template.New(tmplSet.templateFileName)
	tmpl = tmpl.Funcs(template.FuncMap{"unescaped": unescaped})

	tmpl, err := tmpl.ParseFiles(tmplSet.templateFilePath)
	if err != nil {
		return pkgErr.WithStack(err)
	}

	relateDir := filepath.Dir(tmplSet.genFilePath)

	distRelFilePath := filepath.Join(relateDir, filepath.Base(tmplSet.genFilePath))
	distAbsFilePath := filepath.Join(d.AbsGenProjectPath, distRelFilePath)

	g.debugPrintf("distRelFilePath:%s\n", distRelFilePath)
	g.debugPrintf("distAbsFilePath:%s\n", distAbsFilePath)

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

func (g *CategoryGenerator) debugPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
