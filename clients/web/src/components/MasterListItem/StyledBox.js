import styled from "styled-components";
import Box from "../Box";

const StyledBox = styled(Box)`
    background-color: ${props =>
        props.active ? props.theme.colors.gray6 : "white"};
    &:hover {
        background-color: ${props => props.theme.colors.gray6};
    }
`;

export default StyledBox;
