# Loda

![](https://github.com/darkcl/Loda/workflows/Build%20Loda/badge.svg)
![](https://github.com/darkcl/Loda/workflows/Release%20Loda/badge.svg)
![](https://github.com/darkcl/Loda/workflows/Test/badge.svg)

Loda project for golang desktop application development

## What is included?

### Frontend

- React
- Typescript
- Webpack
- Renderer for sending / receiving messages

### Backend

- Golang
- Webview (forked version with additional menu handling)
- IPC for sending / receiving messages

## Usage

### Frontend

```ts
// Send message to golang backend

window.renderer.send({
  evt: "openlink",
  val: elem.getAttribute("href")
});

// Recieve message from golang backend
window.renderer.on("testing", (evt, val) => {
  console.log(`Recieve event "${evt}" from golang`);
});
```

### Backend

```go
// Getting main process
ipcMain := ipc.SharedMain()

// Sending Message to frontend
ipcMain.Send("testing", map[string]string{"testing": "123"})

// Recieving Message from frontend
ipcMain.On(
  "openlink",
  func(event string, value interface{}) interface{} {
    url := value.(string)
    helpers.OpenBrowser(url)
    return nil
  })
```
