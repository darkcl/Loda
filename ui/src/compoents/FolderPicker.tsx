import * as React from "react";
import { useState, useEffect } from "react";
import Box from "ui-box";
import { TextInput, Button } from "evergreen-ui";

interface FolderPickerProps {
  name?: string;
  required?: boolean;
  disabled?: boolean;
  capture?: boolean;
  height?: number;
  onChange?: () => void;
  placeholder?: string;
}

export const CLASS_PREFIX = "loda-folder-picker";

export const FolderPicker: React.FunctionComponent<FolderPickerProps> = props => {
  const {
    name,
    required,
    disabled,
    capture,
    height,
    onChange, // Remove onChange from props
    placeholder
  } = props;
  const [folder, setFolder] = useState("");

  useEffect(() => {
    window.renderer.on("response.open_directory", (evt, value) => {
      setFolder(value.directory || "");
    });
  });

  let inputValue;
  if (folder === "") {
    inputValue = "";
  } else {
    inputValue = folder;
  }

  const buttonText = "Select folder";

  const handleFileChange = e => {};
  const handleBlur = e => {};
  const handleButtonClick = e => {
    window.renderer.send({
      evt: "request.open_directory",
      val: ""
    });
  };

  return (
    <Box display="flex" className={`${CLASS_PREFIX}-root`} {...props}>
      <Box
        innerRef={this.fileInputRef}
        className={`${CLASS_PREFIX}-file-input`}
        name={name}
        required={required}
        disabled={disabled}
        capture={capture}
        onChange={handleFileChange}
        display="none"
      />

      <TextInput
        className={`${CLASS_PREFIX}-text-input`}
        readOnly
        value={inputValue}
        placeholder={placeholder}
        // There's a weird specifity issue when there's two differently sized inputs on the page
        borderTopRightRadius="0 !important"
        borderBottomRightRadius="0 !important"
        height={height}
        flex={1}
        textOverflow="ellipsis"
        onBlur={handleBlur}
      />

      <Button
        className={`${CLASS_PREFIX}-button`}
        onClick={handleButtonClick}
        disabled={disabled}
        borderTopLeftRadius={0}
        borderBottomLeftRadius={0}
        height={height}
        flexShrink={0}
        type="button"
        onBlur={handleBlur}
      >
        {buttonText}
      </Button>
    </Box>
  );
};
