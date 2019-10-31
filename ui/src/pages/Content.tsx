import * as React from "react";
import { Button, TextField } from "@material-ui/core";

const Content: React.FunctionComponent<{}> = () => {
  return (
    <div>
      <TextField
        fullWidth
        id="download-url"
        label="Download URL"
        margin="normal"
      />
      <Button>Start</Button>
    </div>
  );
};

export default Content;
