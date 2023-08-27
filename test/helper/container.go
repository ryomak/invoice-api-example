package helper

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/docker"
	"path"
	"runtime"

	"github.com/friendsofgo/errors"
	"github.com/ory/dockertest"
	"github.com/ryomak/invoice-api-example/infrastructure/client/db"
	"github.com/ryomak/invoice-api-example/infrastructure/env"
	"log"
	"time"
)

const containerName = "test_db"

type dbContainerManager struct {
	Pool *dockertest.Pool

	Resource *dockertest.Resource
	Conn     *db.Conn
}

func NewDBContainerManager() (*dbContainerManager, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, errors.Wrap(err, "Could not connect to docker: %v")
	}

	return &dbContainerManager{
		Pool: pool,
	}, nil
}

func (d *dbContainerManager) CreateOrGetConn() (*db.Conn, func() error, error) {
	// create db
	if err := d.createContainer(); err != nil {
		return nil, nil, err
	}
	conn, err := d.getConn()
	if err != nil {
		return nil, func() error {
			return d.closeContainer()
		}, err
	}

	return conn, func() error {
		return d.closeContainer()
	}, nil
}

func (d *dbContainerManager) createContainer() error {

	_, pwd, _, _ := runtime.Caller(0)

	cfg := env.GetCfg()
	opts := dockertest.RunOptions{
		Name: containerName,
		/*
			docker: m1 macを利用するためにはPlatformにlinux/x86_64を設定する必要あり
			https://matsuand.github.io/docs.docker.jp.onthefly/desktop/mac/apple-silicon/
			dockertestがPlatformに対応していないため、mariadbで設定
			対応された後に、BuildAndRunWithOptionsで設定
		*/
		Repository: "mariadb",
		Tag:        "10.3",
		Env: []string{
			"MYSQL_ALLOW_EMPTY_PASSWORD=yes",
			"MYSQL_DATABASE=" + cfg.MySQLDatabase,
			"MYSQL_PASSWORD=" + cfg.MySQLPassword,
			"MYSQL_ROOT_PASSWORD=" + cfg.MySQLPassword,
			"MYSQL_USER=" + cfg.MySQLUser,
		},
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_ci",
			"--default-time-zone=utc",
		},
		Mounts: []string{
			path.Dir(pwd) + "/../../schema.sql:/docker-entrypoint-initdb.d/schema.sql",
		},
	}

	resource, err := d.Pool.RunWithOptions(&opts, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		return errors.Wrap(err, "Could not start resource")
	}
	d.Resource = resource
	return nil
}

func (d *dbContainerManager) getConn() (*db.Conn, error) {
	if d.Resource == nil || d.Pool == nil {
		return nil, errors.New("getConn: manger not invalid param")
	}

	// DB(コンテナ)との接続
	var (
		conn *db.Conn
		err  error
	)
	for i := 0; i < 3; i++ {
		if err := d.Pool.Retry(func() error {
			// DBコンテナが立ち上がってから疎通可能になるまで少しかかるのでちょっと待ったほうが良さそう
			time.Sleep(time.Second * 3)

			conn, err = db.NewWithOverride(d.Resource.GetPort("3306/tcp"))
			if err != nil {
				return err
			}
			return conn.Ping()
		}); err != nil {
			log.Fatalf("Pool.Retry: %s", err)
		}
	}
	return conn, nil
}

func (d *dbContainerManager) closeContainer() error {
	if err := d.Pool.Purge(d.Resource); err != nil {
		return fmt.Errorf("could not purge resource %w", err)
	}
	return nil
}
