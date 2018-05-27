import styled from "styled-components";

import Box from "../Box";

const GradientBox = styled(Box)`
    background: radial-gradient(
        circle at top left,
        ${props =>
            props.firstcolor in props.theme.colors
                ? props.theme.colors[props.firstcolor]
                : props.firstcolor},
        ${props =>
            props.secondcolor in props.theme.colors
                ? props.theme.colors[props.secondcolor]
                : props.secondcolor}
    );
`;

export default GradientBox;
