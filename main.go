package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"store/app/api/products"
	"store/app/middleware"
	"store/types/cfg"
	"store/util"
	"time"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"

	"store/app/api/auth"
	"store/app/front"
)

func main() {
	var (
		wait            time.Duration
		configPath      string
		doNewConfigPath bool
		doMigrate       bool
		migrateDown     bool
	)

	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&configPath, "cfg", "config.yaml", "path to the cfg file")
	flag.BoolVar(&doNewConfigPath, "new-cfg", false, "create a new cfg file")
	flag.BoolVar(&doMigrate, "migrate", false, "run db migrations")
	flag.BoolVar(&migrateDown, "down", false, "run db migrations down")
	flag.Parse()

	if doMigrate {
		util.Migrate(configPath, migrateDown)
		return
	}

	config, err := util.GetConfig(configPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if doNewConfigPath {
				config = &cfg.Config{
					Db: cfg.Database{
						Host:     "127.0.0.01",
						Port:     3306,
						User:     "store",
						Pass:     "",
						Database: "store",
					},
					Http: cfg.Http{
						Bind:         "127.0.0.1",
						Port:         8080,
						WriteTimeout: 15,
						ReadTimeout:  15,
						IdealTimeout: 60,
					},
					Redis: cfg.Redis{
						Host:   "127.0.0.1",
						Port:   6379,
						Prefix: "store",
						Pass:   "",
					},
				}
				err = util.DumpConfig(configPath, config)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("New cfg file created at %s\n", configPath)
				return
			} else {
				log.Fatalf("Config file not found: %s", configPath)
			}
		}
		log.Fatal(err)
	}

	log.Println("Starting server")

	util.InitStore(config)
	err = util.InitDb(config.Db)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", front.HomeHandler)
	r.HandleFunc("/license", front.LicenseHandler)
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/img/favicon.ico")
	})

	CSRF := csrf.Protect(securecookie.GenerateRandomKey(32), csrf.Secure(false))
	r.Use(CSRF)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	r.Use(middleware.LogRequest)

	ApiRouter := r.PathPrefix("/api").Subrouter()

	ApiRouter.HandleFunc("/products", products.GetProducts).Methods("GET")
	ApiRouter.HandleFunc("/auth", auth.SteamAuth).Methods("GET")
	ApiRouter.HandleFunc("/auth/logout", auth.LogOut).Methods("GET")
	ApiRouter.HandleFunc("/auth/callback", auth.SteamCallback).Methods("GET")

	AdminRouter := r.PathPrefix("/admin").Subrouter()
	AdminRouter.Use(middleware.CheckAdmin)
	AdminRouter.HandleFunc("/", front.AdminHandler).Methods("GET")
	AdminRouter.HandleFunc("/products", products.Add).Methods("POST")
	AdminRouter.HandleFunc("/products/{id}", products.Delete).Methods("DELETE")
	AdminRouter.HandleFunc("/products/{id}", products.Patch).Methods("PUT")

	http.Handle("/", r)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Http.Bind, config.Http.Port),
		Handler: r,

		WriteTimeout: time.Second * time.Duration(config.Http.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(config.Http.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(config.Http.IdealTimeout),

		ErrorLog: log.New(log.Writer(), "ERROR: ", log.LstdFlags),
	}
	go func() {
		log.Println("Listening on", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("ListenAndServe: %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		panic(err)
	}

	if err = util.DumpConfig(configPath, config); err != nil {
		panic(err)
	}
}
