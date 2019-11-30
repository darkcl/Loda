import * as React from "react";
import * as ReactDOM from "react-dom";
import { IPCRenderer, IMessage } from "./ipc";

import { MainLayout } from "./layout";
import { DownloadList } from "./views";

const render = () =>
  ReactDOM.render(
    <MainLayout>
      <DownloadList />
    </MainLayout>,
    document.getElementById("root")
  );

window.renderer = new IPCRenderer();

function isURL(str) {
  var pattern = new RegExp(
    "^(https?:\\/\\/)?" + // protocol
    "((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.?)+[a-z]{2,}|" + // domain name
    "((\\d{1,3}\\.){3}\\d{1,3}))" + // OR ip (v4) address
    "(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*" + // port and path
    "(\\?[;&a-z\\d%_.~+=-]*)?" + // query string
      "(\\#[-a-z\\d_]*)?$",
    "i"
  ); // fragment locator
  return pattern.test(str);
}

window.onclick = function(e) {
  const elem = e.target as Element;
  if (elem.localName === "a") {
    if (!elem.getAttribute("href") && isURL(elem.getAttribute("href"))) {
      e.preventDefault();
      window.renderer.send({
        evt: "openlink",
        val: elem.getAttribute("href")
      });
    }
  }
};

render();
