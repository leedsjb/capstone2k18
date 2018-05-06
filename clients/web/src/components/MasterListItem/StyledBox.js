import styled from "styled-components";
import Box from "../Box";

const StyledBox = styled(Box)`
    background-color: ${props =>
        props.active ? props.theme.colors.wireframe : "white"};
    &:hover {
        background-color: ${props => props.theme.colors.wireframe};
    }
`;

export default StyledBox;
