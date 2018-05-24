import styled from "styled-components";
import Box from "../Box";

const FlexGrid = styled(Box)`
    div:not(:first-of-type) {
        background-color: cyan;
    }
`;

export default FlexGrid;
