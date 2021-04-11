# Hated It

![Hated It](https://media.giphy.com/media/ZKLcZPHPEZnd6/giphy.gif)

A Slack app for personal use. This project is in a beta state, it does some
stuff but nothing is in place to tell you if it's not working properly and
probably never will.

## What does this do?

All that this thing is meant to do is print a list of items using the Slack
block kit format for messages in response to a slash command that needs to be
setup to point at it. The name has nothing to do with what it does.

## Development

### Setup

* [install][gobuffalo-install] `buffalo`
* [install][gobuffalo-install-soda] `soda`
* [install][ngrok-install] `ngrok`

### Execution

These only need to be run once to setup the databases before running.

```shell
cp .env.development .env
docker-compose up db
soda create -a
```

Run the development server with hot reload.

```shell
docker-compose up -d db
soda migrate
buffalo dev
```

### Add Data

At least one user, associated list, and list entry needs to exist for it to
work. Commands to do this have been added that use `grift`.

```shell
grift db:create:user $USER
grift db:create:list $USER
grift db:create:item $USER "A thing in a list"
```

To see this list create a Slack app and add a custom slash command, it can be
whatever is wanted, eg. `/hatedit`. Executing that command in any chat should
print the list.

### Remote Access

Create a public tunnel using `ngrok`.

```shell
ngrok http 3000
```

<!-- links -->

[gobuffalo-install]: https://gobuffalo.io/en/docs/getting-started/installation/
[gobuffalo-install-soda]: https://gobuffalo.io/en/docs/db/toolbox#installing-cli-support
[ngrok-install]: https://dashboard.ngrok.com/get-started/setup
