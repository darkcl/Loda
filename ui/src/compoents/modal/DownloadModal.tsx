import * as React from "react";
import { useState } from "react";
import {
  Dialog,
  Pane,
  Label,
  Tablist,
  FilePicker,
  Tab,
  Textarea
} from "evergreen-ui";
import { FolderPicker } from "../FolderPicker";
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

export const DownloadModal: React.FunctionComponent = () => {
  const modalState = React.useContext(ModalStore.State);
  const modalDispatch = React.useContext(ModalStore.Dispatch);

  const close = () => {
    modalDispatch({
      type: ModalActions.DISMISS
    });
  };

  const [selectedIndex, setSelectedIndex] = useState(0);
  const [inputURL, setInputURL] = useState("");
  return (
    <Dialog
      isShown={modalState.modalType === ModalType.Download}
      title="Download"
      onCloseComplete={() => close()}
      confirmLabel="Confirm"
    >
      <Tablist>
        {["By URL", "By File"].map((tab, index) => (
          <Tab
            key={tab}
            isSelected={selectedIndex === index}
            onSelect={() => setSelectedIndex(index)}
          >
            {tab}
          </Tab>
        ))}
      </Tablist>
      <Pane flex="1" paddingTop={16}>
        {selectedIndex === 0 ? (
          <>
            <Textarea
              id="urls"
              placeholder="Enter multiple URLs, Separated by newline"
              onChange={e => setInputURL(e.target.value)}
              value={inputURL}
            />
          </>
        ) : (
          <FilePicker
            marginBottom={32}
            accept=".png"
            onChange={files => console.log(files)}
            placeholder="Select file to download"
          />
        )}
        <Label>Destination</Label>
        <FolderPicker onChange={val => {}} />
      </Pane>
    </Dialog>
  );
};
