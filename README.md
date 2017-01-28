[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

# Company Repo

This code base evolves from [the React tutorial](http://facebook.github.io/react/docs/tutorial.html).

## TODO / WISHLIST
- Chat app for concurrency
- Upload a file to server. Take pdf and *parse* pdf and present it back to browser.
- Backend to support multiple data layers: in-mem DB, SQL..
- Use JSHint to static analysis on JS code to catch bugs early.
- Backend to support paging
- Autoscrolling
- Animation
- Data visualization/analytics
- Make a website responsive and fun to interact with
- CSS for better looking sites
- Typescript+React+Webpack
- Phabricator for task/project management


## How to start
1. webpack public/index.js (or: webpack -w)
2. mv bundle.js public/
3. python server.py (or: go run server.go)

4. dal$ go generate
   dal$ go test (or go -v test)

## To use

There are several simple server implementations included. They all serve static files from `public/` and handle requests to `/api/comments` to fetch or add data. Start a server with one of the following:

### Start commands for various backend options

```sh
npm install
npm start
--
npm install
node server.js
--
pip install -r requirements.txt
python server.py
--
ruby server.rb
--
php server.php
--
go run server.go
--
cpan Mojolicious
perl server.pl
```
And visit <http://localhost:3000/>. Try opening multiple tabs!

### Changing the port

```sh
PORT=3001 node server.js
```
