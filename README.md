# Lisn Server

Lisn Server allows you to manipulate database and serve files related to the
**[Lisn Music Streaming Service]**.

[Lisn Music Streaming Service]: https://github./com/sharpvik/Lisn



## Getting Started

### Install Go

To compile from source, you need to have **Go** installed on your machine! You
can try installing it through your package manager of choice like this:

```bash
# Using apt manager (Debian, Ubuntu, and related)
apt-get install golang-go

# Using pacman manager (Arch, Manjaro and related)
pacman -S go

# Using yum manager (Fedora, CentOS, and related)
yum install golang
```

Alternatively, you can download [Go binary distributions][bin], and go through
the [installation process][install].

[bin]: https://golang.org/dl/
[install]: https://golang.org/doc/install
[GOPATH]: https://github.com/golang/go/wiki/SettingGOPATH


### Dependencies & Config Files

As soon as you have **Go** installed and running on your machine, you need to do
the following:

```bash
go get github.com/lib/pq    # required to interface with PostrgreSQL
go get github.com/sharpvik/lisn-server  # Lisn Server source files
```

This command will fetch the whole GitHub repo and put it into a specific place
on your computer.

#### Configure RootFolder

Go to the project's root folder:

```bash
cd $(go env GOPATH)/src/github.com/sharpvik/lisn-server
```

Once you are in that folder, you'll discover a folder called `pub`. This folder
is going to contain all public files that your server will serve on demand.
Actually, there will be another important folder called `storage` which contains
all the songs and album covers.

These two folders must lie in the same folder which we will call `RootFolder`.
You may use this project's root folder as your `RootFolder` but you don't have
to. You can place `RootFolder` wherever you have permission to create folders.

```bash
# For example:
~$ mkdir dev
~$ ls
go/     dev/
~$ cd dev
~/dev$ mkdir lisn-root
~/dev$ mv $(go env GOPATH)/src/github.com/sharpvik/lisn-server/pub lisn-root/
~/dev$ cd lisn-root
~/dev/lisn-root$ mkdir storage
~/dev/lisn-root$ cd storage
~/dev/lisn-root/storage$ mkdir albums songs
~/dev/lisn-root/storage$ cd ../../
~/dev$ tree lisn-root
lisn-root
├───pub
│   └───fail
└───storage
    ├───albums
    └───songs
```

After you setup your `RootFolder`, you'll assign its absolute path to the
`RootFolder` constant in `lisn-server/config/config.go` file.

On my machine it looks like this (because I use this project's root folder as my
`RootFolder`):

```go
RootFolder = "/home/sharpvik/go/src/github.com/sharpvik/lisn-server"
```

#### Configure Database Constants

Create `lisn-server/config/secret.go`. This file is part of the `package config`
and it is supposed to contain constants that will allow you to connect to the
database. Here's how it looks like (you can literally `copy+paste` all of the
following in your newly created `secret.go` and substitute with your values):

```go
package config

const (
    DBhost = "localhost"
    DBport = 5432
    DBuser = "user"
    DBpassword = "***"
    DBname = "Lisn"
)
```

Now, your **Go** server is good to go, however you still need to build the
client side if you want to use [Lisn Web App]. Follow the link -- you'll find
all deployment instructions there.

[Lisn Web App]: https://github.com/sharpvik/lisn-web-app


### Database

Lisn is a fairly young project. There isn't a way to quickly upload albums onto
the server and register them in the database. I had to do it myself via the
`psql` prompt while simultaneously saving files to `pub`. You can develop your
own mechanisms if you wish, and if you do, please share!


### Run, Build or Install

```bash
cd $(go env GOPATH)/src/github.com/sharpvik/lisn-server

go run      # compiles and runs without creating any binary executables

go build    # puts binary file called `lisn-server` into the project folder

go install  # creates binary file at $(go env GOPATH)/bin/lisn-server
```



## Contribute

All contributions to the Lisn project are greately appreciated. I know, the
phrase is a cliché but trust me, any contribution you make
**creates a ton of difference**.


### Ways To Help

**Scout through the [Issues]**, look for the ones you think you can fix and
*go ahead*.

[Issues]: https://github.com/sharpvik/lisn-server/issues

Found a bug? -- **create a new issue** for the rest of us to see.

And of course, you are always welcome to `fork + git clone`, and then do
whatever you want. If you think that your version works better than the one we
have published here -- **issue a pull request**!


### Sister Repos

Lisn Server is part of a bigger family. Maybe you could also help with some of
these:

- [Lisn Web App] - web app written in [Vue.js]
- [Lisn Design] - all graphics realted stuff

[Vue.js]: https://vuejs.org
[Lisn Design]: https://github.com/sharpvik/lisn-design



## Authors

- **[Viktor Rozenko]** - *Initial work*

[Viktor Rozenko]: https://github.com/sharpvik



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
