Ok Go Web Template
===================

![Ok Go](http://i.vimeocdn.com/video/38089409_640.jpg)

A Golang web application template with most of the stuff you've come to enjoy in Ruby on Rails.

The `.env` files are also checked into the repo for reference.

Read more about this on my blog: LINK HERE!!!

Setting up a development environment
------------------------------------

ADD MIGRATE

1. Have a working go environment and a running Postgres DB on `5432`
2. Clone this repo to your `GOPATH/src` folder
2. Install all packages (`go get`), and also `github.com/codegangsta/gin` and `gulp.js`
3. Create the development and test DB's with the Postgres `createdb` command
3. Run `gulp server`
4. Open `localhost:3000`

You probably want to remove the `.env` files from being checked into the repo, but I left them in for reference.

Deploying to Heroku
-------------------

1. `heroku create -b https://github.com/kr/heroku-buildpack-go.git`
2. `heroku config:set GO_ENV=production`
3. `godep save`
4. `gulp assets:precompile`
5. `git add .; git commit -m 'adding compiled assets and dependencies'`
6. `git push heroku master`

Keep in mind
------------

- Doesn't restart when you edit `.html` files

TODO
----

- [ ] Test on heroku from clean slate
- [ ] Comment all of the code
