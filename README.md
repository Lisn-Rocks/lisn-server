# [Lisn Music Streaming](http://lisn.rocks)

**Lisn** is the music that *rocks*!



## Getting Started

If you want to host Lisn on your own server, you are absolutely welcome to do
so! 


### Prerequisites

To compile from source, you need to have **Go** installed on your machine! You
can try installing it through your package manager of choice like this:

```bash
apt-get install golang-go

# or

pacman -S go
```

Alternatively, you can download [Go binary distributions][bin], go through the 
[installation process][install], and don't forget to [set the `$GOPATH`
environment variable][GOPATH]!

[bin]: https://golang.org/dl/
[install]: https://golang.org/doc/install
[GOPATH]: https://github.com/golang/go/wiki/SettingGOPATH


### Configure Go Server Files

As soon as you have **Go** installed and running on your machine, you can do the
following:

```bash
go get github.com/sharpvik/Lisn
```

This command will fetch the whole GitHub repo and put it into a specific place
on your computer. For those, who are new to Go programming, I'll give a hack.

```bash
go env GOPATH
# Prints something like /home/username/go on UNIX-based systems.
```

The above command prints out absolute path to the folder where all Go-related
things are supposed to be stored. I don't know what that path is for you, so
here I'll just call it `$GOPATH`. Knowing what your `$GOPATH` is, you can now
easily locate the newly installed `Lisn` package as follows:

```bash
cd $GOPATH/src/github.com/sharpvik/Lisn

# Remember, $GOPATH (here and below) is not actually an environment variable,
# it's just a placeholder for your own Go folder and is supposed to be
# substituted with the path printed by `go env GOPATH`.
```

You can also try ...

```bash
cd $(go env GOPATH)/src/github.com/sharpvik/Lisn
```

... instead of copy-pasting the output from `go env GOPATH` by hand.

Once you are in that folder, you'll need to change the `config.go` file a notch.
Edit the `RootFolder` constant in `Lisn/config/config.go` so that it
reflects the actual path to the `Lisn` folder on your machine.

On my machine it looks like this:

```go
RootFolder = "/home/sharpvik/go/src/github.com/sharpvik/Lisn"

// My $GOPATH is set to /home/sharpvik/go so the string in RootFolder
// reflects the exact location of the Lisn project folder on my machine.
// You need to change this string to be
//
//     RootFolder = "$YOUR_GOPATH/src/github.com/sharpvik/Lisn"
//
```

Now, your **Go** server is good to go, however we still have to build the client
side!


### Build the Client Side

First of all, check that the line

```js
publicPath: '/public', // uncomment before building for deployment
```

in `Lisn/client/vue.config.js` is uncommented! 

Also, change `ROUTE` in `Lisn/client/src/App.vue`. The `ROUTE` is an IP address 
and your server's port on LAN or WLAN you use for testing (e.g. 
`120.116.14.25:8000`) or a proper web link like `my-site.com`.

Then, run the following commands:

```bash
cd $GOPATH/src/github.com/sharpvik/Lisn

npm install     # to install node_modules from package-lock.json
npm run build   # to compile & build client side (outputs into Lisn/public)
cd ..           # to go back to project's root folder
```

Run the following command from the project's root folder to start your server.

```bash
go run main.go
```

It will immediately start serving at `localhost:8000`. To change the port, stop
the server with `CTRL+C`, edit the `Port` constant in
`Lisn/config/config.go`, restart the server.


### Build & Install

Go is actually a compiled language, however the `go run` command doesn't produce
any visible executable files. To compile `Lisn`, you can

```bash
cd $GOPATH/src/github.com/sharpvik/Lisn

go build    # puts binary `Lisn` file into the project folder

# or alternatively, use

go install  # creates binary file at $GOPATH/bin/Lisn
```



## Contributing

All contributions are welcome and appreciated! I'd be glad to accept your help
with this project. `CONTRIBUTING.md` will be added in future commits.



## Authors

- **Viktor A. Rozenko Voitenko** - *Initial work* - [sharpvik]

[sharpvik]: https://github.com/sharpvik



## License

This project is licensed under the **Mozilla Public License Version 2.0** --
see the [LICENSE](LICENSE) file for details.

Please note that this project is distributred as is,
**with absolutely no warranty of any kind** to those who are going to deploy
and/or use it. None of the authors and contributors are responsible (liable)
for **any damage**, including but not limited to, loss of sensitive data and
server machine malfunction.



## Acknowledgments

- Hat tip to [Billie Thompson] for the great [README template].

[Billie Thompson]: https://gist.github.com/PurpleBooth
[README template]: https://gist.github.com/PurpleBooth/109311bb0361f32d87a2
