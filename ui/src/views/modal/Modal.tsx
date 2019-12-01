import * as React from "react";

import { DownloadModal } from "./DownloadModal";
import { SettingsModal } from "./SettingsModal";

export const Modal: React.FunctionComponent = () => {
  return (
    <>
      <DownloadModal />
      <SettingsModal />
    </>
  );
};
