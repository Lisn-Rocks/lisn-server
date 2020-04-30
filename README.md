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


### Dependencies & Config Files

As soon as you have **Go** installed and running on your machine, you need to do
the following:

```bash
go get github.com/lib/pq    # required to interface with PostrgreSQL
go get github.com/sharpvik/lisn-server  # Lisn Server source files
```

This command will fetch the whole GitHub repo and put it into a specific place
on your computer.

Run the `setup.sh` script. It will create a root folder for all server files at
`~/Public/lisn` while also unpacking all the config file templates from `.pkg`.
This is the initial layout of the `~/Public/lisn` folder.

#### RootFolder Tree

```bash
~/Public/lisn
├── logs
├── pub
│   └── fail
└── storage
    ├── albums
    └── songs
```

#### Change Config

In the `config` folder you'll find two files: `config.go` and `secret.go`.
The `config.go` file contains general setup settings, while the `secret.go` file
should contain the login details for your database.

Both these files are merely templates, yet `config.go` is completely functional
out of the box unless you decide to relocate or rename your `~/Public/lisn`
folder. If that's the case, don't forget to change the `RootFolder` constant to
reflect whatever change you've made. Same thing goes for every subfolder that
has a mention in the `config.go` file.

> You must change `secret.go` to match the correct login data if you want your
> server to work.

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
