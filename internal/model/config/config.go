package config

import (
    "errors"
    "fmt"
    gcfg "gopkg.in/gcfg.v1"
)

type (

    Config struct {
        Server            ServerConfig
        Database          map[string]*DBConfig
    }

    ServerConfig struct {
        Port       string
    }

    DBConfig struct {
        Dialect       string
        Host          string
        Port          int
        User          string
        Password      string
        DBname        string
    }

)


func NewConfig() *Config {
    return &Config{}
}

func (cfg *Config) ReadConfig() error {
    var e error

    path := []string{
		"/etc/user/",
		"./configs/etc/user/",
		"../configs/etc/user/",
		"../../configs/etc/user/",
		"../../../configs/etc/user/",
		"../../../../configs/etc/user/",
         // Path is wrong so it will give error. path should be where my ini file is.
    }
    var eMsg string
    for i, val := range path {
        file := fmt.Sprintf("%suser.development.ini", val)
        // Read file using gcfg
        e = gcfg.ReadFileInto(cfg, file)
        if e == nil {
            break
        } else {
            eMsg += "Unable to load the ini file" + ", err:" + e.Error()
            if i != len(path) {
                eMsg += "; "
            }
        }
    }

    if e != nil {
        return errors.New("failed to load file!!!, " + eMsg)
    }

    return nil
}