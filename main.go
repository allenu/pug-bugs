package main

import (
    "net/http"
    "github.com/eknkc/pug"
)

type Item struct {
    Text string
    Value int
}

type SomeObject struct {
    Items []Item
    MatchValue int
}

func init() {
    http.HandleFunc("/referenceOuterWithinEach", eachHandler)
    http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    if tpl, err := pug.CompileFile("templates/index"); err == nil {
        if err == nil {
            params := map[string]interface{}{
                "object": SomeObject{
                    Items: []Item{
                        Item{Text:"apple",  Value: 123},
                        Item{Text:"banana", Value: 456},
                    },
                    MatchValue: 123,
                },
            }
            err = tpl.Execute(w, params)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    } else {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func eachHandler(w http.ResponseWriter, r *http.Request) {
    if tpl, err := pug.CompileFile("templates/each"); err == nil {
        if err == nil {
            params := map[string]interface{}{
                "object": SomeObject{
                    Items: []Item{
                        Item{Text:"apple",  Value: 123},
                        Item{Text:"banana", Value: 456},
                        Item{Text:"orange", Value: 789},
                    },
                    MatchValue: 123,
                },
            }
            err = tpl.Execute(w, params)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    } else {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

