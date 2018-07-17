# Self Contained Web App Boilerplate

A self-contained web app.

## Structure

The root of the project contains the web server/backend.

The `web` directory is also a React app, bootstrapped with `create-react-app`.

## Development flow

The React app in the `web` directory contains a proxy to `http://127.0.0.1:8080`. In development, you will want to take advantage of the hot-reloading functionality present in your web framework, that's why the proxy exists.

In a production build, you won't need the proxy, since the back-end serves the compiled front-end assets.

## Production flow

In order to distribute your application, all you need is to run:

    $ ./build.sh

The script will:

1. Compile your front-end assets
2. Use [Statik](https://github.com/rakyll/statik) to version your assets as a Go file
3. Build your back-end, with your React application built-in the binary

`build.sh` is probably what you should customize if you're using something other than React for the `web` layer.