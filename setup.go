package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lonelyevil/kook"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/spf13/viper"
)

// stdout频道
var stdoutChannel string

// 转发map
var routeMap map[string]string

// 邀请map
var kookInviteUrl map[string]string

var masterID string
var botID string

var localSession *kook.Session

var token string

var db *scribble.Driver

func routeMapSetup() {
	routeMap = make(map[string]string, 0)
	s := viper.Get("kook2qq").(map[string]any)
	for k, v := range s {
		vs := v.(string)
		if k != v {
			if _, ok := routeMap[k]; !ok {
				routeMap[k] = vs
			}
			if _, ok := routeMap[vs]; !ok {
				routeMap[vs] = k
			}
		}
	}
}
func kookInviteUrlSetup() {
	kookInviteUrl = make(map[string]string, 0)
	s := viper.Get("kookinvite").(map[string]any)
	for k, v := range s {
		vs := v.(string)
		if _, ok := kookInviteUrl[k]; !ok {
			kookInviteUrl[k] = vs
		}
	}
}

func getConfig() {
	rand.Seed(time.Now().UnixNano())
	db, _ = scribble.New("./database", nil)
	viper.SetDefault("token", "0")
	viper.SetDefault("stdoutChannel", "0")
	viper.SetDefault("masterID", "")
	viper.SetDefault("oldversion", "0.0.0")
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	masterID = viper.Get("masterID").(string)
	stdoutChannel = viper.Get("stdoutChannel").(string)
	fmt.Println("stdoutChannel=" + stdoutChannel)
	token = viper.Get("token").(string)
	routeMapSetup()
	kookInviteUrlSetup()
	kookLastCacheSetup()
	msgCacheSetup()
}

func beforeShutdown() {
	msgCache.Backup()
}
