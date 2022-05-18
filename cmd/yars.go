package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gitlab.com/yars-ai/yars/ast"
	"gitlab.com/yars-ai/yars/infrastructure"
	"gitlab.com/yars-ai/yars/repository"
	"gitlab.com/yars-ai/yars/services/engine"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	folderFlag     = "folder"
	fileFlag       = "file"
	unitFlag       = "unit"
	fileFormatFlag = "file-format"
	separatorFlag  = "sep"
	projectFlag    = "path-project"

	hostFlag = "host"
	portFlag = "port"

	dbHostFlag     = "db-host"
	dbPortFlag     = "db-port"
	dbUserFlag     = "db-user"
	dbPasswordFlag = "db-password"
	dbNameFlag     = "db-name"

	defaultPort = "3030"
	defaultHost = "localhost"
)

func pathFolder() *cli.PathFlag {
	return &cli.PathFlag{
		Name:     folderFlag,
		Required: true,
	}
}

func pathFile() *cli.PathFlag {
	return &cli.PathFlag{
		Name:     fileFlag,
		Required: true,
	}
}

func stringFileFormat() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     fileFormatFlag,
		Required: true,
	}
}

func stringSeparator() *cli.StringFlag {
	return &cli.StringFlag{
		Name:        separatorFlag,
		DefaultText: ",",
	}
}

func stringUnit() *cli.StringFlag {
	return &cli.StringFlag{
		Name: unitFlag,
	}
}

func pathProject() *cli.PathFlag {
	return &cli.PathFlag{
		Name:     projectFlag,
		Required: true,
	}
}

func stringHost() *cli.StringFlag {
	return &cli.StringFlag{
		Name:  hostFlag,
		Value: defaultHost,
	}
}

func stringPort() *cli.StringFlag {
	return &cli.StringFlag{
		Name:  portFlag,
		Value: defaultPort,
	}
}

func stringDBHost() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     dbHostFlag,
		Required: true,
	}
}

func stringDBUser() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     dbUserFlag,
		Required: true,
	}
}

func stringDBPassword() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     dbPasswordFlag,
		Required: true,
	}
}

func stringDBName() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     dbNameFlag,
		Required: true,
	}
}

func stringDBPort() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     dbPortFlag,
		Required: true,
	}
}

func openGormConnection(host, port, user, password, dbName string) (*gorm.DB, error) {
	log.Print("start connection to database")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		host,
		port,
		user,
		password,
		dbName,
	)

	newLogger := logger.New(
		// TODO change stdout
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		return nil, errors.Wrap(err, "can't open db connection")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(5)

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Print("database connection established")

	return db, nil
}

func readProject(pathProject string) (*ast.Project, error) {
	b, err := os.ReadFile(pathProject)
	if err != nil {
		return nil, fmt.Errorf("can't read file %s: %w", pathProject, err)
	}

	project, err := ast.Parse(string(b))
	if err != nil {
		return nil, fmt.Errorf("can't parse project file %s: %w", pathProject, err)
	}
	log.Printf("project %s parsed", project.Name.Value)

	return project, nil
}

func readProjects(pathProjects string) ([]*ast.Project, error) {
	log.Print("start reading projects")

	files, err := ioutil.ReadDir(pathProjects)
	if err != nil {
		return nil, fmt.Errorf("can't read directory %s: %w", pathProjects, err)
	}

	var projects []*ast.Project

	for _, f := range files {
		if !f.IsDir() {
			project, err := readProject(path.Join(pathProjects, f.Name()))
			if err != nil {
				return nil, err
			}
			projects = append(projects, project)
		}
	}

	log.Print("reading projects done")

	return projects, nil
}

func buildEngine(project *ast.Project, db *gorm.DB) (engine.Engine, error) {
	engine, err := engine.NewBuilder(repository.NewFabric(db, project.Name.Value), project).Build()
	if err != nil {
		return nil, fmt.Errorf("can't build project %s: %w", project.Name.Value, err)
	}
	if err := engine.MigrateAll(context.Background()); err != nil {
		return nil, fmt.Errorf("can't migrate project %s: %w", project.Name.Value, err)
	}
	log.Printf("project %s builded", project.Name.Value)
	return engine, nil
}

func buildEngines(projects []*ast.Project, db *gorm.DB) ([]engine.Engine, error) {
	log.Print("start building projects")

	var engines []engine.Engine
	for _, project := range projects {
		engine, err := buildEngine(project, db)
		if err != nil {
			return nil, err
		}
		engines = append(engines, engine)
	}

	log.Print("building projects done")
	return engines, nil
}

func actionRunProjects(c *cli.Context) error {
	pathProjects := c.Path(folderFlag)
	projects, err := readProjects(pathProjects)
	if err != nil {
		return err
	}

	db, err := openGormConnection(
		c.String(dbHostFlag),
		c.String(dbPortFlag),
		c.String(dbUserFlag),
		c.String(dbPasswordFlag),
		c.String(dbNameFlag),
	)

	if err != nil {
		return err
	}

	engines, err := buildEngines(projects, db)
	if err != nil {
		return err
	}

	server := infrastructure.NewServer(engines)
	server.Run(fmt.Sprintf("%s:%s", c.String(hostFlag), c.String(portFlag)))

	return nil
}

func actionRetrainProject(c *cli.Context) error {
	panic("not implemented")
}

func actionLoadProject(c *cli.Context) error {
	project, err := readProject(c.String(projectFlag))
	if err != nil {
		return err
	}

	db, err := openGormConnection(
		c.String(dbHostFlag),
		c.String(dbPortFlag),
		c.String(dbUserFlag),
		c.String(dbPasswordFlag),
		c.String(dbNameFlag),
	)
	if err != nil {
		return err
	}

	engine, err := buildEngine(project, db)
	if err != nil {
		return err
	}

	file, err := os.Open(c.String(fileFlag))
	if err != nil {
		return err
	}

	log.Print("load started")
	if c.String(unitFlag) != "" {
		unit, err := engine.UnitByName(c.String(unitFlag))
		if err != nil {
			return err
		}
		_, err = unit.Upsert(context.Background(), file)
		if err != nil {
			return err
		}
	}
	log.Print("load finished")

	return nil
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "run",
				Flags: []cli.Flag{
					pathFolder(),
					stringHost(),
					stringPort(),
				},
				Action: actionRunProjects,
			},
			{
				Name: "retrain",
				Flags: []cli.Flag{
					pathProject(),
				},
				Action: actionRetrainProject,
			},
			{
				Name: "load",
				Flags: []cli.Flag{
					pathFile(),
					pathProject(),
					stringFileFormat(),
					stringUnit(),
					stringSeparator(),
				},
				Action: actionLoadProject,
			},
		},
		Flags: []cli.Flag{
			stringDBHost(),
			stringDBUser(),
			stringDBPassword(),
			stringDBName(),
			stringDBPort(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
