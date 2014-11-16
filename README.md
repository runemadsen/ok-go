Ok Go Web Template
===================

![Ok Go](http://i.vimeocdn.com/video/38089409_640.jpg)

This is a Go web application I created while moving from Ruby to Go. It has many of the defaults you have come to know from Ruby on Rails, as long as some helpful boilerplate code to help you get started writing web applications in Go. The `.env` files are also checked into the repo for reference.

Read more about this on my blog: LINK HERE!!!

Setting up a development environment
------------------------------------

1. Have a working go environment and a running Postgres DB on `5432`
2. Clone this repo to your `GOPATH/src` folder
2. Install all packages (`go get`), and also `github.com/codegangsta/gin` and `gulp.js`
3. Create the development and test DB's with the Postgres `createdb` command
3. Run `gulp server`
4. Open `localhost:3000`

You probably want to remove the `.env` files from being checked into the repo, but I left them in for reference.

Deploying to Heroku
--------------

1. `heroku create -b https://github.com/kr/heroku-buildpack-go.git`
2. `heroku config:set GO_ENV=production`
3. `gulp deploy:prepare`
4. `git add .; git commit -m 'adding compiled assets and dependencies'`
5. `git push heroku master`

The `gulp deploy:prepare` task runs a few tasks to compile and minify the assets for production, as well as update th godep dependencies.

Keep in mind
------------

- Only auto-migrations. Not real migrations.
- Doesn't restart when you edit `.html` files

TODO
----

- [ ] Use `pat` instead of `gorilla.mux`
- [ ] Style the posts
- [ ] Write route spec for posts route
- [ ] Write model spec for post
- [ ] Set /assets max-age to loooong time. http://play.golang.org/p/fpETA9_1oo
- [ ] Assets n'stuff when deploying?
