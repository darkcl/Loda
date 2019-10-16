import * as React from "react";
import * as ReactDOM from "react-dom";
import { IPCRenderer, IMessage } from "./ipc";
import { GoogleLink } from "./compoents/google";

const render = () =>
  ReactDOM.render(
    <div>
      <p>Hello World!</p>
      <GoogleLink />
    </div>,
    document.getElementById("root")
  );

window.renderer = new IPCRenderer();

window.onclick = function(e) {
  const elem = e.target as Element;
  if (elem.localName === "a") {
    e.preventDefault();
    window.renderer.send({
      evt: "openlink",
      val: elem.getAttribute("href")
    });
  }
};

render();
