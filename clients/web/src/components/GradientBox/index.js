import styled from "styled-components";

import Box from "../Box";

const GradientBox = styled(Box)`
    background: radial-gradient(
        circle at top left,
        ${props => props.firstcolor},
        ${props => props.secondcolor}
    );
`;

export default GradientBox;
