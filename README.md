# Hated It

A Slack app for personal use.

![Hated It](https://media.giphy.com/media/ZKLcZPHPEZnd6/giphy.gif)

## Development

### Setup

* [install][gobuffalo-install] `buffalo`
* [install][gobuffalo-install-soda] `soda`
* [install][ngrok-install] `ngrok`

### Execution

These only need to be run once to setup the databases before running.

```shell
cp .env.development .env
soda create -a
```

Run the development server with hot reload.

```shell
soda migrate
buffalo dev
```

### Remote Access

Create a public tunnel using `ngrok`.

```shell
ngrok http 3000
```

<!-- links -->

[gobuffalo-install]: https://gobuffalo.io/en/docs/getting-started/installation/
[gobuffalo-install-soda]: https://gobuffalo.io/en/docs/db/toolbox#installing-cli-support
[ngrok-install]: https://dashboard.ngrok.com/get-started/setup
