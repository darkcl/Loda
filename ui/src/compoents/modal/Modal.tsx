import * as React from "react";
import { Pane, Dialog, Text } from "evergreen-ui";

import { DownloadModal } from "./DownloadModal";
import { ModalStore, ModalType, ModalActions } from "../../store";

export const Modal: React.FunctionComponent = () => {
  const modalState = React.useContext(ModalStore.State);
  const modalDispatch = React.useContext(ModalStore.Dispatch);

  const close = () => {
    modalDispatch({
      type: ModalActions.DISMISS
    });
  };

  const getModal = (type: ModalType) => {
    switch (type) {
      case ModalType.Settings: {
        return (
          <Pane>
            <Text>Settings</Text>
          </Pane>
        );
      }
      case ModalType.Download: {
        return <DownloadModal />;
      }
      default: {
        return <p>Hello file picker</p>;
      }
    }
  };

  const getModalTitle = (type: ModalType) => {
    switch (type) {
      case ModalType.Settings: {
        return "Settings";
      }
      case ModalType.Download: {
        return "Download";
      }
      default: {
        return "A Modal";
      }
    }
  };

  return (
    <Dialog
      isShown={modalState.modalType !== ModalType.None}
      title={getModalTitle(modalState.modalType)}
      onCloseComplete={() => close()}
      confirmLabel="Confirm"
    >
      {modalState.modalType !== ModalType.None &&
        getModal(modalState.modalType)}
    </Dialog>
  );
};
