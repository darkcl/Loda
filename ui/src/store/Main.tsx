import * as React from "react";
import { cloneElement } from "react";
import { ModalStore } from "./ModalStore";

const providers = [<ModalStore.Provider />];

const Store = ({ children: initial }) =>
  providers.reduce(
    (children, parent) => cloneElement(parent, { children }),
    initial
  );

export { Store, ModalStore };
