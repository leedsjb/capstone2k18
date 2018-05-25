import styled from "styled-components";

import Box from "../Box";

const GradientBox = styled(Box)`
    background: linear-gradient(
        ${props => props.theme.colors.primaryLight},
        ${props => props.theme.colors.primary}
    );
`;

export default GradientBox;
