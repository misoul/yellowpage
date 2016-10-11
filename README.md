[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

# Company Repo

This code base evolves from [the React tutorial](http://facebook.github.io/react/docs/tutorial.html) code base.

## TODO
- Use JSHint to static analysis on JS code to catch bugs early.
- Backend to support paging
- Autoscrolling

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

You can change the port number by setting the `$PORT` environment variable before invoking any of the scripts above, e.g.,

```sh
PORT=3001 node server.js
```
