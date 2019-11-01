import * as React from "react";
import { useState, useEffect } from "react";
import { Button, TextField } from "@material-ui/core";

const Content: React.FunctionComponent<{}> = () => {
  const [downloadURL, setDownloadURL] = useState("");
  const [progress, setProgress] = useState(0.0);
  const [label, setLabel] = useState("");

  useEffect(() => {
    window.renderer.on("progress.download", (evt, val) => {
      setProgress(val.progress);
    });
    window.renderer.on("download_label", (evt, val) => {
      setLabel(val);
      setInterval(() => {
        window.renderer.send({
          evt: "request.download_progress",
          val: JSON.stringify({
            label: val
          })
        });
      }, 1000);
    });
  });

  return (
    <div>
      <p>Progress: {label}</p>
      <p>Progress: {progress}</p>
      <TextField
        fullWidth
        id="download-url"
        label="Download URL"
        margin="normal"
        onChange={evt => {
          setDownloadURL(evt.target.value);
        }}
      />
      <Button
        onClick={() => {
          window.renderer.send({
            evt: "request.create_download",
            val: JSON.stringify({
              url:
                "http://releases.ubuntu.com/18.04.3/ubuntu-18.04.3-desktop-amd64.iso",
              destination: "/Users/darkcl/Downloads/tmp"
            })
          });
        }}
      >
        Start
      </Button>
    </div>
  );
};

export default Content;
