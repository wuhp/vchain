# vchain

## http sample

    import "github.com/wuhp/vchain"

    ...

    vchain.SetOutput("/tmp/sample.vlog")

    ...

    func CreateUser(w http.ResponseWriter, r *http.Request) {
        vr := NewRequestFromHttp(r, "WebServer", "PostUser")

        ...

        // invoke account service
        client := &http.Client{}
        req, _ := http.NewRequest(...)
        WrapHttpRequest(req, NewChainHeader(vr.Uuid, true))
        res, _ := client.Do(req)

        ...

        vr.EndWithCommit()
    }
