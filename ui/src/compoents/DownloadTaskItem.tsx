import * as React from "react";
import { Pane, Text } from "evergreen-ui";
import ProgressBar from "@atlaskit/progress-bar";

export const DownloadTaskItem: React.FunctionComponent = () => {
  return (
    <Pane flex={1} display="flex" borderBottom="default" flexDirection="column">
      <Pane
        paddingTop={16}
        paddingBottom={8}
        justifyContent="space-between"
        display="flex"
      >
        <Text>Name</Text>
        <Text>30%</Text>
      </Pane>

      <ProgressBar isIndeterminate />
      <Pane
        paddingTop={8}
        paddingBottom={16}
        justifyContent="space-between"
        display="flex"
      >
        <Text>Upload</Text>
        <Text>Download</Text>
      </Pane>
    </Pane>
  );
};
