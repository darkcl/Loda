import * as React from "react";
import { useState } from "react";
import { Dialog, Pane, Label, Tablist, Tab, Textarea } from "evergreen-ui";
import { FolderPicker } from "../FolderPicker";
import { FilePicker } from "../FilePicker";
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

  const submitForm = () => {
    console.log(`Submit Form with ${taskType}`);

    if (destination.length === 0) {
      return;
    }

    if (taskType === "url") {
      const urls = inputURL.split("\n");
      if (urls.length === 0) {
        return;
      }

      urls.forEach(url => {
        window.renderer.send({
          evt: "request.create_download",
          val: JSON.stringify({
            url: url,
            destination: destination
          })
        });
      });
    } else {
    }
  };

  const [selectedIndex, setSelectedIndex] = useState(0);
  const [inputURL, setInputURL] = useState("");
  const [destination, setDestination] = useState("");

  const [taskType, setTaskType] = useState<DownloadType>("url");

  return (
    <Dialog
      isShown={modalState.modalType === ModalType.Download}
      title="Download"
      onCloseComplete={() => close()}
      onConfirm={close => {
        submitForm();
        close();
      }}
      confirmLabel="Confirm"
    >
      <Tablist>
        {["By URL", "By File"].map((tab, index) => (
          <Tab
            key={tab}
            isSelected={selectedIndex === index}
            onSelect={() => {
              setTaskType(index === 0 ? "url" : "file");
              setSelectedIndex(index);
            }}
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
            onChange={files => console.log(files)}
            placeholder="Select file to download"
          />
        )}
      </Pane>
      <Pane flex="1" paddingTop={16}>
        <Label paddingTop={8}>Destination</Label>
        <FolderPicker
          onChange={val => {
            setDestination(val);
          }}
        />
      </Pane>
    </Dialog>
  );
};
