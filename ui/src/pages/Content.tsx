import * as React from "react";
import { useState } from "react";
import { Button, TextField } from "@material-ui/core";

const Content: React.FunctionComponent<{}> = () => {
  const [downloadURL, setDownloadURL] = useState("");

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
            evt: "openlink",
            val: "http://google.com"
          });
        }}
      >
        Start
      </Button>
    </div>
  );
};

export default Content;
