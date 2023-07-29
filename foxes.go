package foxes

import (
    "time"
    "strconv"
    "net/http"
    "io/ioutil"
    "math/rand"
    "strings"
)

type Options struct {
    Width int
    Height int
    AspectRatio string
}

type tags struct {
    time int
    count int
}

var totals = make(map[string]tags)

func internal(tag string, o Options) (string, error) {
    var days = int(time.Now().Unix()/86400)
    if _, ok := totals[tag]; !ok || totals[tag].time != days {
        res, err := http.Get("https://foxes.cool/counts/"+tag)
        if err != nil {
            return "", err
        }
        body, err := ioutil.ReadAll(res.Body)
        res.Body.Close()
        if err != nil {
            return "", err
        }
        count, err := strconv.Atoi(string(body))
        if err != nil {
            return "", err
        }
        totals[tag] = tags{
            time: days,
            count: count,
        }
    }
    var number = strconv.Itoa(rand.Intn(totals[tag].count))
    var arguments = []string{}

    if o.Width != 0 {arguments = append(arguments, "width="+strconv.Itoa(o.Width))}
    if o.Height != 0 {arguments = append(arguments, "height="+strconv.Itoa(o.Height))}
    if o.AspectRatio != "" {arguments = append(arguments, "aspect_ratio="+o.AspectRatio)}

    var argumentString = ""
    if len(arguments) > 0 {
        argumentString = "?"+strings.Join(arguments, "&")
    }

    return "https://img.foxes.cool/"+tag+"/"+number+".jpg"+argumentString, nil
}

func Fox(o Options) (string, error) {
    return internal("fox", o)
}

func Curious(o Options) (string, error) {
    return internal("curious", o)
}

func Happy(o Options) (string, error) {
    return internal("happy", o)
}

func Scary(o Options) (string, error) {
    return internal("scary", o)
}

func Sleeping(o Options) (string, error) {
    return internal("sleeping", o)
}
