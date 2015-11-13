# vchain
Go client for vchain server, should be integated with end user's application.
Usually it can be implemented as a middleware in web application.

## http sample

    import "github.com/wuhp/vchain"

    ...

    vchain.SetOutput("/tmp/sample.vlog")

    ...

    func CreateUser(w http.ResponseWriter, r *http.Request) {
        vr := vchain.NewRequestFromHttp(r, "WebServer", "PostUser")

        ...

        // invoke account service
        client := &http.Client{}
        req, _ := http.NewRequest(...)
        vchain.WrapHttpRequest(req, vchain.NewChainHeader(vr.Uuid, true))
        res, _ := client.Do(req)

        ...

        vr.EndWithCommit()
    }
