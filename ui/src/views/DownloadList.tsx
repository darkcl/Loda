import * as React from "react";

import { Button, Pane, Heading } from "evergreen-ui";
import { DownloadTaskItem } from "../compoents";
import { ModalStore, ModalActions } from "../store";

export const DownloadList: React.FunctionComponent = () => {
  const modalDispatch = React.useContext(ModalStore.Dispatch);

  return (
    <Pane flexDirection="column">
      <Pane display="flex" padding={16} borderRadius={3}>
        <Pane flex={1} alignItems="center" display="flex">
          <Heading size={600}>Tasks</Heading>
        </Pane>
        <Pane>
          <Button
            marginRight={16}
            onClick={() => {
              modalDispatch({
                type: ModalActions.SHOW_SETTINGS_MODAL
              });
            }}
          >
            Settings
          </Button>
          <Button
            appearance="primary"
            onClick={() => {
              modalDispatch({
                type: ModalActions.SHOW_DOWNLOAD_MODAL
              });
            }}
          >
            Add tasks
          </Button>
        </Pane>
      </Pane>

      <Pane flex={1} alignItems="center" display="flex">
        <DownloadTaskItem></DownloadTaskItem>
      </Pane>
    </Pane>
  );
};
