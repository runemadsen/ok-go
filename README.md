Ok Go(lang) Web Application Template
====================================

![Ok Go](http://i.vimeocdn.com/video/38089409_640.jpg)

An opiniated template for golang web applications inspired by Ruby on Rails. But soooooooo much faster.

Read more about this on my blog: LINK HERE!!!

Development environment
-----------------------

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
2. `heroku config:set MARTINI_ENV=production`
3. `gulp deploy:prepare`
4. `git add .; git commit -m 'adding compiled assets and dependencies'`
5. `git push heroku master`

You need to migrate the DB also:

1. `heroku run goose -env production up`

The `gulp deploy:prepare` task runs a few tasks to compile and minify the assets for production, as well as update th godep dependencies.

Not in this
-----------

Only auto-migrations. Not real migrations.

TODO
----

- [ ] Set up model
- [ ] Model specs
- [ ] Small CRUD example
- [ ] Set /assets max-age to loooong time. http://play.golang.org/p/fpETA9_1oo
- [ ] Assets n'stuff when deploying?
