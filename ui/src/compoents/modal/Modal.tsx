import * as React from "react";
import { Pane, Dialog } from "evergreen-ui";

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
        return <p>Settings</p>;
      }
      case ModalType.Download: {
        return <p>Download</p>;
      }
      default: {
        return <p>Hello file picker</p>;
      }
    }
  };

  return (
    <Dialog
      isShown={modalState.modalType !== ModalType.None}
      title="Dialog title"
      onCloseComplete={() => close()}
      confirmLabel="Custom Label"
    >
      {modalState.modalType !== ModalType.None &&
        getModal(modalState.modalType)}
    </Dialog>
  );
};
