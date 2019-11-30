import * as React from "react";
import styled from "styled-components";

import { StyledSidebar } from "../compoents";

export const MainLayout: React.FunctionComponent = props => {
  return (
    <div>
      <StyledSidebar />
      <MainContainer>{props.children}</MainContainer>
    </div>
  );
};

const MainContainer = styled.div`
  margin-left: 160px;
  padding: 0px 10px;
`;
