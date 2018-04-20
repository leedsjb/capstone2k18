import styled from "styled-components";
import Box from "../Box";

const StyledBox = styled(Box)`
    &:hover {
        background-color: ${props => props.theme.colors.wireframe};
    }
`;

export default StyledBox;
