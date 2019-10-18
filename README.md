# webgopher

Quick 'n Dirty prototype of a Web to Gopher proxy service that allows you to
access the greater World-Wide-Web via the GOpher protocol by proxying the URL
selected by the selector to the web and converting the content so something
legible for Gopher clients.

**NB:** This is very much work-in-progress.

## Install

```#!go
$ go get githubcom/prologic/webgopher
```

## Usage

Run the `webgopher` daemon:

```#!sh
$ webgopher
```

Use your favorite Gopher client and pass in the URL you wish to browse on the
WEB as the selector:

```#!sh
$ lynx gopher://localhost:7000/1http://www.wikipedia.org/
```

![Screenshot](/screenshot.png)

## License

webgopher is licensed under the terms of the MIT License.
