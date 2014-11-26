Ok Go Web Template
===================

![Ok Go](http://i.vimeocdn.com/video/38089409_640.jpg)

A Golang web application template with most of the stuff you've come to enjoy from Ruby on Rails. For more information of why I built this, [read the blog post](LINK HEEEERE). Pull requests are more than encouraged!

### Features

The following features are a part of the template. You will probably end up using only a selected handful of them for your specific projects.

- **Routing** with [Gorilla Mux](https://github.com/gorilla/mux). You should easily be able to replace this with your favority Golang routing library.
- Middleware and static asset support with [Negroni](https://github.com/codegangsta/negroni)
- Asset pipeline using [Gulp](http://gulpjs.com/). Generates digested asset paths, `manifest.json`, and ships Golang `asset_path` helpers.
- Database ORM support with [Gorm](https://github.com/jinzhu/gorm)
- Database migrations using [Gomigrate](https://github.com/DavidHuie/gomigrate) and [Gofer](https://github.com/chuckpreslar/gofer)
- Layout/View rendering with [Render](https://github.com/unrolled/render)
- Configuration is stored in the OS environment, with support for dev (`.env`) and test (`.env.test`) environments with [Godotenv](https://github.com/joho/godotenv)
- Testing with [Ginkgo](http://onsi.github.io/ginkgo/) and [Gomega](http://onsi.github.io/gomega/)

### Getting Started

This guide assumes that you have a working Go environment, and a Postgres DB running on port `5432`.

1. Create your dev and test databases. `createdb okgo_dev; createdb okgo_test`
2. `git clone` this repo to your `GOPATH/src` folder.
3. Create a `.env` file and a `.env.test` file in the new folder root. This will be the place to put environment specific variables. You only need one for now: `DATABASE_URL=postgres://MYUSER@localhost:5432/okgo_ENVIRONMENT?sslmode=disable`
4. Install all imported packages with `go get`, and a few extra command-line requirements:
    - `go get -u github.com/codegangsta/gin`
    - `go get -u github.com/chuckpreslar/gofer/gofer`
    - `sudo npm install -g gulp`

That's it.

### Development

To run the development server, run `gin main.go` and open `localhost:3000` in your browser. Gin will recompile the app on any changes to `.go` files, but does currently not support watching `.html` template files.

### Test

Run `ginkgo -r` from the root folder to run the tests.

### Heroku

This template is built to *just work* on Heroku. Right now, it requires you to commit the golang dependencies and precompiled assets to the Git repo.

1. `heroku create -b https://github.com/kr/heroku-buildpack-go.git`
2. `heroku config:set GO_ENV=production`
3. `godep save`
4. `gulp assets:precompile`
5. `git add .; git commit -m 'adding compiled assets and dependencies'`
6. `git push heroku master`
