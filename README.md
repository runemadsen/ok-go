Ok Go Web Template
===================

![Ok Go](http://i.vimeocdn.com/video/38089409_640.jpg)

A Golang web application template with most of the stuff you've come to enjoy from Ruby on Rails. For more information of why I built this, [read the blog post](LINK HEEEERE). Pull requests are more than encouraged!

### Features

The following features are a part of the template. You will probably end up using only a selected handful of them for your specific projects.

- **Middleware** support with [Negroni](https://github.com/codegangsta/negroni). The template has [a single custom middleware handler](https://github.com/runemadsen/ok-go/blob/master/config/assets.go#L14-L47) to serve as inspiration, but [more can be added](https://github.com/runemadsen/ok-go/blob/master/config/app.go#L34-L38) easily if needed.

- **Routing** with [Gorilla Mux](https://github.com/gorilla/mux). You can replace this with your favorite Golang routing library if needed. This is your `controllers` folder in Rails.

- **Database ORM** support with [Gorm](https://github.com/jinzhu/gorm). The template ships with a [simple Go struct](https://github.com/runemadsen/ok-go/blob/master/models/post.go) and some [basic CRUD routes](https://github.com/runemadsen/ok-go/blob/master/routes/posts.go).

- **Database migrations** with [Gomigrate](https://github.com/DavidHuie/gomigrate) and [Gofer](https://github.com/chuckpreslar/gofer) for easy command line syntax.

- **Template** rendering with [Render](https://github.com/unrolled/render). This is your `views` folder in Rails.

- **Asset pipeline** with [Gulp](http://gulpjs.com/). Although not specifically a Golang setup, it generates digested `.coffee` and `.scss`  assets, as well as a `manifest.json` with the file paths. A Golang [`asset_path`](https://github.com/runemadsen/ok-go/blob/master/templates/layouts/layout.html#L4-L5) helper is available in the templates. Digested assets are disabled in development mode for ease of development.

- **Configuration** is stored in the OS environment, with support for development (`.env`) and test (`.env.test`) environments with [Godotenv](https://github.com/joho/godotenv).

- **Testing** with [Ginkgo](http://onsi.github.io/ginkgo/) and [Gomega](http://onsi.github.io/gomega/). Golang prefers the tests to live alongside the actual files, and will automatically ignore files named `*_test.go` during compilation. This template has both simple [route tests](https://github.com/runemadsen/ok-go/blob/master/routes/posts_test.go) and [model tests](https://github.com/runemadsen/ok-go/blob/master/models/post_test.go).


### Getting Started

This guide assumes that you have a working Go environment, and Postgres running on port `5432`.

1. Create your dev and test databases. `createdb okgo_dev; createdb okgo_test`
2. `git clone` this repo to your `GOPATH/src` folder.
3. Create a `.env` file and a `.env.test` file in the new folder root. This will be the place to put environment specific variables. You only need one for now: `DATABASE_URL=postgres://MYUSER@localhost:5432/okgo_ENVIRONMENT?sslmode=disable`
4. Install all imported packages with `go get`, and a few extra command-line requirements:
    - `go get -u github.com/codegangsta/gin`
    - `go get -u github.com/chuckpreslar/gofer/gofer`
    - `sudo npm install -g gulp`

That's it. If you start using this template for your own application, you might want to rename the `runemadsen/ok-go` packages to your own repository name.

### Development

To run the development server, run `gulp server` and open `localhost:3000` in your browser. Gin will recompile the app code on any changes to `.go` files, but does currently not support watching `.html` template files. Gulp will recompile the assets on any changes.

### Test

Run `ginkgo -r` from the root folder to run the tests.

### Heroku

This template is built to *just work* on Heroku. Right now, it requires you to commit the golang dependencies and precompiled assets to the Git repo.

1. `heroku create -b https://github.com/kr/heroku-buildpack-go.git`
2. Add a database to your new heroku server.
2. `heroku config:set GO_ENV=production`
3. `godep save`
4. `gulp assets:precompile`
5. `git add .; git commit -m 'adding compiled assets and dependencies'`
6. `git push heroku master`
7. `heroku run gofer db:migrate`
