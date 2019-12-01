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
  onChange?: (val: string) => void;
  placeholder?: string;
}

export const CLASS_PREFIX = "loda-file-picker";

export const FilePicker: React.FunctionComponent<FolderPickerProps> = props => {
  const {
    name,
    required,
    disabled,
    capture,
    height,
    onChange, // Remove onChange from props
    placeholder
  } = props;
  const [file, setFile] = useState("");

  useEffect(() => {
    window.renderer.on("response.open_file", (evt, value) => {
      setFile(value.file || "");
      onChange(value.file);
    });
  });

  let inputValue;
  if (file === "") {
    inputValue = "";
  } else {
    inputValue = file;
  }

  const buttonText = "Select File";

  const handleFileChange = e => {};
  const handleBlur = e => {};
  const handleButtonClick = e => {
    window.renderer.send({
      evt: "request.open_file",
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
