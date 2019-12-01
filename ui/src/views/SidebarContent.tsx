import * as React from "react";
import { Pane, Text } from "evergreen-ui";

export const SideBarContent: React.FunctionComponent = () => {
  return (
    <Pane>
      <Pane background="tint1" padding={24} marginBottom={16}>
        <Text>tint1</Text>
      </Pane>
      <Pane background="tint2" padding={24}>
        <Text>tint2</Text>
      </Pane>
    </Pane>
  );
};
