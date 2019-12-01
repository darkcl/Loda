import * as React from "react";

export enum ModalActions {
  SHOW_SETTINGS_MODAL = "SHOW_SETTINGS_MODAL",
  SHOW_DOWNLOAD_MODAL = "SHOW_DOWNLOAD_MODAL",
  DISMISS = "DISMISS"
}

export enum ModalType {
  None = -1,
  Download = 0,
  Settings = 1
}

const State = React.createContext(null);
const Dispatch = React.createContext(null);

interface CodeExecution {
  language: string;
  code: string;
}

interface IModalState {
  modalType: ModalType;
}

const initialState: IModalState = {
  modalType: ModalType.None
};

type ModalAction =
  | { type: ModalActions.SHOW_SETTINGS_MODAL }
  | { type: ModalActions.SHOW_DOWNLOAD_MODAL }
  | { type: ModalActions.DISMISS };

function reducer(state: IModalState, action: ModalAction): IModalState {
  switch (action.type) {
    case ModalActions.SHOW_SETTINGS_MODAL: {
      return {
        ...state,
        modalType: ModalType.Settings
      };
    }
    case ModalActions.SHOW_DOWNLOAD_MODAL: {
      return {
        ...state,
        modalType: ModalType.Download
      };
    }
    case ModalActions.DISMISS:
      return {
        ...state,
        modalType: ModalType.None
      };
    default:
      return initialState;
  }
}

// Provider
const Provider: React.FunctionComponent = ({ children }) => {
  const [state, dispatch] = React.useReducer(reducer, initialState);

  return (
    <State.Provider value={state}>
      <Dispatch.Provider value={dispatch}>{children}</Dispatch.Provider>
    </State.Provider>
  );
};

// Export
export const ModalStore = {
  State,
  Dispatch,
  Provider
};
