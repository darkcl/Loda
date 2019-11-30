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

const SidebarItem: React.FunctionComponent<SidebarItemProps> = ({
  className,
  target,
  name
}) => {
  return (
    <a className={className} href={`#${target}`}>
      {name}
    </a>
  );
};

const StyledSidebarItem = styled(SidebarItem)`
  padding: 6px 8px 6px 16px;
  text-decoration: none;
  font-size: 25px;
  color: #818181;
  display: block;

  &:hover {
    color: #f1f1f1;
  }
`;

const Sidebar: React.FunctionComponent<SidebarProps> = ({ className }) => {
  return (
    <div className={className}>
      <StyledSidebarItem target="about" name="About" />
      <StyledSidebarItem target="about" name="About" />
      <StyledSidebarItem target="about" name="About" />
      <StyledSidebarItem target="about" name="About" />
    </div>
  );
};

export const StyledSidebar = styled(Sidebar)`
  height: 100%;
  width: 160px;
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  background-color: #111;
  overflow-x: hidden;
  padding-top: 20px;
`;
