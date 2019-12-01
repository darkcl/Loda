import * as React from "react";
import styled from "styled-components";

interface SidebarProps {
  className?: string;
}

interface SidebarItemProps {
  className?: string;
  target: string;
  name: string;
}

const Sidebar: React.FunctionComponent<SidebarProps> = ({
  className,
  children
}) => {
  return <div className={className}>{children}</div>;
};

export const StyledSidebar = styled(Sidebar)`
  height: 100%;
  width: 160px;
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  overflow-x: hidden;
  padding-top: 20px;
`;
