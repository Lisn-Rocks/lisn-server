# Lisn Server

Lisn Server allows you to manipulate database and serve files related to the
**[Lisn Music Streaming Service]**.

[lisn music streaming service]: https://github.com/Lisn-Rocks/meta

## Getting Started

These are the instructions for your go server to start running on you machine.
For the client side head over to [Lisn Web App].

[lisn web app]: https://github.com/Lisn-Rocks/web-app

### Package Requiremenets

`go` specifically for development purposes

`docker` for setting up the database and for starting an entire
application later on with one command

> Install **ImageMagick** and make sure that it has the proper decode deligates
> for the `.jpg` and `.png` formats (at least). The `convert` command is
> essential for the server's proper functioning.

### Installation

After cloning this girhub repository, run the `setup.sh` script.

This will result in a few things:

1. A root folder will be created for all server files at `~/Public/lisn`
   - This is the initial layout of the `~/Public/lisn` folder.
   - The location of the files can be changed (described in '[Change Config](#change-config)' section)
1. Example config files will be unpacked in the repo under the `config` folder
1. An example `.env` file will be placed in the root of the repo

### Change Config

The container requires the following environmental variables
(should be defined in .env):

- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_DB`

`config.go` has a few more constants that should be changed for security purposes:

- `Hash` that is used to hash the password
- `Salt` that is added to master password before hash

The rest, such as the `RootFolder` for the location of the server files can be kept the same.

> `config.go` inside `config` folder and `.env` are both already functional out of the box
> and can be used for testing purposes, but it is strongly recommended to change the variables/constants above
> when running the service for public use

### Album Uploads

At this point, all you need to do is upload some albums. It is very simple. All
songs must be uploaded as albums to ensure clarity and consistency. One album is
a folder that contains:

1. Album cover image in any appropriate format (preferably `.jpg` or `.png`)
   with _1:1_ side ratio (square) called `cover.jpg` or `cover.png`;
2. Audio files corresponding to every song in the album (preferably `.mp3`) that
   must be named the same way they are called. For example, audio file for song
   "Mustapha" by Queen must be named `Mustapha.mp3`.
3. The metadata file called `meta.json` that contains crucial information
   required to process an upload.

What follows is an example of the `meta.json` file. As you can see from the
comments, some fields are optional and they have default values. This measure
saves maintainers' time and storage space.

> No comments are allowed in the actual `meta.json` file as they will confuse
> the parser! I use them here for explanation purposes only.

```json
{
  "artist": "MadeUp",
  "album": "Fuss In The Air",
  "genres": ["Hard Rock", "Heavy Metal"],
  "coverext": ".png", // ommittable; defaults to ".jpg"
  // songs must be ordered the same way they are in an actual album!
  "songs": [
    {
      "song": "First One",
      "audioext": ".wav", // ommittable; defaults to ".mp3"
      "feat": ["Someone Else"] // ommittable; defaults to []
    },
    {
      "song": "Make Them On The Go!"
    },
    {
      "song": "You Know How It Is"
    },
    {
      "song": "Just Like That"
    }
  ]
}
```

This JSON album data is supposed to represent a folder (album) with the
following structure:

```bash
Fuss In The Air
├── cover.png
├── meta.json
├── "First One.wav"
├── "Make Them On The Go.mp3"
├── "You Know How It Is.mp3"
└── "Just Like That.mp3"
```

Now that you have your album folder ready to go, zip it without the folder
itself (only the files go into the archive) and use `/pub/upload` site on your
server to make Lisn process and save your music!

### Start the Server and Database

Once all the config is changed and you've added a few albums, simply run:

```bash
sudo docker-compose up
```

and both the server and the database will start functioning.

## Development Information

If you want to add features or test the current ones to our server, this section
wil explain some of the aspects you may want to know.

#### RootFolder Initial Tree

```bash
~/Public/lisn
├── logs
├── pub
│   ├── fail
│   └── upload
└── storage
    ├── albums
    └── songs
```

### Database Setup

To only setup the database, run:

```bash
sudo docker-compose up db
```

After the first time, a container with the predefined tables will be created for you,
where you can enter the relevant data.

To stop the database, run:

```bash
sudo docker-compose down
```

> Note that all the data you have entered in the databases will be saved.

If you want to delete the entries in the database, run:

```bash
sudo docker-compose down --volumes
```

You will still have the schemas for the tables in the fresh database the next time you
will start the database.

### Running the Server

```bash
go run main.go
```

Please make sure to run:

```bash
sudo docker-compose build
```

afterwards to rebuild the `go` image for the container, in case
you will want to run the whole thing together again.

### Migrating the database

If you need the entries that were previously entered in the database,
make sure to retreive the `db_data` volume created by your docker upon execution
and mount it accordingly in `docker-compose.yaml` upon migration.

## Known Bugs and TODO's

- [ ] Log file is not actually created
- [ ] Make a defer function that will display the log upon finishing the execution
- [ ] Improve the server structure
- [ ] Find another way for user to change the hash and salt

## Contribute

We would love to get help from you through feedback or pull requests. Every small
contribution will be greatly appreciated!

### Ways To Help

**Scout through the [Issues]**, look for the ones you think you can fix and
_go ahead_.

Found a bug? -- **create a new issue** for the rest of us to see.

And of course, you are always welcome to `fork + git clone`, and then do
whatever you want (use of [Development Information](#development-information) is still
**highly recommended**).
If you think that your version works better than the one we
have published here -- **issue a pull request**!

[issues]: https://github.com/Lisn-Rocks/server/issues

### Sister Repos

Lisn Server is part of a bigger family. Maybe you could also help with some of
these:

- [Lisn Web App] - web app written in [Vue.js]
- [Lisn Design] - all graphics realted stuff

[vue.js]: https://vuejs.org
[lisn design]: https://github.com/Lisn-Rocks/design

## Authors

- **[Viktor Rozenko]** - _Initial work_
- **[Aleksei Martirosov]** - _Docker setup and bug fixes_

[viktor rozenko]: https://github.com/sharpvik
[aleksei martirosov]: https://github.com/aleksimart

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

[billie thompson]: https://gist.github.com/PurpleBooth
[readme template]: https://gist.github.com/PurpleBooth/109311bb0361f32d87a2
