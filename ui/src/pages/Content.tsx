import * as React from "react";
import { useState, useEffect } from "react";
import { useInterval } from "../hooks";
import { Button, TextField } from "@material-ui/core";

const Content: React.FunctionComponent<{}> = () => {
  const [downloadURL, setDownloadURL] = useState("");
  const [errorString, setErrorString] = useState("");
  const [progress, setProgress] = useState(0.0);
  const [isDone, setIsDone] = useState(false);
  const [label, setLabel] = useState("");

  const isDownloading = () => {
    return !isDone && label !== "";
  };

  useEffect(() => {
    window.renderer.on("response.progress.download", (evt, val) => {
      setProgress(val.progress);
    });
    window.renderer.on("error.create_download", (evt, val) => {
      console.log(val);
    });

    window.renderer.on("response.download_list", (evt, val) => {
      console.log(val);
    });

    window.renderer.on("response.progress.download.done", (evt, val) => {
      console.log("response.progress.download.done");
      setIsDone(true);
    });
    window.renderer.on("download_label", (evt, val) => {
      setLabel(val);
    });
  });

  useInterval(
    () => {
      // Your custom logic here
      window.renderer.send({
        evt: "request.download_progress",
        val: JSON.stringify({
          id: parseInt(label)
        })
      });
    },
    isDownloading() ? 1000 : null
  );

  return (
    <div>
      <p>Progress: {label}</p>
      <p>Progress: {progress}</p>
      <p>Error: {errorString}</p>
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
              url: "https://www.youtube.com/watch?v=IoRV6UBFSRM",
              destination: "/Users/darkcl/Downloads/tmp"
            })
          });
        }}
      >
        Start
      </Button>
      <Button
        onClick={() => {
          window.renderer.send({
            evt: "request.download_list",
            val: JSON.stringify({
              type: "url"
            })
          });
        }}
      >
        Get List
      </Button>
    </div>
  );
};

export default Content;
