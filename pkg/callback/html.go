package callback

// This file defines the HTML pages returned by the Auth callback server

const WWW_CALLBACK = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Callback: Code Received</title>
  </head>
  <body>
    <h1>Callback: Code Received</h1>
    <p>This page was used to receive the Authorisation Code from the API.</p>
    <p>You may safely close this window.</p>
  </body>
</html>`

const WWW_ERROR = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Callback: Error</title>
  </head>
  <body>
    <h1>Authorisation Error</h1>
    <p>The APS authorisation request failed.</p>
    <p>You may safely close this window.</p>
  </body>
</html>`
