# 🌳 Z - My Personal Command Tree

Based on [Bonzai](https://github.com/rwxrob/bonzai), this project contains the utilities I use on a day to day basis.
See it as a replacement to the bunch of Bash functions that I usually put in my `.zshrc` file.
Feel free to take stuff from this repository but **don't expect support from me**. 😛

## Features

- Generate UUIDv4
- Generate API key (16 chars hexadecimal string)
- Find which process listen to a given TCP port (MacOS only)
- Kill the process listening to a given TCP port (MacOS only)
- More to come... 👀

## Install

There's a MacOS binary that you can download (e.g using cURL):

```bash
curl https://github.com/arnaudmorisset/z/releases/download/v1.0.0/z
```

For other systems, you'll need to build it yourself. 👷‍♀️

## Tab Completion

To activate bash completion just use the `complete -C` option from your `.bashrc` or command line.
There is no messy sourcing required.
All the completion is done by the program itself.

```bash
complete -C z z
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.
