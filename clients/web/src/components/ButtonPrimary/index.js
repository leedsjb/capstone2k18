import React from "react";
import styled from "styled-components";

import Button from "../Button";

const ButtonPrimary = styled(Button)`
    background: linear-gradient(
        ${props => props.theme.colors.primaryLight},
        ${props => props.theme.colors.primary}
    );
`;

export default ButtonPrimary;
