import styled from "styled-components";
import { Box } from "grid-styled";

const RedBox = styled(Box)`
    background: radial-gradient(
        ${props => props.theme.colors.airlift1} 0%,
        ${props => props.theme.colors.airlift2} 100%
    );
`;

export default RedBox;
