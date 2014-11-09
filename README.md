golang-rails-template
=====================

An opiniated template for golang web applications inspired by Ruby on Rails. But soooooooo much faster.

Read more about this on my blog: LINK HERE!!!

Development environment
-----------------------

1. Have a working go environment
2. Clone this repo to your `GOPATH/src` folder
2. Install all packages (`go get`), and also `github.com/codegangsta/gin` and `gulp.js`
3. Run `gulp server`
4. Open `localhost:3000`

Deploying to Heroku
--------------

1. `heroku create -b https://github.com/kr/heroku-buildpack-go.git`
2. `heroku config:set MARTINI_ENV=production`
3. `git push heroku master`

You need to run a few commands before deploying (compile assets for production, update godep dependencies). Do this by running `gulp deploy:prepare`.

TODO
----

- [ ] Fix test
- [ ] Set /assets max-age to loooong time
- [ ] Model specs
- [ ] Migrations with Goose
- [ ] DB with gorm
- [ ] Gulp won't work on heroku.
- [ ] Small CRUD examples