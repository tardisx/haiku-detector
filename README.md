[![Go Reference](https://pkg.go.dev/badge/github.com/tardisx/haiku-detector.svg)](https://pkg.go.dev/github.com/tardisx/haiku-detector)


    import "github.com/tardisx/haiku-detector"

    haiku := haiku.Find("haiku can be found wherever you are looking with help of some code")

    if len(haiku) > 0 {
        println(haiku[0].String())
    }

    //   haiku can be found
    // wherever you are looking
    //  with help of some code

