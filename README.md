# `license`
> Download open source license files for your project

`license` is a command line tool which allows you to download open source licenses from [osl.kevin.codes](http://osl.kevin.codes).

### Demo

[![asciicast](https://asciinema.org/a/Y58TBBoRlqnHdHlJiNBvN1VcC.png)](https://asciinema.org/a/Y58TBBoRlqnHdHlJiNBvN1VcC)

_Click the image above to see a demo video on asciinema.org_

### Install

#### Binary

Download the binary for your operating system from the [release tab](https://github.com/kevingimbel/license/releases) and place it somewhere inside your PATH.

#### Brew

Run the following commands in a terminal to install `license` via brew.

```sh
$ brew tap kevingimbel/tap
$ brew install license
```

#### `go get`

If brew or binary installations do not suite you, you can install license with `go get`.

```sh
$ go get github.com/kevingimbel/license
```
_Note:_ Go must be installed on the system.

### Usage

See `license help` for a list of all commands.

See `license [cmd] --help` for more about each command. `[cmd]` can be any of: `get`, `list`, `update`, or `version`.

#### `license list`

Run `license list` to see a list of all licenses and their IDs. IDs are used to retrieve a license.

```sh
$ license list
Academic Free License v3.0 (id:  AFL-3.0 )
GNU Affero General Public License v3.0 (id:  AGPL-3.0 )
Apache License 2.0 (id:  Apache-2.0 )
Artistic License 2.0 (id:  Artistic-2.0 )
```

#### `license get [id]`

`license get` downloads the license by ID to a file named `LICENSE`. Use `-f` or `--format` to specify a file format.

Download to a file named `LICENSE` (no file extension).

```sh
$ license get Apache-2.0
```

To download the license to a file named `LICENSE.txt`.

```sh
$ license get --format txt MIT
```

#### `license update`

Update the local license cache file in `$HOME/.license/license.json`. The locale cache file is not used as of version 0.1.2!

#### `license version`

Display the version of `license`. Use `license version -l` to display the Git commit ID and build date.

```sh
$ license version
0.1.2
```
