# Lisn Server

Lisn Server allows you to manipulate database and serve files related to the
**[Lisn Music Streaming Service]**.

[lisn music streaming service]: https://github.com/Lisn-Rocks/meta

## Getting Started

Clone this github repo. Once done, run the `setup.sh` script.

This will result in a few things:

- A root folder will be created for all server files at `~/Public/lisn`
  - This is the initial layout of the `~/Public/lisn` folder.
  - The location of the files can be changed (described in )
- Example config files will be unpacked in the repo under the `config` folder
- An example `.env` file will be placed in the root of the repo

Ensure that you have these packages installed:

- **docker** and **docker-compose**
  - Requied to quickly set up and migrate (if needed) a database and will later on be used to ship the entire application with one command.
- **go** (for development purposes only)

> Install **ImageMagick** and make sure that it has the proper decode deligates
> for the `.jpg` and `.png` formats (at least). The `convert` command is
> essential for the server's proper functioning.

#### RootFolder Tree

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

#### Change Config

The container requires the following environmental variables (should be defined in .env):

- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_DB`

`config.go` has a few more constants that should be changed for security purposes:

- `Hash` that is used to hash the password
- `Salt` that is added to master password before hash

The rest, such as the `RootFolder` for the location of the server files can be kept the same

> `config.go` inside `config` folder and `.env` are both already functional out of the box
> and can be used for testing purposes, but it is strongly recommended to change the variables/constants above
> when running the service for public use

Now, your **Go** server is good to go, however you still need to build the
client side if you want to use [Lisn Web App]. Follow the link -- you'll find
all deployment instructions there.

[lisn web app]: https://github.com/Lisn-Rocks/web-app

### Database

Lisn is a fairly young project, however, there is a mechanism in place that
allows maintainers to quickly initialize database and upload albums into the
service!

#### Initial Setup and Migrations

The predefined volume for empty tables is located in the `sql` directory and
is automatically mounted when running a container. All the additions to the tables
will not be saved unless specified so in docker-compose file

**TODO**: Explain how add persistent data to the database and explain migrations

#### Album Uploads

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

## Contribute

All contributions to the Lisn project are greately appreciated. I know, the
phrase is a cliché but trust me, any contribution you make
**creates a ton of difference**.

### Ways To Help

We are happy to hear your feedback directly or using github:

1. **Scout through the [Issues]**, look for the ones you think you can fix and
   _go ahead_.

1. Found a bug? -- **create a new issue** for the rest of us to see.

1. And of course, you are always welcome to `fork + git clone`, and then do
   whatever you want. If you think that your version works better than the one we
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
[aleksei martirosov]: https://github.com/sharpvik

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
