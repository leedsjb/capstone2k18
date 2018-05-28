import styled from "styled-components";
import Box from "../Box";

const StyledBox = styled(Box)`
    background-color: ${props => (props.active ? "#EBEBEB" : "white")};
    &:hover {
        background-color: #EBEBEB};
    }
`;

export default StyledBox;
