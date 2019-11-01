import * as React from "react";
import { useState, useEffect } from "react";
import { Button, TextField } from "@material-ui/core";

const Content: React.FunctionComponent<{}> = () => {
  const [downloadURL, setDownloadURL] = useState("");

  useEffect(() => {
    window.renderer.on("progress.download", (evt, val) => {
      console.log(`[${evt}]: ${val}`);
    });
  });

  return (
    <div>
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
              url: "http://example.com/example.txt",
              destination: "/tmp"
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
