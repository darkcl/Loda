import * as React from "react";
import { useState } from "react";
import { Dialog, Pane, Text } from "evergreen-ui";
import { ModalStore, ModalType, ModalActions } from "../../store";

type DownloadType = "url" | "file";

export interface DownloadForm {
  type: DownloadType;
  input: string;
  destination: string;
}

export interface DownloadModalProps {
  onFormChange: (form: DownloadForm) => void;
}

export const SettingsModal: React.FunctionComponent = () => {
  const modalState = React.useContext(ModalStore.State);
  const modalDispatch = React.useContext(ModalStore.Dispatch);

  const close = () => {
    modalDispatch({
      type: ModalActions.DISMISS
    });
  };

  return (
    <Dialog
      isShown={modalState.modalType === ModalType.Settings}
      title="Download"
      onCloseComplete={() => close()}
      confirmLabel="Confirm"
    >
      <Pane>
        <Text>Settings</Text>
      </Pane>
    </Dialog>
  );
};
