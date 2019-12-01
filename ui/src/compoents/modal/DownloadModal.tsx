import * as React from "react";
import { useState } from "react";
import { Pane, Label, Tablist, FilePicker, Tab, Textarea } from "evergreen-ui";
import { FolderPicker } from "../FolderPicker";

type DownloadType = "url" | "file";

export interface DownloadModalProps {
  onFormChange: (type: DownloadType, contexts: string[]) => void;
}

export const DownloadModal: React.FunctionComponent = () => {
  const [selectedIndex, setSelectedIndex] = useState(0);
  const [inputURL, setInputURL] = useState("");
  return (
    <>
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
            accept={["torrent"]}
            onChange={files => console.log(files)}
            placeholder="Select file to download"
          />
        )}
        <Label>Destination</Label>
        <FolderPicker />
      </Pane>
    </>
  );
};
